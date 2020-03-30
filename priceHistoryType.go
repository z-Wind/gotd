package gotd

// Candle https://developer.tdameritrade.com/price-history/apis
type Candle struct {
	Close    float64 `json:"close,omitempty"`
	Datetime Time    `json:"datetime,omitempty"`
	High     float64 `json:"high,omitempty"`
	Low      float64 `json:"low,omitempty"`
	Open     float64 `json:"open,omitempty"`
	Volume   float64 `json:"volume,omitempty"`
}

// PriceHistory https://developer.tdameritrade.com/price-history/apis
type PriceHistory struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Candles []Candle `json:"candles,omitempty"`
	Empty   bool     `json:"empty,omitempty"`
	Symbol  string   `json:"symbol,omitempty"` //"string"
}
