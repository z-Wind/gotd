package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// NewAccountsService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

// AccountsService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
type AccountsService struct {
	s *Service
}

// GetAccount https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D-0
// Account balances, positions, and orders for a specific account.
func (r *AccountsService) GetAccount(accountID string) *AccountsGetAccountCall {
	c := &AccountsGetAccountCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// AccountsGetAccountCall https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D-0
// Account balances, positions, and orders for a specific account.
type AccountsGetAccountCall struct {
	DefaultCall

	accountID string
}

// Fields https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D-0
// Balances displayed by default, additional fields can be added here by adding positions or orders
// Example:
// fields=positions,orders
func (c *AccountsGetAccountCall) Fields(s ...string) *AccountsGetAccountCall {
	c.urlParams.Set("fields", strings.Join(s, ","))
	return c
}

func (c *AccountsGetAccountCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *AccountsGetAccountCall) Do() (*Account, error) {
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

	ret := &Account{
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

// GetAccountList https://developer.tdameritrade.com/account-access/apis/get/accounts-0
// Account balances, positions, and orders for all linked accounts.
func (r *AccountsService) GetAccountList() *AccountsGetAccountListCall {
	c := &AccountsGetAccountListCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	return c
}

// AccountsGetAccountListCall https://developer.tdameritrade.com/account-access/apis/get/accounts-0
// Account balances, positions, and orders for all linked accounts.
type AccountsGetAccountListCall struct {
	DefaultCall
}

// Fields https://developer.tdameritrade.com/account-access/apis/get/accounts-0
// Balances displayed by default, additional fields can be added here by adding positions or orders
// Example:
// fields=positions,orders
func (c *AccountsGetAccountListCall) Fields(s ...string) *AccountsGetAccountListCall {
	c.urlParams.Set("fields", strings.Join(s, ","))
	return c
}

func (c *AccountsGetAccountListCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *AccountsGetAccountListCall) Do() (*AccountList, error) {
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
		return nil, err
	}
	defer res.Body.Close()
	if err := CheckResponse(res); err != nil {
		return nil, errors.Wrapf(err, "CheckResponse")
	}

	ret := &AccountList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Account)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Accounts = *target
	return ret, nil

}
