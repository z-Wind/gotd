package gotd

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

func Test_UpdateAccessToken(t *testing.T) {
	ctx := context.Background()
	a := &Auth{}

	b, err := ioutil.ReadFile(clientsecretPath)
	if err != nil {
		t.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying the scope, delete your previously saved credentials
	// at ~/.credentials/youtube-go.json
	err = a.configFromJSON(b)
	if err != nil {
		t.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	cacheFile, err := a.tokenCacheFile("TDAmeritrade-go.json")
	if err != nil {
		t.Fatalf("Unable to get path to cached credential file. %v", err)
	}
	token, err := a.tokenFromFile(cacheFile)
	if err != nil {
		t.Fatalf("Unable to get token from %v", cacheFile)
	}

	call := postAccessToken(ctx, http.DefaultClient, TokenGrantTypeRefreshToken, a.config.ClientID)
	call = call.RefreshToken(token.RefreshToken)
	token, err = call.Do()
	if err != nil {
		t.Fatalf("Unable to update refresh_token. %v", err)
	}
	t.Logf("%+v", token)
}

func Test_postAccessTokenCall_doRequest(t *testing.T) {
	client := clientTest("", http.StatusOK)
	ctx := context.Background()

	tests := []struct {
		name    string
		c       *postAccessTokenCall
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", postAccessToken(ctx, client, TokenGrantTypeRefreshToken, "client_id").RefreshToken("refresh_token").AccessType(TokenAccessTypeOffline).RedirectURI("http://localhost:8080"), "https://api.tdameritrade.com/v1/oauth2/token", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rsp, err := tt.c.doRequest()
			if (err != nil) != tt.wantErr {
				t.Errorf("postAccessTokenCall.doRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := rsp.Request.URL.String()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postAccessTokenCall.doRequest() = %v, want %v", got, tt.want)

				buf := new(bytes.Buffer)
				buf.ReadFrom(rsp.Request.Body)
				t.Errorf("%s", buf.String())
			}
		})
	}
}

func Test_postAccessTokenCall_Do(t *testing.T) {
	client := clientTest(`{
		"access_token" : "access_token",
		"refresh_token" : "refresh_token",
		"scope" : "PlaceTrades AccountAccess MoveMoney",
		"expires_in" : 1800,
		"refresh_token_expires_in" : 7776000,
		"token_type" : "Bearer"
	  }`, http.StatusOK)
	ctx := context.Background()

	tests := []struct {
		name    string
		c       *postAccessTokenCall
		want    *oauth2.Token
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Normal", postAccessToken(ctx, client, TokenGrantTypeRefreshToken, "client_id").RefreshToken("refresh_token").AccessType(TokenAccessTypeOffline).RedirectURI("http://localhost:8080"), &oauth2.Token{AccessToken: "access_token", TokenType: "Bearer", RefreshToken: "refresh_token", Expiry: time.Date(2020, 1, 1, 0, 0, 0, 0, time.Local)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Do()
			if (err != nil) != tt.wantErr {
				t.Errorf("postAccessTokenCall.Do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !(got.AccessToken == tt.want.AccessToken && got.RefreshToken == tt.want.RefreshToken && got.TokenType == tt.want.TokenType) {
				t.Errorf("postAccessTokenCall.Do() = %v, want %v", got, tt.want)
			}
		})
	}
}
