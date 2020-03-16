package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// NewUserPrincipalService https://developer.tdameritrade.com/user-principal/apis
// APIs to access user-authorized accounts and their preferences
func NewUserPrincipalService(s *Service) *UserPrincipalService {
	rs := &UserPrincipalService{s: s}
	return rs
}

// UserPrincipalService https://developer.tdameritrade.com/user-principal/apis
// APIs to access user-authorized accounts and their preferences
type UserPrincipalService struct {
	s *Service
}

// GetPreferences https://developer.tdameritrade.com/user-principal/apis/get/accounts/%7BaccountId%7D/preferences-0
// Preferences for a specific account.
func (r *UserPrincipalService) GetPreferences(accountID string) *UserPrincipalGetPreferencesCall {
	c := &UserPrincipalGetPreferencesCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// UserPrincipalGetPreferencesCall https://developer.tdameritrade.com/user-principal/apis/get/accounts/%7BaccountId%7D/preferences-0
// Preferences for a specific account.
type UserPrincipalGetPreferencesCall struct {
	DefaultCall

	accountID string
}

func (c *UserPrincipalGetPreferencesCall) doRequest() (*http.Response, error) {
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
func (c *UserPrincipalGetPreferencesCall) Do() (*Preferences, error) {
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
func (r *UserPrincipalService) GetStreamerSubscriptionKeys(accountIDs ...string) *UserPrincipalGetStreamerSubscriptionKeysCall {
	c := &UserPrincipalGetStreamerSubscriptionKeysCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}

	c.urlParams.Set("accountIds", strings.Join(accountIDs, ","))
	return c
}

// UserPrincipalGetStreamerSubscriptionKeysCall https://developer.tdameritrade.com/user-principal/apis/get/userprincipals/streamersubscriptionkeys-0
// SubscriptionKey for provided accounts or default accounts.
type UserPrincipalGetStreamerSubscriptionKeysCall struct {
	DefaultCall
}

func (c *UserPrincipalGetStreamerSubscriptionKeysCall) doRequest() (*http.Response, error) {
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
func (c *UserPrincipalGetStreamerSubscriptionKeysCall) Do() (*SubscriptionKey, error) {
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
func (r *UserPrincipalService) GetUserPrincipals() *UserPrincipalGetUserPrincipalsCall {
	c := &UserPrincipalGetUserPrincipalsCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	return c
}

// UserPrincipalGetUserPrincipalsCall https://developer.tdameritrade.com/user-principal/apis/get/userprincipals-0
// User Principal details.
type UserPrincipalGetUserPrincipalsCall struct {
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
func (c *UserPrincipalGetUserPrincipalsCall) Fields(fields ...string) *UserPrincipalGetUserPrincipalsCall {
	c.urlParams.Set("fields", strings.Join(fields, ","))

	return c
}

func (c *UserPrincipalGetUserPrincipalsCall) doRequest() (*http.Response, error) {
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
func (c *UserPrincipalGetUserPrincipalsCall) Do() (*UserPrincipal, error) {
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
func (r *UserPrincipalService) UpdatePreferences(accountID string, preferences *Preferences) *UserPrincipalUpdatePreferencesCall {
	c := &UserPrincipalUpdatePreferencesCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		preferences: preferences,
	}
	return c
}

// UserPrincipalUpdatePreferencesCall https://developer.tdameritrade.com/user-principal/apis/put/accounts/%7BaccountId%7D/preferences-0
// Update preferences for a specific account.
// Please note that the directOptionsRouting and directEquityRouting values cannot be modified via this operation.
type UserPrincipalUpdatePreferencesCall struct {
	DefaultCall

	accountID   string
	preferences *Preferences
}

func (c *UserPrincipalUpdatePreferencesCall) doRequest() (*http.Response, error) {
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
func (c *UserPrincipalUpdatePreferencesCall) Do() (*UserPrincipal, error) {
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
