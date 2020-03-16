package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// NewOrdersService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
func NewOrdersService(s *Service) *OrdersService {
	rs := &OrdersService{s: s}
	return rs
}

// OrdersService https://developer.tdameritrade.com/account-access/apis
// APIs to access Account Balances, Positions, Trade Info and place Trades
type OrdersService struct {
	s *Service
}

// ReplaceOrder https://developer.tdameritrade.com/account-access/apis/put/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Replace an existing order for an account. The existing order will be replaced by the new order. Once replaced, the old order will be canceled and a new order will be created.
func (r *OrdersService) ReplaceOrder(accountID string, OrderID int64, Order *Order) *OrdersReplaceOrderCall {
	c := &OrdersReplaceOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
		OrderID:   OrderID,
		Order:     Order,
	}
	return c
}

// OrdersReplaceOrderCall https://developer.tdameritrade.com/account-access/apis/put/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Replace an existing order for an account. The existing order will be replaced by the new order. Once replaced, the old order will be canceled and a new order will be created.
type OrdersReplaceOrderCall struct {
	DefaultCall

	accountID string
	OrderID   int64
	Order     *Order
}

func (c *OrdersReplaceOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.Order)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "orders", strconv.FormatInt(c.OrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersReplaceOrderCall) Do() (*ServerResponse, error) {
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

	serverResponse := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return serverResponse, nil
}

// GetOrdersByQuery https://developer.tdameritrade.com/account-access/apis/get/orders-0
// All orders for a specific account or, if account ID isn't specified, orders will be returned for all linked accounts
func (r *OrdersService) GetOrdersByQuery(accountID string) *OrdersGetOrdersByQueryCall {
	c := &OrdersGetOrdersByQueryCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("accountId", accountID)
	return c
}

// OrdersGetOrdersByQueryCall https://developer.tdameritrade.com/account-access/apis/get/orders-0
// All orders for a specific account or, if account ID isn't specified, orders will be returned for all linked accounts
type OrdersGetOrdersByQueryCall struct {
	DefaultCall
}

// MaxResults https://developer.tdameritrade.com/account-access/apis/get/orders-0
// The maximum number of orders to retrieve.
func (c *OrdersGetOrdersByQueryCall) MaxResults(n int) *OrdersGetOrdersByQueryCall {
	c.urlParams.Set("maxResults", strconv.Itoa(n))

	return c
}

// FromEnteredTime https://developer.tdameritrade.com/account-access/apis/get/orders-0
// Specifies that no orders entered before this time should be returned.Valid ISO-8601 formats are :
// yyyy-MM-dd. Date must be within 60 days from today's date. 'toEnteredTime' must also be set.
func (c *OrdersGetOrdersByQueryCall) FromEnteredTime(date time.Time) *OrdersGetOrdersByQueryCall {
	c.urlParams.Set("fromEnteredTime", date.UTC().Format("2006-01-02"))

	return c
}

// ToEnteredTime https://developer.tdameritrade.com/account-access/apis/get/orders-0
// Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :
// yyyy-MM-dd. 'fromEnteredTime' must also be set.
func (c *OrdersGetOrdersByQueryCall) ToEnteredTime(date time.Time) *OrdersGetOrdersByQueryCall {
	c.urlParams.Set("toEnteredTime", date.UTC().Format("2006-01-02"))

	return c
}

// Status https://developer.tdameritrade.com/account-access/apis/get/orders-0
// Specifies that only orders of this status should be returned.
func (c *OrdersGetOrdersByQueryCall) Status(s string) *OrdersGetOrdersByQueryCall {
	c.urlParams.Set("status", s)

	return c
}

func (c *OrdersGetOrdersByQueryCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "orders")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersGetOrdersByQueryCall) Do() (*OrderList, error) {
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

	ret := &OrderList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Order)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Orders = *target
	return ret, nil
}

// GetOrdersByPath https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// Orders for a specific account.
func (r *OrdersService) GetOrdersByPath(accountID string) *OrdersGetOrdersByPathCall {
	c := &OrdersGetOrdersByPathCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// OrdersGetOrdersByPathCall https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// Orders for a specific account.
type OrdersGetOrdersByPathCall struct {
	DefaultCall

	accountID string
}

// MaxResults https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// The maximum number of orders to retrieve.
func (c *OrdersGetOrdersByPathCall) MaxResults(n int) *OrdersGetOrdersByPathCall {
	c.urlParams.Set("maxResults", strconv.Itoa(n))

	return c
}

// FromEnteredTime https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are :
// yyyy-MM-dd. Date must be within 60 days from today's date. 'toEnteredTime' must also be set.
func (c *OrdersGetOrdersByPathCall) FromEnteredTime(date time.Time) *OrdersGetOrdersByPathCall {
	c.urlParams.Set("fromEnteredTime", date.UTC().Format("2006-01-02"))

	return c
}

// ToEnteredTime https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are :
// yyyy-MM-dd. 'fromEnteredTime' must also be set.
func (c *OrdersGetOrdersByPathCall) ToEnteredTime(date time.Time) *OrdersGetOrdersByPathCall {
	c.urlParams.Set("toEnteredTime", date.UTC().Format("2006-01-02"))

	return c
}

// Status https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders-0
// Specifies that only orders of this status should be returned.
func (c *OrdersGetOrdersByPathCall) Status(s string) *OrdersGetOrdersByPathCall {
	c.urlParams.Set("status", s)

	return c
}

func (c *OrdersGetOrdersByPathCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "orders")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersGetOrdersByPathCall) Do() (*OrderList, error) {
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

	ret := &OrderList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Order)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Orders = *target
	return ret, nil
}

// GetOrder https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Get a specific order for a specific account.
func (r *OrdersService) GetOrder(accountID string, OrderID int64) *OrdersGetOrderCall {
	c := &OrdersGetOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
		OrderID:   OrderID,
	}
	return c
}

// OrdersGetOrderCall https://developer.tdameritrade.com/account-access/apis/get/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Get a specific order for a specific account.
type OrdersGetOrderCall struct {
	DefaultCall

	accountID string
	OrderID   int64
}

func (c *OrdersGetOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "orders", strconv.FormatInt(c.OrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersGetOrderCall) Do() (*Order, error) {
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

	ret := &Order{
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

// CancelOrder https://developer.tdameritrade.com/account-access/apis/delete/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Cancel a specific order for a specific account.
func (r *OrdersService) CancelOrder(accountID string, OrderID int64) *OrdersCancelOrderCall {
	c := &OrdersCancelOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
		OrderID:   OrderID,
	}
	return c
}

// OrdersCancelOrderCall https://developer.tdameritrade.com/account-access/apis/delete/accounts/%7BaccountId%7D/orders/%7BorderId%7D-0
// Cancel a specific order for a specific account.
type OrdersCancelOrderCall struct {
	DefaultCall

	accountID string
	OrderID   int64
}

func (c *OrdersCancelOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "orders", strconv.FormatInt(c.OrderID, 10))
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersCancelOrderCall) Do() (*ServerResponse, error) {
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

// PlaceOrder https://developer.tdameritrade.com/account-access/apis/post/accounts/%7BaccountId%7D/orders-0
// Place an order for a specific account.
func (r *OrdersService) PlaceOrder(accountID string, Order *Order) *OrdersPlaceOrderCall {
	c := &OrdersPlaceOrderCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
		Order:     Order,
	}
	return c
}

// OrdersPlaceOrderCall https://developer.tdameritrade.com/account-access/apis/post/accounts/%7BaccountId%7D/orders-0
// Place an order for a specific account.
type OrdersPlaceOrderCall struct {
	DefaultCall

	accountID string
	Order     *Order
}

func (c *OrdersPlaceOrderCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.Order)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "orders")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OrdersPlaceOrderCall) Do() (*ServerResponse, error) {
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
