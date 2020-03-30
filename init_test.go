package gotd

import (
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	redirectURL      = "http://localhost:8090"
	clientsecretPath = "./instance/client_secret.json"
)

// account setting please reference to init_test_secret.go.sample

var (
	// if true, test online
	onlineTest = true
	tdReal     *Service
)

func init() {
	if !onlineTest {
		return
	}

	auth := NewAuth()
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")

	var err error
	tdReal, err = New(client)
	if err != nil {
		panic(err)
	}
}

type TestTransport struct {
	body       string
	statusCode int
}

// RoundTrip add apikey
func (t *TestTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var res http.Response
	res.StatusCode = t.statusCode
	res.Body = ioutil.NopCloser(strings.NewReader(t.body))
	res.Header = http.Header{}
	res.Request = req

	return &res, nil
}

func clientTest(body string, statuscode int) *http.Client {
	transport := &TestTransport{body: body, statusCode: statuscode}

	client := &http.Client{
		Transport: transport,
	}

	return client
}
