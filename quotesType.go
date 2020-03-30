package gotd

// QuoteMap https://developer.tdameritrade.com/quotes/apis
type QuoteMap struct {
	Quotes map[string]*Quote

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Quote https://developer.tdameritrade.com/quotes/apis
type Quote struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	// Quote
	// ===========================================================================
	AssetType string `json:"assetType,omitempty"`

	// MutualFund
	Symbol          string  `json:"symbol,omitempty"`
	Description     string  `json:"description,omitempty"`
	ClosePrice      float64 `json:"closePrice,omitempty"`
	NetChange       float64 `json:"netChange,omitempty"`
	TotalVolume     float64 `json:"totalVolume,omitempty"`
	TradeTimeInLong float64 `json:"tradeTimeInLong,omitempty"`
	Exchange        string  `json:"exchange,omitempty"`
	ExchangeName    string  `json:"exchangeName,omitempty"`
	Digits          float64 `json:"digits,omitempty"`
	Wk52High        float64 `json:"52WkHigh,omitempty"`
	Wk52Low         float64 `json:"52WkLow,omitempty"`
	NAV             float64 `json:"nAV,omitempty"`
	PeRatio         float64 `json:"peRatio,omitempty"`
	DivAmount       float64 `json:"divAmount,omitempty"`
	DivYield        float64 `json:"divYield,omitempty"`
	DivDate         string  `json:"divDate,omitempty"`
	SecurityStatus  string  `json:"securityStatus,omitempty"`
	Delayed         bool    `json:"Delayed,omitempty"`

	// Future
	//Symbol                string  `json:"symbol,omitempty"`
	BidPriceInDouble   float64 `json:"bidPriceInDouble,omitempty"`
	AskPriceInDouble   float64 `json:"askPriceInDouble,omitempty"`
	LastPriceInDouble  float64 `json:"lastPriceInDouble,omitempty"`
	BidID              string  `json:"bidId,omitempty"`
	BidTick            string  `json:"bidTick,omitempty"`
	AskID              string  `json:"askId,omitempty"`
	HighPriceInDouble  float64 `json:"highPriceInDouble,omitempty"`
	LowPriceInDouble   float64 `json:"lowPriceInDouble,omitempty"`
	ClosePriceInDouble float64 `json:"closePriceInDouble,omitempty"`
	//Exchange              string  `json:"exchange,omitempty"`
	//Description           string  `json:"description,omitempty"`
	LastID              string  `json:"lastId,omitempty"`
	OpenPriceInDouble   float64 `json:"openPriceInDouble,omitempty"`
	ChangeInDouble      float64 `json:"changeInDouble,omitempty"`
	FuturePercentChange float64 `json:"futurePercentChange,omitempty"`
	//ExchangeName          string  `json:"exchangeName,omitempty"`
	//SecurityStatus        string  `json:"securityStatus,omitempty"`
	OpenInterest          float64 `json:"openInterest,omitempty"`
	Mark                  float64 `json:"mark,omitempty"`
	Tick                  float64 `json:"tick,omitempty"`
	TickAmount            float64 `json:"tickAmount,omitempty"`
	Product               string  `json:"product,omitempty"`
	FuturePriceFormat     string  `json:"futurePriceFormat,omitempty"`
	FutureTradingHours    string  `json:"futureTradingHours,omitempty"`
	FutureIsTradable      bool    `json:"futureIsTradable,omitempty"`
	FutureMultiplier      float64 `json:"futureMultiplier,omitempty"`
	FutureIsActive        bool    `json:"futureIsActive,omitempty"`
	FutureSettlementPrice float64 `json:"futureSettlementPrice,omitempty"`
	FutureActiveSymbol    string  `json:"futureActiveSymbol,omitempty"`
	FutureExpirationDate  string  `json:"futureExpirationDate,omitempty"`

	// FutureOptions
	//Symbol                      string  `json:"symbol,omitempty"`
	//BidPriceInDouble            float64 `json:"bidPriceInDouble,omitempty"`
	//AskPriceInDouble            float64 `json:"askPriceInDouble,omitempty"`
	//LastPriceInDouble           float64 `json:"lastPriceInDouble,omitempty"`
	//HighPriceInDouble           float64 `json:"highPriceInDouble,omitempty"`
	//LowPriceInDouble            float64 `json:"lowPriceInDouble,omitempty"`
	//ClosePriceInDouble          float64 `json:"closePriceInDouble,omitempty"`
	//Description                 string  `json:"description,omitempty"`
	//OpenPriceInDouble           float64 `json:"openPriceInDouble,omitempty"`
	NetChangeInDouble float64 `json:"netChangeInDouble,omitempty"`
	//OpenInterest                float64 `json:"openInterest,omitempty"`
	//ExchangeName                string  `json:"exchangeName,omitempty"`
	//SecurityStatus              string  `json:"securityStatus,omitempty"`
	Volatility                  float64 `json:"volatility,omitempty"`
	MoneyIntrinsicValueInDouble float64 `json:"moneyIntrinsicValueInDouble,omitempty"`
	MultiplierInDouble          float64 `json:"multiplierInDouble,omitempty"`
	//Digits                      float64 `json:"digits,omitempty"`
	StrikePriceInDouble float64 `json:"strikePriceInDouble,omitempty"`
	ContractType        string  `json:"contractType,omitempty"`
	Underlying          string  `json:"underlying,omitempty"`
	TimeValueInDouble   float64 `json:"timeValueInDouble,omitempty"`
	DeltaInDouble       float64 `json:"deltaInDouble,omitempty"`
	GammaInDouble       float64 `json:"gammaInDouble,omitempty"`
	ThetaInDouble       float64 `json:"thetaInDouble,omitempty"`
	VegaInDouble        float64 `json:"vegaInDouble,omitempty"`
	RhoInDouble         float64 `json:"rhoInDouble,omitempty"`
	//Mark                        float64 `json:"mark,omitempty"`
	//Tick                        float64 `json:"tick,omitempty"`
	//TickAmount                  float64 `json:"tickAmount,omitempty"`
	//FutureIsTradable            bool    `json:"futureIsTradable,omitempty"`
	//FutureTradingHours          string  `json:"futureTradingHours,omitempty"`
	//FuturePercentChange         float64 `json:"futurePercentChange,omitempty"`
	//FutureIsActive              bool    `json:"futureIsActive,omitempty"`
	//FutureExpirationDate        float64 `json:"futureExpirationDate,omitempty"`
	ExpirationType string `json:"expirationType,omitempty"`
	ExerciseType   string `json:"exerciseType,omitempty"`
	InTheMoney     bool   `json:"inTheMoney,omitempty"`

	// Index
	//Symbol          string  `json:"symbol,omitempty"`
	//Description     string  `json:"description,omitempty"`
	LastPrice float64 `json:"lastPrice,omitempty"`
	OpenPrice float64 `json:"openPrice,omitempty"`
	HighPrice float64 `json:"highPrice,omitempty"`
	LowPrice  float64 `json:"lowPrice,omitempty"`
	//ClosePrice      float64 `json:"closePrice,omitempty"`
	//NetChange       float64 `json:"netChange,omitempty"`
	//TotalVolume     float64 `json:"totalVolume,omitempty"`
	//TradeTimeInLong float64 `json:"tradeTimeInLong,omitempty"`
	//Exchange        string  `json:"exchange,omitempty"`
	//ExchangeName    string  `json:"exchangeName,omitempty"`
	//Digits          float64 `json:"digits,omitempty"`
	//Wk52High        float64 `json:"52WkHigh,omitempty"`
	//Wk52Low         float64 `json:"52WkLow,omitempty"`
	//SecurityStatus  string  `json:"securityStatus,omitempty"`

	// Option
	//Symbol                 string  `json:"symbol,omitempty"`
	//Description            string  `json:"description,omitempty"`
	BidPrice float64 `json:"bidPrice,omitempty"`
	BidSize  float64 `json:"bidSize,omitempty"`
	AskPrice float64 `json:"askPrice,omitempty"`
	AskSize  float64 `json:"askSize,omitempty"`
	//LastPrice              float64 `json:"lastPrice,omitempty"`
	LastSize float64 `json:"lastSize,omitempty"`
	//OpenPrice              float64 `json:"openPrice,omitempty"`
	//HighPrice              float64 `json:"highPrice,omitempty"`
	//LowPrice               float64 `json:"lowPrice,omitempty"`
	//ClosePrice             float64 `json:"closePrice,omitempty"`
	//NetChange              float64 `json:"netChange,omitempty"`
	//TotalVolume            float64 `json:"totalVolume,omitempty"`
	QuoteTimeInLong float64 `json:"quoteTimeInLong,omitempty"`
	//TradeTimeInLong        float64 `json:"tradeTimeInLong,omitempty"`
	//Mark                   float64 `json:"mark,omitempty"`
	//OpenInterest           float64 `json:"openInterest,omitempty"`
	//Volatility             float64 `json:"volatility,omitempty"`
	MoneyIntrinsicValue float64 `json:"moneyIntrinsicValue,omitempty"`
	Multiplier          float64 `json:"multiplier,omitempty"`
	StrikePrice         float64 `json:"strikePrice,omitempty"`
	//ContractType           string  `json:"contractType,omitempty"`
	//Underlying             string  `json:"underlying,omitempty"`
	TimeValue    float64 `json:"timeValue,omitempty"`
	Deliverables string  `json:"deliverables,omitempty"`
	Delta        float64 `json:"delta,omitempty"`
	Gamma        float64 `json:"gamma,omitempty"`
	Theta        float64 `json:"theta,omitempty"`
	Vega         float64 `json:"vega,omitempty"`
	Rho          float64 `json:"rho,omitempty"`
	//SecurityStatus         string  `json:"securityStatus,omitempty"`
	TheoreticalOptionValue float64 `json:"theoreticalOptionValue,omitempty"`
	UnderlyingPrice        float64 `json:"underlyingPrice,omitempty"`
	UvExpirationType       string  `json:"uvExpirationType,omitempty"`
	//Exchange               string  `json:"exchange,omitempty"`
	//ExchangeName           string  `json:"exchangeName,omitempty"`
	SettlementType string `json:"settlementType,omitempty"`

	// Forex
	//Symbol             string  `json:"symbol,omitempty"`
	//BidPriceInDouble   float64 `json:"bidPriceInDouble,omitempty"`
	//AskPriceInDouble   float64 `json:"askPriceInDouble,omitempty"`
	//LastPriceInDouble  float64 `json:"lastPriceInDouble,omitempty"`
	//HighPriceInDouble  float64 `json:"highPriceInDouble,omitempty"`
	//LowPriceInDouble   float64 `json:"lowPriceInDouble,omitempty"`
	//ClosePriceInDouble float64 `json:"closePriceInDouble,omitempty"`
	//Exchange           string  `json:"exchange,omitempty"`
	//Description        string  `json:"description,omitempty"`
	//OpenPriceInDouble  float64 `json:"openPriceInDouble,omitempty"`
	//ChangeInDouble     float64 `json:"changeInDouble,omitempty"`
	PercentChange float64 `json:"percentChange,omitempty"`
	//ExchangeName       string  `json:"exchangeName,omitempty"`
	//Digits             float64 `json:"digits,omitempty"`
	//SecurityStatus     string  `json:"securityStatus,omitempty"`
	//Tick               float64 `json:"tick,omitempty"`
	//TickAmount         float64 `json:"tickAmount,omitempty"`
	//Product            string  `json:"product,omitempty"`
	TradingHours string `json:"tradingHours,omitempty"`
	IsTradable   bool   `json:"isTradable,omitempty"`
	MarketMaker  string `json:"marketMaker,omitempty"`
	//Wk52High           float64 `json:"52WkHigh,omitempty"`
	//Wk52Low            float64 `json:"52WkLow,omitempty"`
	//Mark               float64 `json:"mark,omitempty"`

	// ETF
	//Symbol                       string  `json:"symbol,omitempty"`
	//Description                  string  `json:"description,omitempty"`
	//BidPrice                     float64 `json:"bidPrice,omitempty"`
	//BidSize                      float64 `json:"bidSize,omitempty"`
	//BidID                        string  `json:"bidId,omitempty"`
	//AskPrice                     float64 `json:"askPrice,omitempty"`
	//AskSize                      float64 `json:"askSize,omitempty"`
	//AskID                        string  `json:"askId,omitempty"`
	//LastPrice                    float64 `json:"lastPrice,omitempty"`
	//LastSize                     float64 `json:"lastSize,omitempty"`
	//LastID                       string  `json:"lastId,omitempty"`
	//OpenPrice                    float64 `json:"openPrice,omitempty"`
	//HighPrice                    float64 `json:"highPrice,omitempty"`
	//LowPrice                     float64 `json:"lowPrice,omitempty"`
	//ClosePrice                   float64 `json:"closePrice,omitempty"`
	//NetChange                    float64 `json:"netChange,omitempty"`
	//TotalVolume                  float64 `json:"totalVolume,omitempty"`
	//QuoteTimeInLong              float64 `json:"quoteTimeInLong,omitempty"`
	//TradeTimeInLong              float64 `json:"tradeTimeInLong,omitempty"`
	//Mark                         float64 `json:"mark,omitempty"`
	//Exchange                     string  `json:"exchange,omitempty"`
	//ExchangeName                 string  `json:"exchangeName,omitempty"`
	Marginable bool `json:"marginable,omitempty"`
	Shortable  bool `json:"shortable,omitempty"`
	//Volatility                   float64 `json:"volatility,omitempty"`
	//Digits                       float64 `json:"digits,omitempty"`
	//Wk52High                     float64 `json:"52WkHigh,omitempty"`
	//Wk52Low                      float64 `json:"52WkLow,omitempty"`
	//PeRatio                      float64 `json:"peRatio,omitempty"`
	//DivAmount                    float64 `json:"divAmount,omitempty"`
	//DivYield                     float64 `json:"divYield,omitempty"`
	//DivDate                      string  `json:"divDate,omitempty"`
	//SecurityStatus               string  `json:"securityStatus,omitempty"`
	RegularMarketLastPrice       float64 `json:"regularMarketLastPrice,omitempty"`
	RegularMarketLastSize        float64 `json:"regularMarketLastSize,omitempty"`
	RegularMarketNetChange       float64 `json:"regularMarketNetChange,omitempty"`
	RegularMarketTradeTimeInLong float64 `json:"regularMarketTradeTimeInLong,omitempty"`

	// Equity
	//Symbol                       string  `json:"symbol,omitempty"`
	//Description                  string  `json:"description,omitempty"`
	//BidPrice                     float64 `json:"bidPrice,omitempty"`
	//BidSize                      string   `json:"bidSize,omitempty"`
	//BidID                        string  `json:"bidId,omitempty"`
	//AskPrice                     float64 `json:"askPrice,omitempty"`
	//AskSize                      string   `json:"askSize,omitempty"`
	//AskID                        string   `json:"askId,omitempty"`
	//LastPrice                    float64 `json:"lastPrice,omitempty"`
	//LastSize                     string   `json:"lastSize,omitempty"`
	//LastID                       string   `json:"lastId,omitempty"`
	//OpenPrice                    float64 `json:"openPrice,omitempty"`
	//HighPrice                    float64 `json:"highPrice,omitempty"`
	//LowPrice                     float64 `json:"lowPrice,omitempty"`
	//ClosePrice                   float64 `json:"closePrice,omitempty"`
	//NetChange                    float64 `json:"netChange,omitempty"`
	//TotalVolume                  float64 `json:"totalVolume,omitempty"`
	//QuoteTimeInLong              string   `json:"quoteTimeInLong,omitempty"`
	//TradeTimeInLong              string   `json:"tradeTimeInLong,omitempty"`
	//Mark                         float64 `json:"mark,omitempty"`
	//Exchange                     string  `json:"exchange,omitempty"`
	//ExchangeName                 string  `json:"exchangeName,omitempty"`
	//Marginable                   bool    `json:"marginable,omitempty"`
	//Shortable                    bool    `json:"shortable,omitempty"`
	//Volatility                   float64 `json:"volatility,omitempty"`
	//Digits                       float64 `json:"digits,omitempty"`
	//Wk52High                     float64 `json:"52WkHigh,omitempty"`
	//Wk52Low                      float64 `json:"52WkLow,omitempty"`
	//PeRatio                      float64 `json:"peRatio,omitempty"`
	//DivAmount                    float64 `json:"divAmount,omitempty"`
	//DivYield                     float64 `json:"divYield,omitempty"`
	//DivDate                      string  `json:"divDate,omitempty"`
	//SecurityStatus               string  `json:"securityStatus,omitempty"`
	//RegularMarketLastPrice       float64 `json:"regularMarketLastPrice,omitempty"`
	//RegularMarketLastSize        float64 `json:"regularMarketLastSize,omitempty"`
	//RegularMarketNetChange       float64 `json:"regularMarketNetChange,omitempty"`
	//RegularMarketTradeTimeInLong float64 `json:"regularMarketTradeTimeInLong,omitempty"`
	// ===========================================================================
}
