package gotd

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewInstrumentsService https://developer.tdameritrade.com/instruments/apis
// Search for instrument and fundamental data
func NewInstrumentsService(s *Service) *InstrumentsService {
	rs := &InstrumentsService{s: s}
	return rs
}

// InstrumentsService https://developer.tdameritrade.com/instruments/apis
// Search for instrument and fundamental data
type InstrumentsService struct {
	s *Service
}

// SearchInstruments https://developer.tdameritrade.com/instruments/apis/get/instruments
// Search or retrieve instrument data, including fundamental data.
func (r *InstrumentsService) SearchInstruments(symbol, projection string) *InstrumentsSearchInstrumentsCall {
	c := &InstrumentsSearchInstrumentsCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("symbol", symbol)
	c.urlParams.Set("projection", projection)
	return c
}

// InstrumentsSearchInstrumentsCall https://developer.tdameritrade.com/instruments/apis/get/instruments
// Search or retrieve instrument data, including fundamental data.
type InstrumentsSearchInstrumentsCall struct {
	DefaultCall
}

func (c *InstrumentsSearchInstrumentsCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "instruments")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *InstrumentsSearchInstrumentsCall) Do() (*InstrumentMap, error) {
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

	ret := &InstrumentMap{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new(map[string]*Instrument)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.Instruments = *target
	return ret, nil
}

// GetInstrument https://developer.tdameritrade.com/instruments/apis/get/instruments/%7Bcusip%7D
// Get an instrument by CUSIP
func (r *InstrumentsService) GetInstrument(cusip string) *InstrumentsGetInstrumentCall {
	c := &InstrumentsGetInstrumentCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
		cusip: cusip,
	}

	return c
}

// InstrumentsGetInstrumentCall https://developer.tdameritrade.com/instruments/apis/get/instruments/%7Bcusip%7D
// Get an instrument by CUSIP
type InstrumentsGetInstrumentCall struct {
	DefaultCall

	cusip string
}

func (c *InstrumentsGetInstrumentCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "instruments", c.cusip)
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *InstrumentsGetInstrumentCall) Do() (*Instrument, error) {
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

	target := new([]*Instrument)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	if len(*target) == 0 {
		return nil, fmt.Errorf("%s could not be found", c.cusip)
	}

	ret := (*target)[0]
	ret.ServerResponse = ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}
