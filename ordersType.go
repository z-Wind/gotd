package gotd

// OrderList https://developer.tdameritrade.com/account-access/apis
type OrderList struct {
	Orders []*Order

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Order https://developer.tdameritrade.com/account-access/apis
type Order struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Session                   string                `json:"session,omitempty"`   //"'NORMAL' or 'AM' or 'PM' or 'SEAMLESS'"
	Duration                  string                `json:"duration,omitempty"`  //"'DAY' or 'GOOD_TILL_CANCEL' or 'FILL_OR_KILL'"
	OrderType                 string                `json:"orderType,omitempty"` //"'MARKET' or 'LIMIT' or 'STOP' or 'STOP_LIMIT' or 'TRAILING_STOP' or 'MARKET_ON_CLOSE' or 'EXERCISE' or 'TRAILING_STOP_LIMIT' or 'NET_DEBIT' or 'NET_CREDIT' or 'NET_ZERO'"
	CancelTime                string                `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType  string                `json:"complexOrderStrategyType,omitempty"` //"'NONE' or 'COVERED' or 'VERTICAL' or 'BACK_RATIO' or 'CALENDAR' or 'DIAGONAL' or 'STRADDLE' or 'STRANGLE' or 'COLLAR_SYNTHETIC' or 'BUTTERFLY' or 'CONDOR' or 'IRON_CONDOR' or 'VERTICAL_ROLL' or 'COLLAR_WITH_STOCK' or 'DOUBLE_DIAGONAL' or 'UNBALANCED_BUTTERFLY' or 'UNBALANCED_CONDOR' or 'UNBALANCED_IRON_CONDOR' or 'UNBALANCED_VERTICAL_ROLL' or 'CUSTOM'"
	Quantity                  float64               `json:"quantity,omitempty"`
	FilledQuantity            float64               `json:"filledQuantity,omitempty"`
	RemainingQuantity         float64               `json:"remainingQuantity,omitempty"`
	RequestedDestination      string                `json:"requestedDestination,omitempty"` //"'INET' or 'ECN_ARCA' or 'CBOE' or 'AMEX' or 'PHLX' or 'ISE' or 'BOX' or 'NYSE' or 'NASDAQ' or 'BATS' or 'C2' or 'AUTO'"
	DestinationLinkName       string                `json:"destinationLinkName,omitempty"`
	ReleaseTime               string                `json:"releaseTime,omitempty"`
	StopPrice                 float64               `json:"stopPrice,omitempty"`
	StopPriceLinkBasis        string                `json:"stopPriceLinkBasis,omitempty"` //"'MANUAL' or 'BASE' or 'TRIGGER' or 'LAST' or 'BID' or 'ASK' or 'ASK_BID' or 'MARK' or 'AVERAGE'"
	StopPriceLinkType         string                `json:"stopPriceLinkType,omitempty"`  //"'VALUE' or 'PERCENT' or 'TICK'"
	StopPriceOffset           float64               `json:"stopPriceOffset,omitempty"`
	StopType                  string                `json:"stopType,omitempty"`       //"'STANDARD' or 'BID' or 'ASK' or 'LAST' or 'MARK'"
	PriceLinkBasis            string                `json:"priceLinkBasis,omitempty"` //"'MANUAL' or 'BASE' or 'TRIGGER' or 'LAST' or 'BID' or 'ASK' or 'ASK_BID' or 'MARK' or 'AVERAGE'"
	PriceLinkType             string                `json:"priceLinkType,omitempty"`  //"'VALUE' or 'PERCENT' or 'TICK'"
	Price                     float64               `json:"price,omitempty"`
	TaxLotMethod              string                `json:"taxLotMethod,omitempty"` //"'FIFO' or 'LIFO' or 'HIGH_COST' or 'LOW_COST' or 'AVERAGE_COST' or 'SPECIFIC_LOT'"
	OrderLegCollections       []*OrderLegCollection `json:"orderLegCollection,omitempty"`
	ActivationPrice           float64               `json:"activationPrice,omitempty"`
	SpecialInstruction        string                `json:"specialInstruction,omitempty"` //"'ALL_OR_NONE' or 'DO_NOT_REDUCE' or 'ALL_OR_NONE_DO_NOT_REDUCE'"
	OrderStrategyType         string                `json:"orderStrategyType,omitempty"`  //"'SINGLE' or 'OCO' or 'TRIGGER'"
	OrderID                   int64                 `json:"orderId,omitempty"`
	Cancelable                bool                  `json:"cancelable,omitempty"`
	Editable                  bool                  `json:"editable,omitempty"`
	Status                    string                `json:"status,omitempty"` //"'AWAITING_PARENT_ORDER' or 'AWAITING_CONDITION' or 'AWAITING_MANUAL_REVIEW' or 'ACCEPTED' or 'AWAITING_UR_OUT' or 'PENDING_ACTIVATION' or 'QUEUED' or 'WORKING' or 'REJECTED' or 'PENDING_CANCEL' or 'CANCELED' or 'PENDING_REPLACE' or 'REPLACED' or 'FILLED' or 'EXPIRED'"
	EnteredTime               string                `json:"enteredTime,omitempty"`
	CloseTime                 string                `json:"closeTime,omitempty"`
	Tag                       string                `json:"tag,omitempty"`
	AccountID                 float64               `json:"accountId,omitempty"`
	OrderActivityCollections  []*Execution          `json:"orderActivityCollection,omitempty"`
	ReplacingOrderCollections []*Order              `json:"replacingOrderCollection,omitempty"`
	ChildOrderStrategies      []*Order              `json:"childOrderStrategies,omitempty"`
	StatusDescription         string                `json:"statusDescription,omitempty"`
}
