package gotd

// WatchlistItem https://developer.tdameritrade.com/watchlist/apis
type WatchlistItem struct {
	SequenceID    int64       `json:"sequenceId,omitempty"`
	Quantity      float64     `json:"quantity,omitempty"`
	AveragePrice  float64     `json:"averagePrice,omitempty"`
	Commission    float64     `json:"commission,omitempty"`
	PurchasedDate string      `json:"purchasedDate,omitempty"` //"DateParam\"",
	Instrument    *Instrument `json:"instrument,omitempty"`
	Status        string      `json:"status,omitempty"` //"'UNCHANGED' or 'CREATED' or 'UPDATED' or 'DELETED'"
}

// WatchlistList https://developer.tdameritrade.com/watchlist/apis
type WatchlistList struct {
	Watchlists []*Watchlist

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Watchlist https://developer.tdameritrade.com/watchlist/apis
type Watchlist struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Name           string           `json:"name,omitempty"`        //"string"
	WatchlistID    string           `json:"watchlistId,omitempty"` //"string"
	AccountID      string           `json:"accountId,omitempty"`   //"string"
	Status         string           `json:"status,omitempty"`      //"'UNCHANGED' or 'CREATED' or 'UPDATED' or 'DELETED'"
	WatchlistItems []*WatchlistItem `json:"watchlistItems,omitempty"`
}
