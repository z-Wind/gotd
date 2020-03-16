package gotd

// SavedOrderList https://developer.tdameritrade.com/account-access/apis
type SavedOrderList struct {
	SavedOrders []*SavedOrder

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// SavedOrder https://developer.tdameritrade.com/account-access/apis
type SavedOrder struct {
	*Order
	SavedOrderID int64  `json:"savedOrderId,omitempty"` // for saved order
	SavedTime    string `json:"savedTime,omitempty"`    // for saved order

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}
