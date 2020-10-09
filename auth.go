package gotd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

// This variable indicates whether the script should launch a web server to
// initiate the authorization flow or just display the URL in the terminal
// window. Note the following instructions based on this setting:
// * launchWebServer = true
//   1. Use OAuth2 credentials for a web application
//   2. Define authorized redirect URIs for the credential in the Google APIs
//      Console and set the RedirectURL property on the config object to one
//      of those redirect URIs. For example:
//      config.RedirectURL = "http://localhost:8090"
//   3. In the startWebServer function below, update the URL in this line
//      to match the redirect URI you selected:
//         listener, err := net.Listen("tcp", "localhost:8090")
//      The redirect URI identifies the URI to which the user is sent after
//      completing the authorization flow. The listener then captures the
//      authorization code in the URL and passes it back to this script.
// * launchWebServer = false
//   1. Use OAuth2 credentials for an installed application. (When choosing
//      the application type for the OAuth2 client ID, select "Other".)
//   2. Set the redirect URI to "urn:ietf:wg:oauth:2.0:oob", like this:
//      config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"
//   3. When running the script, complete the auth flow. Then copy the
//      authorization code from the browser and enter it on the command line.
const launchWebServer = true

const missingClientSecretsMessage = `
Please configure OAuth 2.0
To make this sample run, you need to populate the client_secrets.json file
found at:
   %v
with information from the {{ Google Cloud Console }}{{ https://cloud.google.com/console }}For more information about the client_secrets.json file format, please visit:
https://developers.google.com/api-client-library/python/guide/aaa_client_secrets
`

// Auth for oauth2
type Auth struct {
	redirectURL *url.URL
	config      *oauth2.Config

	useTLS          bool
	tlsCert, tlsKey string
}

// NewAuth create Auth
func NewAuth() *Auth {
	a := &Auth{}

	return a
}

// SetTLS when redirectURL is TLS, it should be setted
func (a *Auth) SetTLS(TLSCertPath, TLSKeyPath string) {
	a.tlsCert = TLSCertPath
	a.tlsKey = TLSKeyPath
	a.useTLS = true
}

// GetClient generate a Client. It returns the generated Client.
// clientsecret file format reference to client_secret.json.sample
func (a *Auth) GetClient(clientsecretPath, tokenFile string) *http.Client {
	ctx := context.Background()

	b, err := ioutil.ReadFile(clientsecretPath)
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying the scope, delete your previously saved credentials
	// at ~/.credentials/youtube-go.json
	err = a.configFromJSON(b)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	// Use a redirect URI like this for a web app. The redirect URI must be a
	// valid one for your OAuth2 credentials.
	a.config.RedirectURL = a.redirectURL.String()
	// Use the following redirect URI if launchWebServer=false in oauth2.go
	// config.RedirectURL = "urn:ietf:wg:oauth:2.0:oob"

	cacheFile, err := a.tokenCacheFile(tokenFile)
	if err != nil {
		log.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	token, err := a.tokenFromFile(cacheFile)
	if err != nil {
		if launchWebServer {
			fmt.Println("Trying to get token from web")
			token, err = a.getTokenFromWeb()
		} else {
			fmt.Println("Trying to get token from prompt")
			token, err = a.getTokenFromPrompt()
		}
		if err != nil {
			log.Fatalf("Unable get token. %v", err)
		}
		a.saveToken(cacheFile, token)
	}

	// token.Expiry is the expiry of Access Token
	// oauth2 does not change the token.Expiry even it refresh Access Token
	// so we can use it to estimate the expiry of Refresh Token
	// A Refresh Token is valid for 90 days
	// Update Refresh Token ten days before the expiry date
	Day := time.Hour * 24
	valid_duration := Day * 90
	refresh_token_expired := token.Expiry.Round(0).Add(valid_duration - Day*10).Before(time.Now())

	if refresh_token_expired {
		call := postAccessToken(ctx, http.DefaultClient, TokenGrantTypeRefreshToken, a.config.ClientID)
		call = call.RefreshToken(token.RefreshToken)
		call = call.AccessType(TokenAccessTypeOffline)
		token, err = call.Do()
		if err != nil {
			log.Fatalf("Unable update refresh_token. %v", err)
		}
		a.saveToken(cacheFile, token)
	}

	client := a.config.Client(ctx, token)
	return client
}

func (a *Auth) configFromJSON(jsonKey []byte, scope ...string) error {
	type cred struct {
		ClientID     string   `json:"client_id,omitempty"`
		ClientSecret string   `json:"client_secret,omitempty"`
		RedirectURIs []string `json:"redirect_uris,omitempty"`
		AuthURI      string   `json:"auth_uri,omitempty"`
		TokenURI     string   `json:"token_uri,omitempty"`
	}
	var j struct {
		Web       *cred `json:"web,omitempty"`
		Installed *cred `json:"installed,omitempty"`
	}
	if err := json.Unmarshal(jsonKey, &j); err != nil {
		return err
	}
	var c *cred
	switch {
	case j.Web != nil:
		c = j.Web
	case j.Installed != nil:
		c = j.Installed
	default:
		return fmt.Errorf("oauth2: no credentials found")
	}
	if len(c.RedirectURIs) < 1 {
		return errors.New("oauth2: missing redirect URL in the client_credentials.json")
	}

	for _, redirectURL := range c.RedirectURIs {
		var err error
		a.redirectURL, err = url.Parse(redirectURL)
		if err == nil {
			break
		}
	}

	a.config = &oauth2.Config{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		RedirectURL:  c.RedirectURIs[0],
		Scopes:       scope,
		Endpoint: oauth2.Endpoint{
			AuthURL:  c.AuthURI,
			TokenURL: c.TokenURI,
		},
	}

	return nil
}

// startWebServer starts a web server that listens on http://localhost:8080.
// The webserver waits for an oauth code in the three-legged auth flow.
func (a *Auth) startWebServer() (codeCh chan string, err error) {
	listener, err := net.Listen("tcp", a.redirectURL.Host)
	if err != nil {
		return nil, errors.Wrapf(err, "net.Listen")
	}
	codeCh = make(chan string)

	handFunc := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		codeCh <- code // send code to OAuth flow
		listener.Close()
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Received code: %v\r\nYou can now safely close this browser window.", code)
	})
	if a.useTLS {
		log.Printf("create TLS Server\n")
		log.Printf("certPath: %s\n", a.tlsCert)
		log.Printf("keyPath: %s\n", a.tlsKey)
		go http.ServeTLS(listener, handFunc, a.tlsCert, a.tlsKey)
	} else {
		go http.Serve(listener, handFunc)
	}

	return codeCh, nil
}

// Exchange the authorization code for an access token
func (a *Auth) exchangeToken(code string) (*oauth2.Token, error) {
	// set access_type offline for refresh token
	token, err := a.config.Exchange(oauth2.NoContext, code, oauth2.AccessTypeOffline)
	if err != nil {
		log.Fatalf("Unable to retrieve token %v", err)
	}
	return token, nil
}

// getTokenFromPrompt uses Config to request a Token and prompts the user
// to enter the token on the command line. It returns the retrieved Token.
func (a *Auth) getTokenFromPrompt() (*oauth2.Token, error) {
	authURL := a.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)

	var code string
	fmt.Printf("Go to the following link in your browser. After completing "+
		"the authorization flow, enter the authorization code on the command "+
		"line: \n%v\n", authURL)

	if _, err := fmt.Scan(&code); err != nil {
		log.Fatalf("Unable to read authorization code %v", err)
	}

	fmt.Println(authURL)
	return a.exchangeToken(code)
}

// getTokenFromWeb uses Config to request a Token.
// It returns the retrieved Token.
func (a *Auth) getTokenFromWeb() (*oauth2.Token, error) {
	codeCh, err := a.startWebServer()
	if err != nil {
		fmt.Printf("Unable to start a web server. %s", err)
		return nil, errors.Wrapf(err, "a.startWebServer")
	}

	authURL := a.config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	err = openURL(authURL)
	if err != nil {
		log.Fatalf("Unable to open authorization URL in web server: %v", err)
	} else {
		fmt.Println("Your browser has been opened to an authorization URL.",
			"This program will resume once authorization has been provided.")
		fmt.Println(authURL)
	}

	// Wait for the web server to get the code.
	code := <-codeCh
	return a.exchangeToken(code)
}

// tokenCacheFile generates credential file path/filename.
// It returns the generated credential path/filename.
func (a *Auth) tokenCacheFile(fileName string) (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}
	tokenCacheDir := filepath.Join(usr.HomeDir, ".credentials")
	os.MkdirAll(tokenCacheDir, 0700)
	return filepath.Join(tokenCacheDir,
		url.QueryEscape(fileName)), err
}

// tokenFromFile retrieves a Token from a given file path.
// It returns the retrieved Token and any read error encountered.
func (a *Auth) tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	t := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(t)
	defer f.Close()
	return t, err
}

// saveToken uses a file path to create a file and store the
// token in it.
func (a *Auth) saveToken(file string, token *oauth2.Token) {
	fmt.Println("trying to save token")
	fmt.Printf("Saving credential file to: %s\n", file)
	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

// postAccessToken https://developer.tdameritrade.com/authentication/apis/post/token-0
// The token endpoint returns an access token along with an optional refresh token
func postAccessToken(ctx context.Context, client *http.Client, grant_type, client_id string) *postAccessTokenCall {
	c := &postAccessTokenCall{
		ctx:    ctx,
		client: client,
		body: url.Values{
			"grant_type": []string{grant_type},
			"client_id":  []string{client_id},
		},
	}

	return c
}

// postAccessTokenCall https://developer.tdameritrade.com/authentication/apis/post/token-0
// The token endpoint returns an access token along with an optional refresh token
type postAccessTokenCall struct {
	ctx context.Context
	// 不能用 oauth2 產生的，因會放在 Header 導致忽略 body 中的參數
	// 建議使用 http.DefaultClient
	client *http.Client

	body url.Values
}

// RefreshToken https://developer.tdameritrade.com/authentication/apis/post/token-0
// Required if using refresh token grant
func (c *postAccessTokenCall) RefreshToken(refreshToken string) *postAccessTokenCall {
	c.body.Set("refresh_token", refreshToken)
	return c
}

// AccessType https://developer.tdameritrade.com/authentication/apis/post/token-0
// Set to offline to receive a refresh token on an authorization_code grant type request. Do not set to offline on a refresh_token grant type request.
func (c *postAccessTokenCall) AccessType(accessType string) *postAccessTokenCall {
	c.body.Set("access_type", accessType)
	return c
}

// Code https://developer.tdameritrade.com/authentication/apis/post/token-0
// Required if trying to use authorization code grant
func (c *postAccessTokenCall) Code(code string) *postAccessTokenCall {
	c.body.Set("code", code)
	return c
}

// RedirectURI https://developer.tdameritrade.com/authentication/apis/post/token-0
// Required if trying to use authorization code grant
func (c *postAccessTokenCall) RedirectURI(redirectURI string) *postAccessTokenCall {
	c.body.Set("redirect_uri", redirectURI)
	return c
}

func (c *postAccessTokenCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("User-Agent", UserAgent)
	reqHeaders.Set("Content-Type", "application/x-www-form-urlencoded")

	body := strings.NewReader(c.body.Encode())
	urls := ResolveRelative(basePath, "oauth2", "token")

	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.client, req)
}

// Do send request
func (c *postAccessTokenCall) Do() (*oauth2.Token, error) {
	res, err := c.doRequest()

	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, errors.Wrapf(err, "doRequest")
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	token := &oauth2.Token{}
	err = json.NewDecoder(res.Body).Decode(token)
	token.Expiry = time.Now().Add(time.Second * 1800)
	if err != nil {
		return nil, errors.Wrapf(err, "json.NewDecoder")
	}

	return token, nil
}
