package gotd

// TransactionItem https://developer.tdameritrade.com/transaction-history/apis
type TransactionItem struct {
	AccountID            int64       `json:"accountId,omitempty"`
	Amount               float64     `json:"amount,omitempty"`
	Price                float64     `json:"price,omitempty"`
	Cost                 float64     `json:"cost,omitempty"`
	ParentOrderKey       float64     `json:"parentOrderKey,omitempty"`
	ParentChildIndicator string      `json:"parentChildIndicator,omitempty"` //"string"
	Instruction          string      `json:"instruction,omitempty"`          //"string"
	PositionEffect       string      `json:"positionEffect,omitempty"`       //"string"
	Instrument           *Instrument `json:"instrument,omitempty"`
}

// Fees https://developer.tdameritrade.com/transaction-history/apis
type Fees struct {
	RFee          float64 `json:"rFee,omitempty"`
	AdditionalFee float64 `json:"additionalFee,omitempty"`
	CDSCFee       float64 `json:"cdscFee,omitempty"`
	RegFee        float64 `json:"regFee,omitempty"`
	OtherCharges  float64 `json:"otherCharges,omitempty"`
	Commission    float64 `json:"commission,omitempty"`
	OptRegFee     float64 `json:"optRegFee,omitempty"`
	SecFee        float64 `json:"secFee,omitempty"`
}

// TransactionList https://developer.tdameritrade.com/transaction-history/apis
type TransactionList struct {
	Transactions []*Transaction

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Transaction https://developer.tdameritrade.com/transaction-history/apis
type Transaction struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Type                          string           `json:"type,omitempty"`                    //"'TRADE' or 'RECEIVE_AND_DELIVER' or 'DIVIDEND_OR_INTEREST' or 'ACH_RECEIPT' or 'ACH_DISBURSEMENT' or 'CASH_RECEIPT' or 'CASH_DISBURSEMENT' or 'ELECTRONIC_FUND' or 'WIRE_OUT' or 'WIRE_IN' or 'JOURNAL' or 'MEMORANDUM' or 'MARGIN_CALL' or 'MONEY_MARKET' or 'SMA_ADJUSTMENT'"
	ClearingReferenceNumber       string           `json:"clearingReferenceNumber,omitempty"` //"string"
	SubAccount                    string           `json:"subAccount,omitempty"`              //"string"
	SettlementDate                string           `json:"settlementDate,omitempty"`          //"string"
	OrderID                       string           `json:"orderId,omitempty"`                 //"string"
	Sma                           float64          `json:"sma,omitempty"`
	RequirementReallocationAmount float64          `json:"requirementReallocationAmount,omitempty"`
	DayTradeBuyingPowerEffect     float64          `json:"dayTradeBuyingPowerEffect,omitempty"`
	NetAmount                     float64          `json:"netAmount,omitempty"`
	TransactionDate               string           `json:"transactionDate,omitempty"`    //"string"
	OrderDate                     string           `json:"orderDate,omitempty"`          //"string"
	TransactionSubType            string           `json:"transactionSubType,omitempty"` //"string"
	TransactionID                 int64            `json:"transactionId,omitempty"`
	CashBalanceEffectFlag         bool             `json:"cashBalanceEffectFlag,omitempty"`
	Description                   string           `json:"description,omitempty"` //"string"
	AchStatus                     string           `json:"achStatus,omitempty"`   //"'Approved' or 'Rejected' or 'Cancel' or 'Error'"
	AccruedInterest               float64          `json:"accruedInterest,omitempty"`
	Fees                          *Fees            `json:"fees,omitempty"` //"object"
	TransactionItem               *TransactionItem `json:"transactionItem,omitempty"`
}
