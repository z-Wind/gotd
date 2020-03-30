package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strconv"

	"github.com/pkg/errors"
)

// NewSavedOrdersService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
func NewSavedOrdersService(s *Service) *SavedOrdersService {
	rs := &SavedOrdersService{s: s}
	return rs
}

// SavedOrdersService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
type SavedOrdersService struct {
	s *Service
}

// ReplaceSavedOrder https://developer.tdameritrade.com/account-access/apis/put/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Replace an existing saved order for an account. The existing saved order will be replaced by the new order.
func (r *SavedOrdersService) ReplaceSavedOrder(accountID string, savedOrderID int64, savedOrder *SavedOrder) *SavedOrdersReplaceSavedOrderCall {
	c := &SavedOrdersReplaceSavedOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:    accountID,
		savedOrderID: savedOrderID,
		savedOrder:   savedOrder,
	}
	return c
}

// SavedOrdersReplaceSavedOrderCall https://developer.tdameritrade.com/account-access/apis/put/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Replace an existing saved order for an account. The existing saved order will be replaced by the new order.
type SavedOrdersReplaceSavedOrderCall struct {
	DefaultCall

	accountID    string
	savedOrderID int64
	savedOrder   *SavedOrder
}

func (c *SavedOrdersReplaceSavedOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.savedOrder)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "savedorders", strconv.FormatInt(c.savedOrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *SavedOrdersReplaceSavedOrderCall) Do() (*ServerResponse, error) {
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

// GetSavedOrdersByPath https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/savedorders-0
// Saved orders for a specific account.
func (r *SavedOrdersService) GetSavedOrdersByPath(accountID string) *SavedOrdersGetSavedOrdersByPathCall {
	c := &SavedOrdersGetSavedOrdersByPathCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// SavedOrdersGetSavedOrdersByPathCall https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/savedorders-0
// Saved orders for a specific account.
type SavedOrdersGetSavedOrdersByPathCall struct {
	DefaultCall

	accountID string
}

func (c *SavedOrdersGetSavedOrdersByPathCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "savedorders")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *SavedOrdersGetSavedOrdersByPathCall) Do() (*SavedOrderList, error) {
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

	ret := &SavedOrderList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*SavedOrder)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.SavedOrders = *target
	return ret, nil
}

// GetSavedOrder https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Specific saved order by its ID, for a specific account.
func (r *SavedOrdersService) GetSavedOrder(accountID string, savedOrderID int64) *SavedOrdersGetSavedOrderCall {
	c := &SavedOrdersGetSavedOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:    accountID,
		savedOrderID: savedOrderID,
	}
	return c
}

// SavedOrdersGetSavedOrderCall https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Specific saved order by its ID, for a specific account.
type SavedOrdersGetSavedOrderCall struct {
	DefaultCall

	accountID    string
	savedOrderID int64
}

func (c *SavedOrdersGetSavedOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "savedorders", strconv.FormatInt(c.savedOrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *SavedOrdersGetSavedOrderCall) Do() (*SavedOrder, error) {
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

	ret := &SavedOrder{
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

// DeleteSavedOrder https://developer.tdameritrade.com/account-access/apis/delete/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Delete a specific saved order for a specific account.
func (r *SavedOrdersService) DeleteSavedOrder(accountID string, savedOrderID int64) *SavedOrdersDeleteSavedOrderCall {
	c := &SavedOrdersDeleteSavedOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:    accountID,
		savedOrderID: savedOrderID,
	}
	return c
}

// SavedOrdersDeleteSavedOrderCall https://developer.tdameritrade.com/account-access/apis/delete/accounts/%7BaccountId%7D/savedorders/%7BsavedOrderId%7D-0
// Delete a specific saved order for a specific account.
type SavedOrdersDeleteSavedOrderCall struct {
	DefaultCall

	accountID    string
	savedOrderID int64
}

func (c *SavedOrdersDeleteSavedOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "savedorders", strconv.FormatInt(c.savedOrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *SavedOrdersDeleteSavedOrderCall) Do() (*ServerResponse, error) {
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

// CreateSavedOrder https://developer.tdameritrade.com/account-access/apis/post/accounts/%7BaccountId%7D/savedorders-0
// Save an order for a specific account.
func (r *SavedOrdersService) CreateSavedOrder(accountID string, savedOrder *SavedOrder) *SavedOrdersCreateSavedOrderCall {
	c := &SavedOrdersCreateSavedOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:  accountID,
		savedOrder: savedOrder,
	}
	return c
}

// SavedOrdersCreateSavedOrderCall https://developer.tdameritrade.com/account-access/apis/post/accounts/%7BaccountId%7D/savedorders-0
// Save an order for a specific account.
type SavedOrdersCreateSavedOrderCall struct {
	DefaultCall

	accountID  string
	savedOrder *SavedOrder
}

func (c *SavedOrdersCreateSavedOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.savedOrder)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "savedorders")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *SavedOrdersCreateSavedOrderCall) Do() (*ServerResponse, error) {
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
