package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// NewMarketHoursService https://developer.tdameritrade.com/market-hours/apis
// Operating hours of markets
func NewMarketHoursService(s *Service) *MarketHoursService {
	rs := &MarketHoursService{s: s}
	return rs
}

// MarketHoursService https://developer.tdameritrade.com/market-hours/apis
// Operating hours of markets
type MarketHoursService struct {
	s *Service
}

// GetHoursForMultipleMarkets https://developer.tdameritrade.com/market-hours/apis/get/marketdata/hours
// Retrieve market hours for specified markets
// The markets for which you're requesting market hours, comma-separated. Valid markets are EQUITY, OPTION, FUTURE, BOND, or FOREX.
func (r *MarketHoursService) GetHoursForMultipleMarkets(markets ...string) *MarketHoursGetHoursForMultipleMarketsCall {
	c := &MarketHoursGetHoursForMultipleMarketsCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}
	c.urlParams.Set("markets", strings.Join(markets, ","))

	return c
}

// MarketHoursGetHoursForMultipleMarketsCall https://developer.tdameritrade.com/market-hours/apis/get/marketdata/hours
// Retrieve market hours for specified markets
type MarketHoursGetHoursForMultipleMarketsCall struct {
	DefaultCall
}

// Date https://developer.tdameritrade.com/market-hours/apis/get/marketdata/hours
// "The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz."
func (c *MarketHoursGetHoursForMultipleMarketsCall) Date(date time.Time) *MarketHoursGetHoursForMultipleMarketsCall {
	c.urlParams.Set("date", date.UTC().Format("2006-01-02"))

	return c
}

func (c *MarketHoursGetHoursForMultipleMarketsCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata/hours")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *MarketHoursGetHoursForMultipleMarketsCall) Do() (*MarketHourMap, error) {
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

	ret := &MarketHourMap{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new(map[string]map[string]*MarketHour)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}
	ret.MarketHourProductMaps = *target

	return ret, nil
}

// GetHoursForASingleMarket https://developer.tdameritrade.com/market-hours/apis/get/marketdata/%7Bmarket%7D/hours
// Retrieve market hours for specified single market
func (r *MarketHoursService) GetHoursForASingleMarket(market string) *MarketHoursGetHoursForASingleMarketCall {
	c := &MarketHoursGetHoursForASingleMarketCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
		market: market,
	}

	return c
}

// MarketHoursGetHoursForASingleMarketCall https://developer.tdameritrade.com/market-hours/apis/get/marketdata/%7Bmarket%7D/hours
// Retrieve market hours for specified single market
type MarketHoursGetHoursForASingleMarketCall struct {
	DefaultCall

	market string
}

// Date https://developer.tdameritrade.com/market-hours/apis/get/marketdata/%7Bmarket%7D/hours
// "The date for which market hours information is requested. Valid ISO-8601 formats are : yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz."
func (c *MarketHoursGetHoursForASingleMarketCall) Date(date time.Time) *MarketHoursGetHoursForASingleMarketCall {
	c.urlParams.Set("date", date.UTC().Format("2006-01-02"))

	return c
}

func (c *MarketHoursGetHoursForASingleMarketCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata", c.market, "hours")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *MarketHoursGetHoursForASingleMarketCall) Do() (*MarketHourProductMap, error) {
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

	ret := &MarketHourProductMap{
		ServerResponse: ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := new(map[string]map[string]*MarketHour)
	if err := DecodeResponse(target, res); err != nil {
		return nil, errors.Wrapf(err, "DecodeResponse")
	}

	ret.MarketHours = (*target)[strings.ToLower(c.market)]

	return ret, nil
}
