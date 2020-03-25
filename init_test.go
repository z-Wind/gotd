package gotd

const (
	redirectURL      = "http://localhost:8090"
	clientsecretPath = "./instance/client_secret.json"
)

var td *Service

func init() {
	auth := NewAuth()
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath, "TDAmeritrade-go.json")

	var err error
	td, err = New(client)
	if err != nil {
		panic(err)
	}
}
