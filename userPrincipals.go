package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// NewUserPrincipalsService https://developer.tdameritrade.com/user-principal/apis
// APIs to access user-authorized accounts and their preferences
func NewUserPrincipalsService(s *Service) *UserPrincipalsService {
	rs := &UserPrincipalsService{s: s}
	return rs
}

// UserPrincipalsService https://developer.tdameritrade.com/user-principal/apis
// APIs to access user-authorized accounts and their preferences
type UserPrincipalsService struct {
	s *Service
}

// GetPreferences https://developer.tdameritrade.com/user-principal/apis/get/accounts/%7BaccountId%7D/preferences-0
// Preferences for a specific account.
func (r *UserPrincipalsService) GetPreferences(accountID string) *UserPrincipalsGetPreferencesCall {
	c := &UserPrincipalsGetPreferencesCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// UserPrincipalsGetPreferencesCall https://developer.tdameritrade.com/user-principal/apis/get/accounts/%7BaccountId%7D/preferences-0
// Preferences for a specific account.
type UserPrincipalsGetPreferencesCall struct {
	DefaultCall

	accountID string
}

func (c *UserPrincipalsGetPreferencesCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "preferences")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *UserPrincipalsGetPreferencesCall) Do() (*Preferences, error) {
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

	ret := &Preferences{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	return ret, nil
}

// GetStreamerSubscriptionKeys https://developer.tdameritrade.com/user-principal/apis/get/userprincipals/streamersubscriptionkeys-0
// SubscriptionKey for provided accounts or default accounts.
func (r *UserPrincipalsService) GetStreamerSubscriptionKeys(accountIDs ...string) *UserPrincipalsGetStreamerSubscriptionKeysCall {
	c := &UserPrincipalsGetStreamerSubscriptionKeysCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}

	c.urlParams.Set("accountIds", strings.Join(accountIDs, ","))
	return c
}

// UserPrincipalsGetStreamerSubscriptionKeysCall https://developer.tdameritrade.com/user-principal/apis/get/userprincipals/streamersubscriptionkeys-0
// SubscriptionKey for provided accounts or default accounts.
type UserPrincipalsGetStreamerSubscriptionKeysCall struct {
	DefaultCall
}

func (c *UserPrincipalsGetStreamerSubscriptionKeysCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "userprincipals", "streamersubscriptionkeys")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *UserPrincipalsGetStreamerSubscriptionKeysCall) Do() (*SubscriptionKey, error) {
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

	ret := &SubscriptionKey{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	return ret, nil
}

// GetUserPrincipals https://developer.tdameritrade.com/user-principal/apis/get/userprincipals-0
// User Principal details.
func (r *UserPrincipalsService) GetUserPrincipals() *UserPrincipalsGetUserPrincipalsCall {
	c := &UserPrincipalsGetUserPrincipalsCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	return c
}

// UserPrincipalsGetUserPrincipalsCall https://developer.tdameritrade.com/user-principal/apis/get/userprincipals-0
// User Principal details.
type UserPrincipalsGetUserPrincipalsCall struct {
	DefaultCall
}

// Fields https://developer.tdameritrade.com/user-principal/apis/get/userprincipals-0
// A comma separated String which allows one to specify additional fields to return. None of these fields are returned by default. Possible values in this String can be:
//
// streamerSubscriptionKeys
// streamerConnectionInfo
// preferences
// surrogateIds
//
// Example:
// fields=streamerSubscriptionKeys,streamerConnectionInfo
func (c *UserPrincipalsGetUserPrincipalsCall) Fields(fields ...string) *UserPrincipalsGetUserPrincipalsCall {
	c.urlParams.Set("fields", strings.Join(fields, ","))

	return c
}

func (c *UserPrincipalsGetUserPrincipalsCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "userprincipals")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *UserPrincipalsGetUserPrincipalsCall) Do() (*UserPrincipal, error) {
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

	ret := &UserPrincipal{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	return ret, nil
}

// UpdatePreferences https://developer.tdameritrade.com/user-principal/apis/put/accounts/%7BaccountId%7D/preferences-0
// Update preferences for a specific account.
// Please note that the directOptionsRouting and directEquityRouting values cannot be modified via this operation.
func (r *UserPrincipalsService) UpdatePreferences(accountID string, preferences *Preferences) *UserPrincipalsUpdatePreferencesCall {
	c := &UserPrincipalsUpdatePreferencesCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		preferences: preferences,
	}
	return c
}

// UserPrincipalsUpdatePreferencesCall https://developer.tdameritrade.com/user-principal/apis/put/accounts/%7BaccountId%7D/preferences-0
// Update preferences for a specific account.
// Please note that the directOptionsRouting and directEquityRouting values cannot be modified via this operation.
type UserPrincipalsUpdatePreferencesCall struct {
	DefaultCall

	accountID   string
	preferences *Preferences
}

func (c *UserPrincipalsUpdatePreferencesCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.preferences)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "preferences")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *UserPrincipalsUpdatePreferencesCall) Do() (*ServerResponse, error) {
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

	ServerResponse := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ServerResponse, nil
}
