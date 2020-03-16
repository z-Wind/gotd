package gotd

// Period https://developer.tdameritrade.com/market-hours/apis
type Period struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

// SessionHours https://developer.tdameritrade.com/market-hours/apis
type SessionHours struct {
	PreMarket     []Period `json:"preMarket,omitempty"`
	RegularMarket []Period `json:"regularMarket,omitempty"`
	PostMarket    []Period `json:"postMarket,omitempty"`
}

// MarketHourMap https://developer.tdameritrade.com/market-hours/apis
type MarketHourMap struct {
	MarketHourProductMaps map[string]map[string]*MarketHour

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// MarketHourProductMap https://developer.tdameritrade.com/market-hours/apis
type MarketHourProductMap struct {
	MarketHours map[string]*MarketHour

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// MarketHour https://developer.tdameritrade.com/market-hours/apis
type MarketHour struct {
	Category     string       `json:"category,omitempty"`
	Date         string       `json:"date,omitempty"`
	Exchange     string       `json:"exchange,omitempty"`
	IsOpen       bool         `json:"isOpen,omitempty"`
	MarketType   string       `json:"marketType,omitempty"` //"'BOND' or 'EQUITY' or 'ETF' or 'FOREX' or 'FUTURE' or 'FUTURE_OPTION' or 'INDEX' or 'INDICATOR' or 'MUTUAL_FUND' or 'OPTION' or 'UNKNOWN'"
	Product      string       `json:"product,omitempty"`
	ProductName  string       `json:"productName,omitempty"`
	SessionHours SessionHours `json:"sessionHours,omitempty"` //"object"
}
