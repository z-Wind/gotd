package gotd

import (
	"io"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// NewMoversService https://developer.tdameritrade.com/movers/apis
// Retrieve mover information by index symbol, direction type and change
func NewMoversService(s *Service) *MoversService {
	rs := &MoversService{s: s}
	return rs
}

// MoversService https://developer.tdameritrade.com/movers/apis
// Retrieve mover information by index symbol, direction type and change
type MoversService struct {
	s *Service
}

// GetMovers https://developer.tdameritrade.com/movers/apis/get/marketdata/%7Bindex%7D/movers
// Top 10 (up or down) movers by value or percent for a particular market
func (r *MoversService) GetMovers(index string) *MoversGetMoversCall {
	c := &MoversGetMoversCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		index: index,
	}
	return c
}

// MoversGetMoversCall https://developer.tdameritrade.com/movers/apis/get/marketdata/%7Bindex%7D/movers
// Top 10 (up or down) movers by value or percent for a particular market
type MoversGetMoversCall struct {
	DefaultCall

	index string
}

// Direction https://developer.tdameritrade.com/movers/apis/get/marketdata/%7Bindex%7D/movers
// To return movers with the specified directions of up or down
func (c *MoversGetMoversCall) Direction(direction string) *MoversGetMoversCall {
	c.urlParams.Set("direction", direction)
	return c
}

// Change https://developer.tdameritrade.com/movers/apis/get/marketdata/%7Bindex%7D/movers
// To return movers with the specified change types of percent or value
func (c *MoversGetMoversCall) Change(change string) *MoversGetMoversCall {
	c.urlParams.Set("change", change)
	return c
}

func (c *MoversGetMoversCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata", c.index, "movers")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *MoversGetMoversCall) Do() (*Mover, error) {
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

	target := new(map[string]*Mover)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret := (*target)[c.index]
	ret.ServerResponse = ServerResponse{
		Header:         res.Header,
		HTTPStatusCode: res.StatusCode,
	}

	return ret, nil
}
