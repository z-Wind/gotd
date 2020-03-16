package gotd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"path"
	"runtime"

	"github.com/pkg/errors"
)

// openURL opens a browser window to the specified location.
// This code originally appeared at:
//   http://stackoverflow.com/questions/10377243/how-can-i-launch-a-process-that-is-not-a-file-in-go
func openURL(url string) error {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("Cannot open URL %s on this platform", url)
	}
	return err
}

// CheckResponse returns an error (of type *Error) if the response
// status code is not 2xx.
func CheckResponse(res *http.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrapf(err, "ioutil.ReadAll")
	}

	errReply := &errorReply{}
	err = json.Unmarshal(b, errReply)
	if err != nil {
		return nil
	}

	return &Error{
		Code:    res.StatusCode,
		Message: errReply.Message,
		Body:    string(b),
		Header:  res.Header,
	}
}

// DecodeResponse decodes the body of res into target. If there is no body,
// target is unchanged.
func DecodeResponse(target interface{}, res *http.Response) error {
	if res.StatusCode == http.StatusNoContent {
		return nil
	}
	// for test
	//b, _ := ioutil.ReadAll(res.Body)
	//fmt.Printf("%s\n", string(b))
	//fmt.Printf("====================================\n")
	//return json.NewDecoder(bytes.NewBuffer(b)).Decode(target)

	return json.NewDecoder(res.Body).Decode(target)
}

// SendRequest sends a single HTTP request using the given client.
// If ctx is non-nil, it sends the request with req.WithContext
func SendRequest(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if ctx == nil {
		return client.Do(req)
	}

	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}

	return resp, err
}

// JSONReader convert struct to reader for json request
func JSONReader(v interface{}) (io.Reader, error) {
	buf := new(bytes.Buffer)

	err := json.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// ResolveRelative join path
func ResolveRelative(basePath string, elem ...string) string {
	u, err := url.Parse(basePath)
	if err != nil {
		panic(fmt.Sprintf("url.Parse failed to parse %q", basePath))
	}

	u.Path = path.Join(u.Path, path.Join(elem...))

	return u.String()
}
