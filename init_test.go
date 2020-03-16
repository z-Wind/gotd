package gotd

const (
	redirectURL      = "http://localhost:8090"
	clientsecretPath = "client_secret.json"
)

var td *Service

func init() {
	auth := NewAuth(redirectURL)
	auth.SetTLS("./instance/cert.pem", "./instance/key.pem")
	client := auth.GetClient(clientsecretPath)

	var err error
	td, err = New(client)
	if err != nil {
		panic(err)
	}
}
