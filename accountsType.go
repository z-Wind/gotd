package gotd

import (
	"encoding/json"
)

// AccountList https://developer.tdameritrade.com/account-access/apis
type AccountList struct {
	Accounts []*Account

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Account https://developer.tdameritrade.com/account-access/apis
type Account struct {
	SecuritiesAccount `json:"securitiesAccount,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// SecuritiesAccount https://developer.tdameritrade.com/account-access/apis
type SecuritiesAccount struct {
	Type                    string             `json:"type,omitempty"`
	AccountID               string             `json:"accountId,omitempty"`
	RoundTrips              float64            `json:"roundTrips,omitempty"`
	IsDayTrader             bool               `json:"isDayTrader,omitempty"`
	IsClosingOnlyRestricted bool               `json:"isClosingOnlyRestricted,omitempty"`
	Positions               []*Position        `json:"positions,omitempty"`
	OrderStrategies         []*OrderStrategie  `json:"orderStrategies,omitempty"`
	InitialBalances         *InitialBalances   `json:"initialBalances,omitempty"`
	CurrentBalances         *CurrentBalances   `json:"currentBalances,omitempty"`
	ProjectedBalances       *ProjectedBalances `json:"projectedBalances,omitempty"`
}

// Position https://developer.tdameritrade.com/account-access/apis
type Position struct {
	ShortQuantity                  float64     `json:"shortQuantity,omitempty"`
	AveragePrice                   float64     `json:"averagePrice,omitempty"`
	CurrentDayProfitLoss           float64     `json:"currentDayProfitLoss,omitempty"`
	CurrentDayProfitLossPercentage float64     `json:"currentDayProfitLossPercentage,omitempty"`
	LongQuantity                   float64     `json:"longQuantity,omitempty"`
	SettledLongQuantity            float64     `json:"settledLongQuantity,omitempty"`
	SettledShortQuantity           float64     `json:"settledShortQuantity,omitempty"`
	AgedQuantity                   float64     `json:"agedQuantity,omitempty"`
	InstrumenA                     *Instrument `json:"instrument,omitempty"`
	MarketValue                    float64     `json:"marketValue,omitempty"`
}

// OrderLegCollection https://developer.tdameritrade.com/account-access/apis
type OrderLegCollection struct {
	OrderLegType   string      `json:"orderLegType,omitempty"` //"'EQUITY' or 'OPTION' or 'INDEX' or 'MUTUAL_FUND' or 'CASH_EQUIVALENT' or 'FIXED_INCOME' or 'CURRENCY'",
	LegID          int64       `json:"legId,omitempty"`
	Instrument     *Instrument `json:"instrument,omitempty"`     //"\"The type <Instrument> has the following subclasses [Option, MutualFund, CashEquivalent, Equity, FixedIncome] descriptions are listed below\"",
	Instruction    string      `json:"instruction,omitempty"`    //"'BUY' or 'SELL' or 'BUY_TO_COVER' or 'SELL_SHORT' or 'BUY_TO_OPEN' or 'BUY_TO_CLOSE' or 'SELL_TO_OPEN' or 'SELL_TO_CLOSE' or 'EXCHANGE'",
	PositionEffect string      `json:"positionEffect,omitempty"` //"'OPENING' or 'CLOSING' or 'AUTOMATIC'",
	Quantity       float64     `json:"quantity,omitempty"`
	QuantityType   string      `json:"quantityType,omitempty"` //"'ALL_SHARES' or 'DOLLARS' or 'SHARES'"
}

// OrderStrategie https://developer.tdameritrade.com/account-access/apis
type OrderStrategie struct {
	Session                  string                `json:"session,omitempty"`   //"'NORMAL' or 'AM' or 'PM' or 'SEAMLESS'",
	Duration                 string                `json:"duration,omitempty"`  //"'DAY' or 'GOOD_TILL_CANCEL' or 'FILL_OR_KILL'",
	OrderType                string                `json:"orderType,omitempty"` //"'MARKET' or 'LIMIT' or 'STOP' or 'STOP_LIMIT' or 'TRAILING_STOP' or 'MARKET_ON_CLOSE' or 'EXERCISE' or 'TRAILING_STOP_LIMIT' or 'NET_DEBIT' or 'NET_CREDIT' or 'NET_ZERO'",
	CancelTime               string                `json:"cancelTime,omitempty"`
	ComplexOrderStrategyType string                `json:"complexOrderStrategyType,omitempty"` //"'NONE' or 'COVERED' or 'VERTICAL' or 'BACK_RATIO' or 'CALENDAR' or 'DIAGONAL' or 'STRADDLE' or 'STRANGLE' or 'COLLAR_SYNTHETIC' or 'BUTTERFLY' or 'CONDOR' or 'IRON_CONDOR' or 'VERTICAL_ROLL' or 'COLLAR_WITH_STOCK' or 'DOUBLE_DIAGONAL' or 'UNBALANCED_BUTTERFLY' or 'UNBALANCED_CONDOR' or 'UNBALANCED_IRON_CONDOR' or 'UNBALANCED_VERTICAL_ROLL' or 'CUSTOM'",
	Quantity                 float64               `json:"quantity,omitempty"`
	FilledQuantity           float64               `json:"filledQuantity,omitempty"`
	RemainingQuantity        float64               `json:"remainingQuantity,omitempty"`
	RequestedDestination     string                `json:"requestedDestination,omitempty"` //"'INET' or 'ECN_ARCA' or 'CBOE' or 'AMEX' or 'PHLX' or 'ISE' or 'BOX' or 'NYSE' or 'NASDAQ' or 'BATS' or 'C2' or 'AUTO'",
	DestinationLinkName      string                `json:"destinationLinkName,omitempty"`
	ReleaseTime              string                `json:"releaseTime,omitempty"`
	StopPrice                float64               `json:"stopPrice,omitempty"`
	StopPriceLinkBasis       string                `json:"stopPriceLinkBasis,omitempty"` //"'MANUAL' or 'BASE' or 'TRIGGER' or 'LAST' or 'BID' or 'ASK' or 'ASK_BID' or 'MARK' or 'AVERAGE'",
	StopPriceLinkType        string                `json:"stopPriceLinkType,omitempty"`  //"'VALUE' or 'PERCENT' or 'TICK'",
	StopPriceOffset          float64               `json:"stopPriceOffset,omitempty"`
	StopType                 string                `json:"stopType,omitempty"`       //"'STANDARD' or 'BID' or 'ASK' or 'LAST' or 'MARK'",
	PriceLinkBasis           string                `json:"priceLinkBasis,omitempty"` //"'MANUAL' or 'BASE' or 'TRIGGER' or 'LAST' or 'BID' or 'ASK' or 'ASK_BID' or 'MARK' or 'AVERAGE'",
	PriceLinkType            string                `json:"priceLinkType,omitempty"`  //"'VALUE' or 'PERCENT' or 'TICK'",
	Price                    float64               `json:"price,omitempty"`
	TaxLotMethod             string                `json:"taxLotMethod,omitempty"` //"'FIFO' or 'LIFO' or 'HIGH_COST' or 'LOW_COST' or 'AVERAGE_COST' or 'SPECIFIC_LOT'",
	OrderLegCollections      []*OrderLegCollection `json:"orderLegCollection,omitempty"`
	ActivationPrice          float64               `json:"activationPrice,omitempty"`
	SpecialInstruction       string                `json:"specialInstruction,omitempty"` //"'ALL_OR_NONE' or 'DO_NOT_REDUCE' or 'ALL_OR_NONE_DO_NOT_REDUCE'",
	OrderStrategyType        string                `json:"orderStrategyType,omitempty"`  //"'SINGLE' or 'OCO' or 'TRIGGER'",
	OrderID                  int64                 `json:"orderId,omitempty"`
	Cancelable               bool                  `json:"cancelable,omitempty"`
	Editable                 bool                  `json:"editable,omitempty"`
	Status                   string                `json:"status,omitempty"` //"'AWAITING_PARENT_ORDER' or 'AWAITING_CONDITION' or 'AWAITING_MANUAL_REVIEW' or 'ACCEPTED' or 'AWAITING_UR_OUT' or 'PENDING_ACTIVATION' or 'QUEUED' or 'WORKING' or 'REJECTED' or 'PENDING_CANCEL' or 'CANCELED' or 'PENDING_REPLACE' or 'REPLACED' or 'FILLED' or 'EXPIRED'",
	EnteredTime              string                `json:"enteredTime,omitempty"`
	CloseTime                string                `json:"closeTime,omitempty"`
	Tag                      string                `json:"tag,omitempty"`
	AccountID                int64                 `json:"accountId,omitempty"`
	OrderActivityCollection  []*Execution          `json:"orderActivityCollection,omitempty"`  //: ["\"The type <OrderActivity> has the following subclasses [Execution] descriptions are listed below\""],
	ReplacingOrderCollection []json.RawMessage     `json:"replacingOrderCollection,omitempty"` //: [ {} ],
	ChildOrderStrategies     []json.RawMessage     `json:"childOrderStrategies,omitempty"`     //: [ {}  ],
	StatusDescription        string                `json:"statusDescription,omitempty"`
}

// InitialBalances https://developer.tdameritrade.com/account-access/apis
type InitialBalances struct {
	AccruedInterest            float64 `json:"accruedInterest,omitempty"`
	CashAvailableForTrading    float64 `json:"cashAvailableForTrading,omitempty"`
	CashAvailableForWithdrawal float64 `json:"cashAvailableForWithdrawal,omitempty"`
	CashBalance                float64 `json:"cashBalance,omitempty"`
	BondValue                  float64 `json:"bondValue,omitempty"`
	CashReceipts               float64 `json:"cashReceipts,omitempty"`
	LiquidationValue           float64 `json:"liquidationValue,omitempty"`
	LongOptionMarketValue      float64 `json:"longOptionMarketValue,omitempty"`
	LongStockValue             float64 `json:"longStockValue,omitempty"`
	MoneyMarketFund            float64 `json:"moneyMarketFund,omitempty"`
	MutualFundValue            float64 `json:"mutualFundValue,omitempty"`
	ShortOptionMarketValue     float64 `json:"shortOptionMarketValue,omitempty"`
	ShortStockValue            float64 `json:"shortStockValue,omitempty"`
	IsInCall                   bool    `json:"isInCall,omitempty"`
	UnsettledCash              float64 `json:"unsettledCash,omitempty"`
	CashDebitCallValue         float64 `json:"cashDebitCallValue,omitempty"`
	PendingDeposits            float64 `json:"pendingDeposits,omitempty"`
	AccountValue               float64 `json:"accountValue,omitempty"`
}

// CurrentBalances https://developer.tdameritrade.com/account-access/apis
type CurrentBalances struct {
	AccruedInterest              float64 `json:"accruedInterest,omitempty"`
	CashBalance                  float64 `json:"cashBalance,omitempty"`
	CashReceipts                 float64 `json:"cashReceipts,omitempty"`
	LongOptionMarketValue        float64 `json:"longOptionMarketValue,omitempty"`
	LiquidationValue             float64 `json:"liquidationValue,omitempty"`
	LongMarketValue              float64 `json:"longMarketValue,omitempty"`
	MoneyMarketFund              float64 `json:"moneyMarketFund,omitempty"`
	Savings                      float64 `json:"savings,omitempty"`
	ShortMarketValue             float64 `json:"shortMarketValue,omitempty"`
	PendingDeposits              float64 `json:"pendingDeposits,omitempty"`
	CashAvailableForTrading      float64 `json:"cashAvailableForTrading,omitempty"`
	CashAvailableForWithdrawal   float64 `json:"cashAvailableForWithdrawal,omitempty"`
	CashCall                     float64 `json:"cashCall,omitempty"`
	LongNonMarginableMarketValue float64 `json:"longNonMarginableMarketValue,omitempty"`
	TotalCash                    float64 `json:"totalCash,omitempty"`
	ShortOptionMarketValue       float64 `json:"shortOptionMarketValue,omitempty"`
	MutualFundValue              float64 `json:"mutualFundValue,omitempty"`
	BondValue                    float64 `json:"bondValue,omitempty"`
	CashDebitCallValue           float64 `json:"cashDebitCallValue,omitempty"`
	UnsettledCash                float64 `json:"unsettledCash,omitempty"`
}

// ProjectedBalances https://developer.tdameritrade.com/account-access/apis
type ProjectedBalances struct {
	AccruedInterest              float64 `json:"accruedInterest,omitempty"`
	CashBalance                  float64 `json:"cashBalance,omitempty"`
	CashReceipts                 float64 `json:"cashReceipts,omitempty"`
	LongOptionMarketValue        float64 `json:"longOptionMarketValue,omitempty"`
	LiquidationValue             float64 `json:"liquidationValue,omitempty"`
	LongMarketValue              float64 `json:"longMarketValue,omitempty"`
	MoneyMarketFund              float64 `json:"moneyMarketFund,omitempty"`
	Savings                      float64 `json:"savings,omitempty"`
	ShortMarketValue             float64 `json:"shortMarketValue,omitempty"`
	PendingDeposits              float64 `json:"pendingDeposits,omitempty"`
	CashAvailableForTrading      float64 `json:"cashAvailableForTrading,omitempty"`
	CashAvailableForWithdrawal   float64 `json:"cashAvailableForWithdrawal,omitempty"`
	CashCall                     float64 `json:"cashCall,omitempty"`
	LongNonMarginableMarketValue float64 `json:"longNonMarginableMarketValue,omitempty"`
	TotalCash                    float64 `json:"totalCash,omitempty"`
	ShortOptionMarketValue       float64 `json:"shortOptionMarketValue,omitempty"`
	MutualFundValue              float64 `json:"mutualFundValue,omitempty"`
	BondValue                    float64 `json:"bondValue,omitempty"`
	CashDebitCallValue           float64 `json:"cashDebitCallValue,omitempty"`
	UnsettledCash                float64 `json:"unsettledCash,omitempty"`
}

// OptionDeliverable https://developer.tdameritrade.com/account-access/apis
type OptionDeliverable struct {
	Symbol           string  `json:"symbol,omitempty"`
	DeliverableUnits float64 `json:"deliverableUnits,omitempty"`
	CurrencyType     string  `json:"currencyType,omitempty"` //"'USD' or 'CAD' or 'EUR' or 'JPY'",
	AssetType        string  `json:"assetType,omitempty"`    //"'EQUITY' or 'OPTION' or 'INDEX' or 'MUTUAL_FUND' or 'CASH_EQUIVALENT' or 'FIXED_INCOME' or 'CURRENCY'"
}

// ExecutionLeg https://developer.tdameritrade.com/account-access/apis
type ExecutionLeg struct {
	LegID             int64   `json:"legId,omitempty"`
	Quantity          float64 `json:"quantity,omitempty"`
	MismarkedQuantity float64 `json:"mismarkedQuantity,omitempty"`
	Price             float64 `json:"price,omitempty"`
	Time              string  `json:"time,omitempty"`
}

// Execution https://developer.tdameritrade.com/account-access/apis
type Execution struct {
	ActivityType           string          `json:"activityType,omitempty"`  //"'EXECUTION' or 'ORDER_ACTION'",
	ExecutionType          string          `json:"executionType,omitempty"` //"'FILL'",
	Quantity               float64         `json:"quantity,omitempty"`
	OrderRemainingQuantity float64         `json:"orderRemainingQuantity,omitempty"`
	ExecutionLegs          []*ExecutionLeg `json:"executionLegs,omitempty"`
}
