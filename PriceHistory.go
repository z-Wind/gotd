package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

// NewPriceHistoryService https://developer.tdameritrade.com/price-history/apis
// Historical price data for charts
func NewPriceHistoryService(s *Service) *PriceHistoryService {
	rs := &PriceHistoryService{s: s}
	return rs
}

// PriceHistoryService https://developer.tdameritrade.com/price-history/apis
// Historical price data for charts
type PriceHistoryService struct {
	s *Service
}

// GetPriceHistory https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// Get price history for a symbol
func (r *PriceHistoryService) GetPriceHistory(symbol string) *PriceHistoryGetPriceHistoryCall {
	c := &PriceHistoryGetPriceHistoryCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},

		symbol: symbol,
	}

	return c
}

// PriceHistoryGetPriceHistoryCall https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// Get price history for a symbol
type PriceHistoryGetPriceHistoryCall struct {
	DefaultCall

	symbol string
}

// PeriodType https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// The type of period to show. Valid values are day, month, year, or ytd (year to date). Default is day.
func (c *PriceHistoryGetPriceHistoryCall) PeriodType(periodType string) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("periodType", periodType)
	return c
}

// Period https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// The number of periods to show.
//
// Example: For a 2 day / 1 min chart, the values would be:
//
// period: 2
// periodType: day
// frequency: 1
// frequencyType: min
//
// Valid periods by periodType (defaults marked with an asterisk):
//
// day: 1, 2, 3, 4, 5, 10*
// month: 1*, 2, 3, 6
// year: 1*, 2, 3, 5, 10, 15, 20
// ytd: 1*
func (c *PriceHistoryGetPriceHistoryCall) Period(period int) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("period", strconv.Itoa(period))
	return c
}

// FrequencyType https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// The type of frequency with which a new candle is formed.
//
// Valid frequencyTypes by periodType (defaults marked with an asterisk):
//
// day: minute*
// month: daily, weekly*
// year: daily, weekly, monthly*
// ytd: daily, weekly*
func (c *PriceHistoryGetPriceHistoryCall) FrequencyType(frequencyType string) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("frequencyType", frequencyType)
	return c
}

// Frequency https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// The number of the frequencyType to be included in each candle.
//
// Valid frequencies by frequencyType (defaults marked with an asterisk):
//
// minute: 1*, 5, 10, 15, 30
// daily: 1*
// weekly: 1*
// monthly: 1*
func (c *PriceHistoryGetPriceHistoryCall) Frequency(frequency int) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("frequency", strconv.Itoa(frequency))
	return c
}

// EndDate https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// End date as milliseconds since epoch. If startDate and endDate are provided, period should not be provided. Default is previous trading day.
func (c *PriceHistoryGetPriceHistoryCall) EndDate(date time.Time) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("endDate", strconv.FormatInt(date.UnixNano()/int64(time.Millisecond), 10))
	return c
}

// StartDate https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// Start date as milliseconds since epoch. If startDate and endDate are provided, period should not be provided.
func (c *PriceHistoryGetPriceHistoryCall) StartDate(date time.Time) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("startDate", strconv.FormatInt(date.UnixNano()/int64(time.Millisecond), 10))
	return c
}

// NeedExtendedHoursData https://developer.tdameritrade.com/price-history/apis/get/marketdata/%7Bsymbol%7D/pricehistory
// true to return extended hours data, false for regular market hours only. Default is true
func (c *PriceHistoryGetPriceHistoryCall) NeedExtendedHoursData(needExtendedHoursData bool) *PriceHistoryGetPriceHistoryCall {
	c.urlParams.Set("needExtendedHoursData", strconv.FormatBool(needExtendedHoursData))
	return c
}

func (c *PriceHistoryGetPriceHistoryCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata", c.symbol, "pricehistory")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *PriceHistoryGetPriceHistoryCall) Do() (*PriceHistory, error) {
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

	ret := &PriceHistory{
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
