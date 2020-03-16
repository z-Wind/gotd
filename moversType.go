package gotd

// Mover https://developer.tdameritrade.com/movers/apis
type Mover struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Change      float64 `json:"change,omitempty"`
	Description string  `json:"description,omitempty"` //"string"
	Direction   string  `json:"direction,omitempty"`   //"'up' or 'down'"
	Last        float64 `json:"last,omitempty"`
	Symbol      string  `json:"symbol,omitempty"` //"string"
	TotalVolume float64 `json:"totalVolume,omitempty"`
}
