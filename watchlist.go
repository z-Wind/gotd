package gotd

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewWatchlistService https://developer.tdameritrade.com/watchlist/apis
// APIs to perform CRUD operations on Account Watchlist
func NewWatchlistService(s *Service) *WatchlistService {
	rs := &WatchlistService{s: s}
	return rs
}

// WatchlistService https://developer.tdameritrade.com/watchlist/apis
// APIs to perform CRUD operations on Account Watchlist
type WatchlistService struct {
	s *Service
}

// CreateWatchlist https://developer.tdameritrade.com/watchlist/apis/post/accounts/%7BaccountId%7D/watchlists-0
// Create watchlist for specific account.This method does not verify that the symbol or asset type are valid.
func (r *WatchlistService) CreateWatchlist(accountID string, watchlist *Watchlist) *WatchlistCreateWatchlistCall {
	c := &WatchlistCreateWatchlistCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
		watchlist: watchlist,
	}
	return c
}

// WatchlistCreateWatchlistCall https://developer.tdameritrade.com/watchlist/apis/post/accounts/%7BaccountId%7D/watchlists-0
// Create watchlist for specific account.This method does not verify that the symbol or asset type are valid.
type WatchlistCreateWatchlistCall struct {
	DefaultCall

	accountID string
	watchlist *Watchlist
}

func (c *WatchlistCreateWatchlistCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.watchlist)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("POST", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistCreateWatchlistCall) Do() (*ServerResponse, error) {
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

	ret := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}

// DeleteWatchlist https://developer.tdameritrade.com/watchlist/apis/delete/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Delete watchlist for a specific account.
func (r *WatchlistService) DeleteWatchlist(accountID, watchlistID string) *WatchlistDeleteWatchlistCall {
	c := &WatchlistDeleteWatchlistCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		watchlistID: watchlistID,
	}
	return c
}

// WatchlistDeleteWatchlistCall https://developer.tdameritrade.com/watchlist/apis/delete/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Delete watchlist for a specific account.
type WatchlistDeleteWatchlistCall struct {
	DefaultCall

	accountID, watchlistID string
}

func (c *WatchlistDeleteWatchlistCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists", c.watchlistID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("DELETE", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistDeleteWatchlistCall) Do() (*ServerResponse, error) {
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

	ret := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}

// GetWatchlist https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Specific watchlist for a specific account.
func (r *WatchlistService) GetWatchlist(accountID, watchlistID string) *WatchlistGetWatchlistCall {
	c := &WatchlistGetWatchlistCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		watchlistID: watchlistID,
	}
	return c
}

// WatchlistGetWatchlistCall https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Specific watchlist for a specific account.
type WatchlistGetWatchlistCall struct {
	DefaultCall

	accountID, watchlistID string
}

func (c *WatchlistGetWatchlistCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists", c.watchlistID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistGetWatchlistCall) Do() (*Watchlist, error) {
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

	ret := &Watchlist{
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

// GetWatchlistsForMultipleAccounts https://developer.tdameritrade.com/watchlist/apis/get/accounts/watchlists-0
// All watchlists for all of the user's linked accounts.
func (r *WatchlistService) GetWatchlistsForMultipleAccounts() *WatchlistGetWatchlistsForMultipleAccountsCall {
	c := &WatchlistGetWatchlistsForMultipleAccountsCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	return c
}

// WatchlistGetWatchlistsForMultipleAccountsCall https://developer.tdameritrade.com/watchlist/apis/get/accounts/watchlists-0
// All watchlists for all of the user's linked accounts.
type WatchlistGetWatchlistsForMultipleAccountsCall struct {
	DefaultCall
}

func (c *WatchlistGetWatchlistsForMultipleAccountsCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", "watchlists")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistGetWatchlistsForMultipleAccountsCall) Do() (*WatchlistList, error) {
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

	ret := &WatchlistList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Watchlist)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	ret.Watchlists = *target

	return ret, nil
}

// GetWatchlistsForSingleAccount https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists-0
// All watchlists of an account.
func (r *WatchlistService) GetWatchlistsForSingleAccount(accountID string) *WatchlistGetWatchlistsForSingleAccountCall {
	c := &WatchlistGetWatchlistsForSingleAccountCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID: accountID,
	}
	return c
}

// WatchlistGetWatchlistsForSingleAccountCall https://developer.tdameritrade.com/watchlist/apis/get/accounts/%7BaccountId%7D/watchlists-0
// All watchlists of an account.
type WatchlistGetWatchlistsForSingleAccountCall struct {
	DefaultCall

	accountID string
}

func (c *WatchlistGetWatchlistsForSingleAccountCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistGetWatchlistsForSingleAccountCall) Do() (*WatchlistList, error) {
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

	ret := &WatchlistList{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new([]*Watchlist)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	ret.Watchlists = *target

	return ret, nil
}

// ReplaceWatchlist https://developer.tdameritrade.com/watchlist/apis/put/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Replace watchlist for a specific account.
// This method does not verify that the symbol or asset type are valid.
func (r *WatchlistService) ReplaceWatchlist(accountID, watchlistID string, watchlist *Watchlist) *WatchlistReplaceWatchlistCall {
	c := &WatchlistReplaceWatchlistCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		watchlistID: watchlistID,
		watchlist:   watchlist,
	}
	return c
}

// WatchlistReplaceWatchlistCall https://developer.tdameritrade.com/watchlist/apis/put/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Replace watchlist for a specific account.
// This method does not verify that the symbol or asset type are valid.
type WatchlistReplaceWatchlistCall struct {
	DefaultCall

	accountID, watchlistID string
	watchlist              *Watchlist
}

func (c *WatchlistReplaceWatchlistCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.watchlist)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists", c.watchlistID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("PUT", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistReplaceWatchlistCall) Do() (*ServerResponse, error) {
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

	ret := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}

// UpdateWatchlist https://developer.tdameritrade.com/watchlist/apis/patch/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Partially update watchlist for a specific account: change watchlist name, add to the beginning/end of a watchlist, update or delete items in a watchlist.
// This method does not verify that the symbol or asset type are valid.
func (r *WatchlistService) UpdateWatchlist(accountID, watchlistID string, watchlist *Watchlist) *WatchlistUpdateWatchlistCall {
	c := &WatchlistUpdateWatchlistCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		accountID:   accountID,
		watchlistID: watchlistID,
		watchlist:   watchlist,
	}
	return c
}

// WatchlistUpdateWatchlistCall https://developer.tdameritrade.com/watchlist/apis/patch/accounts/%7BaccountId%7D/watchlists/%7BwatchlistId%7D-0
// Partially update watchlist for a specific account: change watchlist name, add to the beginning/end of a watchlist, update or delete items in a watchlist.
// This method does not verify that the symbol or asset type are valid.
type WatchlistUpdateWatchlistCall struct {
	DefaultCall

	accountID, watchlistID string
	watchlist              *Watchlist
}

func (c *WatchlistUpdateWatchlistCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	body, err := JSONReader(c.watchlist)
	if err != nil {
		return nil, errors.Wrapf(err, "JSONReader")
	}
	reqHeaders.Set("Content-Type", "application/json")

	urls := ResolveRelative(c.s.BasePath, "accounts", c.accountID, "watchlists", c.watchlistID)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *WatchlistUpdateWatchlistCall) Do() (*ServerResponse, error) {
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

	ret := &ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}
