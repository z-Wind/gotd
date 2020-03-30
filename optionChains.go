package gotd

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// NewOptionChainsService https://developer.tdameritrade.com/option-chains/apis
// Get Option Chains for optionable symbols
func NewOptionChainsService(s *Service) *OptionChainsService {
	rs := &OptionChainsService{s: s}
	return rs
}

// OptionChainsService https://developer.tdameritrade.com/option-chains/apis
// Get Option Chains for optionable symbols
type OptionChainsService struct {
	s *Service
}

// GetOptionChain https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Get option chain for an optionable Symbol
func (r *OptionChainsService) GetOptionChain(symbol string) *OptionChainsGetOptionChainCall {
	c := &OptionChainsGetOptionChainCall{
		DefaultCall: DefaultCall{
			s:         r.s,
			urlParams: url.Values{},
		},
	}

	c.urlParams.Set("symbol", symbol)
	return c
}

// OptionChainsGetOptionChainCall https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Get option chain for an optionable Symbol
type OptionChainsGetOptionChainCall struct {
	DefaultCall
}

// ContractType https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Type of contracts to return in the chain. Can be CALL, PUT, or ALL. Default is ALL.
func (c *OptionChainsGetOptionChainCall) ContractType(contractType string) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("contractType", contractType)
	return c
}

// StrikeCount https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// 	The number of strikes to return above and below the at-the-money price.
func (c *OptionChainsGetOptionChainCall) StrikeCount(strikeCount int) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("strikeCount", strconv.Itoa(strikeCount))
	return c
}

// IncludeQuotes https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Include quotes for options in the option chain. Can be TRUE or FALSE. Default is FALSE.
func (c *OptionChainsGetOptionChainCall) IncludeQuotes(includeQuotes bool) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("includeQuotes", strings.ToUpper(strconv.FormatBool(includeQuotes)))
	return c
}

// Strategy https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Passing a value returns a Strategy Chain. Possible values are SINGLE, ANALYTICAL (allows use of the volatility, underlyingPrice, interestRate, and daysToExpiration params to calculate theoretical values), COVERED, VERTICAL, CALENDAR, STRANGLE, STRADDLE, BUTTERFLY, CONDOR, DIAGONAL, COLLAR, or ROLL. Default is SINGLE.
func (c *OptionChainsGetOptionChainCall) Strategy(strategy string) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("strategy", strategy)
	return c
}

// Interval https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Strike interval for spread strategy chains (see strategy param).
func (c *OptionChainsGetOptionChainCall) Interval(interval float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("interval", strconv.FormatFloat(interval, 'g', -1, 64))
	return c
}

// Strike https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Provide a strike price to return options only at that strike price.
func (c *OptionChainsGetOptionChainCall) Strike(strike float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("strike", strconv.FormatFloat(strike, 'g', -1, 64))
	return c
}

// Range https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Returns options for the given range. Possible values are:
//
// ITM: In-the-money
// NTM: Near-the-money
// OTM: Out-of-the-money
// SAK: Strikes Above Market
// SBK: Strikes Below Market
// SNK: Strikes Near Market
// ALL: All Strikes
//
// Default is ALL.
func (c *OptionChainsGetOptionChainCall) Range(r string) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("range", r)
	return c
}

// FromDate https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Only return expirations after this date. For strategies, expiration refers to the nearest term expiration in the strategy. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
func (c *OptionChainsGetOptionChainCall) FromDate(date time.Time) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("fromDate", date.Format("2006-01-02"))
	return c
}

// ToDate https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Only return expirations before this date. For strategies, expiration refers to the nearest term expiration in the strategy. Valid ISO-8601 formats are: yyyy-MM-dd and yyyy-MM-dd'T'HH:mm:ssz.
func (c *OptionChainsGetOptionChainCall) ToDate(date time.Time) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("toDate", date.Format("2006-01-02"))
	return c
}

// Volatility https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Volatility to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param).
func (c *OptionChainsGetOptionChainCall) Volatility(volatility float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("volatility", strconv.FormatFloat(volatility, 'g', -1, 64))
	return c
}

// UnderlyingPrice https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Underlying price to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param).
func (c *OptionChainsGetOptionChainCall) UnderlyingPrice(underlyingPrice float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("underlyingPrice", strconv.FormatFloat(underlyingPrice, 'g', -1, 64))
	return c
}

// InterestRate https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Interest rate to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param).
func (c *OptionChainsGetOptionChainCall) InterestRate(interestRate float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("interestRate", strconv.FormatFloat(interestRate, 'g', -1, 64))
	return c
}

// DaysToExpiration https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Days to expiration to use in calculations. Applies only to ANALYTICAL strategy chains (see strategy param).
func (c *OptionChainsGetOptionChainCall) DaysToExpiration(daysToExpiration float64) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("daysToExpiration", strconv.FormatFloat(daysToExpiration, 'g', -1, 64))
	return c
}

// ExpMonth https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Return only options expiring in the specified month. Month is given in the three character format.
// Example: JAN
// Default is ALL.
func (c *OptionChainsGetOptionChainCall) ExpMonth(expMonth string) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("expMonth", expMonth)
	return c
}

// OptionType https://developer.tdameritrade.com/option-chains/apis/get/marketdata/chains
// Type of contracts to return. Possible values are:
//
// S: Standard contracts
// NS: Non-standard contracts
// ALL: All contracts
//
// Default is ALL.
func (c *OptionChainsGetOptionChainCall) OptionType(optionType string) *OptionChainsGetOptionChainCall {
	c.urlParams.Set("optionType", optionType)
	return c
}

func (c *OptionChainsGetOptionChainCall) doRequest() (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())

	var body io.Reader = nil
	urls := ResolveRelative(c.s.BasePath, "marketdata/chains")
	urls += "?" + c.urlParams.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, errors.Wrapf(err, "http.NewRequest")
	}
	req.Header = reqHeaders

	return SendRequest(c.ctx, c.s.client, req)
}

// Do send request
func (c *OptionChainsGetOptionChainCall) Do() (*OptionChain, error) {
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

	ret := &OptionChain{
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
