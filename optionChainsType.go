package gotd

// Underlying https://developer.tdameritrade.com/option-chains/apis
type Underlying struct {
	Ask               float64 `json:"ask,omitempty"`
	AskSize           string  `json:"askSize,omitempty"`
	Bid               float64 `json:"bid,omitempty"`
	BidSize           string  `json:"bidSize,omitempty"`
	Change            float64 `json:"change,omitempty"`
	Close             float64 `json:"close,omitempty"`
	Delayed           bool    `json:"delayed,omitempty"`
	Description       string  `json:"description,omitempty"`  //"string"
	ExchangeName      string  `json:"exchangeName,omitempty"` //"string"
	FiftyTwoWeekHigh  float64 `json:"fiftyTwoWeekHigh,omitempty"`
	FiftyTwoWeekLow   float64 `json:"fiftyTwoWeekLow,omitempty"`
	HighPrice         float64 `json:"highPrice,omitempty"`
	Last              float64 `json:"last,omitempty"`
	LowPrice          float64 `json:"lowPrice,omitempty"`
	Mark              float64 `json:"mark,omitempty"`
	MarkChange        float64 `json:"markChange,omitempty"`
	MarkPercentChange float64 `json:"markPercentChange,omitempty"`
	OpenPrice         float64 `json:"openPrice,omitempty"`
	PercentChange     float64 `json:"percentChange,omitempty"`
	QuoteTime         string  `json:"quoteTime,omitempty"`
	Symbol            string  `json:"symbol,omitempty"` //"string"
	TotalVolume       string  `json:"totalVolume,omitempty"`
	TradeTime         string  `json:"tradeTime,omitempty"`
}

// Leg https://developer.tdameritrade.com/option-chains/apis
type Leg struct {
	Symbol      string  `json:"symbol,omitempty"`
	PutCallInd  string  `json:"putCallInd,omitempty"`
	Description string  `json:"description,omitempty"`
	Bid         float64 `json:"bid,omitempty"`
	Ask         float64 `json:"ask,omitempty"`
	Range       string  `json:"range,omitempty"`
	StrikePrice float64 `json:"strikePrice,omitempty"`
	TotalVolume float64 `json:"totalVolume,omitempty"`
}

// OptionStrategy https://developer.tdameritrade.com/option-chains/apis
type OptionStrategy struct {
	PrimaryLeg     *Leg    `json:"primaryLeg,omitempty"`
	SecondaryLeg   *Leg    `json:"secondaryLeg,omitempty"`
	StrategyStrike string  `json:"strategyStrike,omitempty"`
	StrategyBid    float64 `json:"strategyBid,omitempty"`
	StrategyAsk    float64 `json:"strategyAsk,omitempty"`
}

// MonthlyStrategy https://developer.tdameritrade.com/option-chains/apis
type MonthlyStrategy struct {
	Month              string            `json:"month,omitempty"`
	Year               int64             `json:"year,omitempty"`
	Day                int64             `json:"day,omitempty"`
	DaysToExp          int64             `json:"daysToExp,omitempty"`
	SecondaryMonth     string            `json:"secondaryMonth,omitempty"`
	SecondaryYear      int64             `json:"secondaryYear,omitempty"`
	SecondaryDay       int64             `json:"secondaryDay,omitempty"`
	SecondaryDaysToExp int64             `json:"secondaryDaysToExp,omitempty"`
	Type               string            `json:"type,omitempty"`
	SecondaryType      string            `json:"secondaryType,omitempty"`
	SecondaryLeap      bool              `json:"secondaryLeap,omitempty"`
	OptionStrategyList []*OptionStrategy `json:"optionStrategyList,omitempty"`
}

// OptionChain https://developer.tdameritrade.com/option-chains/apis
type OptionChain struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Symbol              string                          `json:"symbol,omitempty"` //"string"
	Status              string                          `json:"status,omitempty"` //"string"
	Underlying          *Underlying                     `json:"underlying,omitempty"`
	Strategy            string                          `json:"strategy,omitempty"` //"'SINGLE' or 'ANALYTICAL' or 'COVERED' or 'VERTICAL' or 'CALENDAR' or 'STRANGLE' or 'STRADDLE' or 'BUTTERFLY' or 'CONDOR' or 'DIAGONAL' or 'COLLAR' or 'ROLL'"
	Interval            float64                         `json:"interval,omitempty"`
	Intervals           []float64                       `json:"intervals,omitempty"`
	IsDelayed           bool                            `json:"isDelayed,omitempty"`
	IsIndex             bool                            `json:"isIndex,omitempty"`
	DaysToExpiration    float64                         `json:"daysToExpiration,omitempty"`
	InterestRate        float64                         `json:"interestRate,omitempty"`
	UnderlyingPrice     float64                         `json:"underlyingPrice,omitempty"`
	Volatility          float64                         `json:"volatility,omitempty"`
	MonthlyStrategyList []*MonthlyStrategy              `json:"monthlyStrategyList,omitempty"`
	CallExpDateMap      map[string]map[string][]*Option `json:"callExpDateMap,omitempty"` //"object"
	PutExpDateMap       map[string]map[string][]*Option `json:"putExpDateMap,omitempty"`  //"object"
}

// OptionDeliverables https://developer.tdameritrade.com/option-chains/apis
type OptionDeliverables struct {
	Symbol           string `json:"symbol,omitempty"`           //"string"
	AssetType        string `json:"assetType,omitempty"`        //"string"
	DeliverableUnits string `json:"deliverableUnits,omitempty"` //"string"
	CurrencyType     string `json:"currencyType,omitempty"`     //"string"
}

// Option https://developer.tdameritrade.com/option-chains/apis
type Option struct {
	PutCall                string                `json:"putCall,omitempty"`      //"'PUT' or 'CALL'"
	Symbol                 string                `json:"symbol,omitempty"`       //"string"
	Description            string                `json:"description,omitempty"`  //"string"
	ExchangeName           string                `json:"exchangeName,omitempty"` //"string"
	Bid                    float64               `json:"bid,omitempty"`
	Ask                    float64               `json:"ask,omitempty"`
	Last                   float64               `json:"last,omitempty"`
	Mark                   float64               `json:"mark,omitempty"`
	BidSize                float64               `json:"bidSize,omitempty"`
	AskSize                float64               `json:"askSize,omitempty"`
	LastSize               float64               `json:"lastSize,omitempty"`
	HighPrice              float64               `json:"highPrice,omitempty"`
	LowPrice               float64               `json:"lowPrice,omitempty"`
	OpenPrice              float64               `json:"openPrice,omitempty"`
	ClosePrice             float64               `json:"closePrice,omitempty"`
	TotalVolume            float64               `json:"totalVolume,omitempty"`
	TradeDate              string                `json:"tradeDate,omitempty"`
	QuoteTimeInLong        float64               `json:"quoteTimeInLong,omitempty"`
	TradeTimeInLong        float64               `json:"tradeTimeInLong,omitempty"`
	NetChange              float64               `json:"netChange,omitempty"`
	Volatility             float64               `json:"volatility,omitempty"`
	Delta                  float64               `json:"delta,omitempty"`
	Gamma                  float64               `json:"gamma,omitempty"`
	Theta                  float64               `json:"theta,omitempty"`
	Vega                   float64               `json:"vega,omitempty"`
	Rho                    float64               `json:"rho,omitempty"`
	TimeValue              float64               `json:"timeValue,omitempty"`
	OpenInterest           float64               `json:"openInterest,omitempty"`
	IsInTheMoney           bool                  `json:"isInTheMoney,omitempty"`
	TheoreticalOptionValue float64               `json:"theoreticalOptionValue,omitempty"`
	TheoreticalVolatility  float64               `json:"theoreticalVolatility,omitempty"`
	IsMini                 bool                  `json:"isMini,omitempty"`
	IsNonStandard          bool                  `json:"isNonStandard,omitempty"`
	OptionDeliverablesList []*OptionDeliverables `json:"optionDeliverablesList,omitempty"`
	StrikePrice            float64               `json:"strikePrice,omitempty"`
	ExpirationDate         float64               `json:"expirationDate,omitempty"`
	DaysToExpiration       float64               `json:"daysToExpiration,omitempty"`
	ExpirationType         string                `json:"expirationType,omitempty"` //"string"
	LastTradingDay         float64               `json:"lastTradingDay,omitempty"`
	Multiplier             float64               `json:"multiplier,omitempty"`
	SettlementType         string                `json:"settlementType,omitempty"`  //"string"
	DeliverableNote        string                `json:"deliverableNote,omitempty"` //"string"
	IsIndexOption          bool                  `json:"isIndexOption,omitempty"`
	PercentChange          float64               `json:"percentChange,omitempty"`
	MarkChange             float64               `json:"markChange,omitempty"`
	MarkPercentChange      float64               `json:"markPercentChange,omitempty"`
	NonStandard            bool                  `json:"nonStandard,omitempty"`
	Mini                   bool                  `json:"mini,omitempty"`
	InTheMoney             bool                  `json:"inTheMoney,omitempty"`
}
