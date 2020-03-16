package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// NewQuotesService https://developer.tdameritrade.com/quotes/apis
// Request real-time and delayed top level quote data
func NewQuotesService(s *Service) *QuotesService {
	rs := &QuotesService{s: s}
	return rs
}

// QuotesService https://developer.tdameritrade.com/quotes/apis
// Request real-time and delayed top level quote data
type QuotesService struct {
	s *Service
}

// GetQuote https://developer.tdameritrade.com/quotes/apis/get/marketdata/%7Bsymbol%7D/quotes
// Get quote for a symbol
func (r *QuotesService) GetQuote(symbol string) *QuotesGetQuoteCall {
	c := &QuotesGetQuoteCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}
	return c
}

// QuotesGetQuoteCall https://developer.tdameritrade.com/quotes/apis/get/marketdata/%7Bsymbol%7D/quotes
// Get quote for a symbol
type QuotesGetQuoteCall struct {
	DefaultCall

	symbol string
}

func (c *QuotesGetQuoteCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata", c.symbol, "quotes")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *QuotesGetQuoteCall) Do() (*Quote, error) {
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

	target := new(map[string]*Quote)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	ret := (*target)[c.symbol]
	ret.ServerResponse = ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}
	return ret, nil

}

// GetQuoteList https://developer.tdameritrade.com/quotes/apis/get/marketdata/quotes
// Get quote for one or more symbols
func (r *QuotesService) GetQuoteList(symbol ...string) *QuotesGetQuoteListCall {
	c := &QuotesGetQuoteListCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("symbol", strings.Join(symbol, ","))
	return c
}

// QuotesGetQuoteListCall https://developer.tdameritrade.com/quotes/apis/get/marketdata/quotes
// Get quote for one or more symbols
type QuotesGetQuoteListCall struct {
	DefaultCall
}

func (c *QuotesGetQuoteListCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata/quotes")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *QuotesGetQuoteListCall) Do() (*QuoteMap, error) {
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

	ret := &QuoteMap{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new(map[string]*Quote)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Quotes = *target
	return ret, nil

}
