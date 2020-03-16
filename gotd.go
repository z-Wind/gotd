package gotd

import (
	"net/http"

	"github.com/pkg/errors"
)

const (
	// Version defines the gax version being used. This is typically sent
	// in an HTTP header to services.
	Version = "0.5"

	// UserAgent is the header string used to identify this package.
	UserAgent = "td-api-go-client/" + Version

	basePath = "https://api.tdameritrade.com/v1"
)

// Service TD api Service
type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService
	Quotes   *QuotesService
}

// New TDAmeritrade API Server
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)

	return s, nil
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return UserAgent
	}
	return UserAgent + " " + s.UserAgent
}
