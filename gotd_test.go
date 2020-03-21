package gotd

import (
	"fmt"
	"log"
	"testing"
)

func TestNewServer(t *testing.T) {
	auth := NewAuth(redirectURL)
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
	td, err := New(client)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", td)
}

func ExampleAuth_GetClient() {
	auth := NewAuth("http://localhost:8090")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
	_, err := New(client)
	if err != nil {
		log.Fatal(err)
	}
}
func ExampleAuth_GetClient_tls() {
	auth := NewAuth("https://localhost:8090")
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
	_, err := New(client)
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleQuotesService_GetQuote() {
	auth := NewAuth("https://localhost:8090")
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")
	td, err := New(client)
	if err != nil {
		log.Fatal(err)
	}

	call := td.Quotes.GetQuote("VTI")
	quote, err := call.Do()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", quote.Symbol)

	// Output:
	// VTI
}
