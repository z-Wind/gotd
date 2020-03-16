package gotd

// InstrumentMap https://developer.tdameritrade.com/instruments/apis
type InstrumentMap struct {
	Instruments map[string]*Instrument

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`
}

// Instrument https://developer.tdameritrade.com/instruments/apis
type Instrument struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Cusip       string `json:"cusip,omitempty"`
	Symbol      string `json:"symbol,omitempty"`
	Description string `json:"description,omitempty"`
	Exchange    string `json:"exchange,omitempty"`
	AssetType   string `json:"assetType,omitempty"` //"'EQUITY' or 'ETF' or 'FOREX' or 'FUTURE' or 'FUTURE_OPTION' or 'INDEX' or 'INDICATOR' or 'MUTUAL_FUND' or 'OPTION' or 'UNKNOWN' or 'BOND'"

	Fundamental *FundamentalData `json:"fundamental,omitempty"`
	BondPrice   float64          `json:"bondPrice,omitempty"` // for "'BOND'"

	// Accounts
	//================================================================
	//AssetType string `json:"assetType,omitempty"`

	// EQUITY
	//Cusip       string `json:"cusip,omitempty"`
	//Symbol      string `json:"symbol,omitempty"`
	//Description string `json:"description,omitempty"`

	// OPTION
	//Cusip              string               `json:"cusip,omitempty"`
	//Symbol             string               `json:"symbol,omitempty"`
	//Description        string               `json:"description,omitempty"`
	Type               string               `json:"type,omitempty"`    //"'VANILLA' or 'BINARY' or 'BARRIER'",
	PutCall            string               `json:"putCall,omitempty"` //"'PUT' or 'CALL'",
	UnderlyingSymbol   string               `json:"underlyingSymbol,omitempty"`
	OptionMultiplier   float64              `json:"optionMultiplier,omitempty"`
	OptionDeliverables []*OptionDeliverable `json:"optionDeliverables,omitempty"`

	// MUTUAL_FUND
	//Cusip       string `json:"cusip,omitempty"`
	//Symbol      string `json:"symbol,omitempty"`
	//Description string `json:"description,omitempty"`
	//Type        string `json:"type,omitempty"` //"'NOT_APPLICABLE' or 'OPEN_END_NON_TAXABLE' or 'OPEN_END_TAXABLE' or 'NO_LOAD_NON_TAXABLE' or 'NO_LOAD_TAXABLE'"

	// CASH_EQUIVALENT
	//Cusip       string `json:"cusip,omitempty"`
	//Symbol      string `json:"symbol,omitempty"`
	//Description string `json:"description,omitempty"`
	//Type        string `json:"type,omitempty"` //"'SAVINGS' or 'MONEY_MARKET_FUND'"

	// FIXED_INCOME
	//Cusip        string  `json:"cusip,omitempty"`
	//Symbol       string  `json:"symbol,omitempty"`
	//Description  string  `json:"description,omitempty"`
	MaturityDate string  `json:"maturityDate,omitempty"`
	VariableRate float64 `json:"variableRate,omitempty"`
	Factor       float64 `json:"factor,omitempty"`
	// CURRENCY 未定義
	// INDEX 未定義
	//================================================================

	// transactionHistory
	//================================================================
	//Symbol               string  `json:"symbol,omitempty"`               //"string"
	//UnderlyingSymbol     string  `json:"underlyingSymbol,omitempty"`     //"string"
	OptionExpirationDate string  `json:"optionExpirationDate,omitempty"` //"string"
	OptionStrikePrice    float64 `json:"optionStrikePrice,omitempty"`
	//PutCall              string  `json:"putCall,omitempty"`          //"string"
	//Cusip                string  `json:"cusip,omitempty"`            //"string"
	//Description          string  `json:"description,omitempty"`      //"string"
	//AssetType            string  `json:"assetType,omitempty"`        //"string"
	BondMaturityDate string  `json:"bondMaturityDate,omitempty"` //"string"
	BondInterestRate float64 `json:"bondInterestRate,omitempty"`
	//================================================================

	// watchlist
	//================================================================
	//Symbol      string `json:"symbol,omitempty"`      //"string"
	//AssetType   string `json:"assetType,omitempty"`   //"'EQUITY' or 'OPTION' or 'MUTUAL_FUND' or 'FIXED_INCOME' or 'INDEX'"
	//Description string `json:"description,omitempty"` //"string"
	//================================================================
}

// FundamentalData https://developer.tdameritrade.com/instruments/apis
type FundamentalData struct {
	Symbol              string  `json:"symbol,omitempty"`
	High52              float64 `json:"high52,omitempty"`
	Low52               float64 `json:"low52,omitempty"`
	DividendAmount      float64 `json:"dividendAmount,omitempty"`
	DividendYield       float64 `json:"dividendYield,omitempty"`
	DividendDate        string  `json:"dividendDate,omitempty"`
	PeRatio             float64 `json:"peRatio,omitempty"`
	PegRatio            float64 `json:"pegRatio,omitempty"`
	PbRatio             float64 `json:"pbRatio,omitempty"`
	PrRatio             float64 `json:"prRatio,omitempty"`
	PcfRatio            float64 `json:"pcfRatio,omitempty"`
	GrossMarginTTM      float64 `json:"grossMarginTTM,omitempty"`
	GrossMarginMRQ      float64 `json:"grossMarginMRQ,omitempty"`
	NetProfitMarginTTM  float64 `json:"netProfitMarginTTM,omitempty"`
	NetProfitMarginMRQ  float64 `json:"netProfitMarginMRQ,omitempty"`
	OperatingMarginTTM  float64 `json:"operatingMarginTTM,omitempty"`
	OperatingMarginMRQ  float64 `json:"operatingMarginMRQ,omitempty"`
	ReturnOnEquity      float64 `json:"returnOnEquity,omitempty"`
	ReturnOnAssets      float64 `json:"returnOnAssets,omitempty"`
	ReturnOnInvestment  float64 `json:"returnOnInvestment,omitempty"`
	QuickRatio          float64 `json:"quickRatio,omitempty"`
	CurrentRatio        float64 `json:"currentRatio,omitempty"`
	InterestCoverage    float64 `json:"interestCoverage,omitempty"`
	TotalDebtToCapital  float64 `json:"totalDebtToCapital,omitempty"`
	LtDebtToEquity      float64 `json:"ltDebtToEquity,omitempty"`
	TotalDebtToEquity   float64 `json:"totalDebtToEquity,omitempty"`
	EpsTTM              float64 `json:"epsTTM,omitempty"`
	EpsChangePercentTTM float64 `json:"epsChangePercentTTM,omitempty"`
	EpsChangeYear       float64 `json:"epsChangeYear,omitempty"`
	EpsChange           float64 `json:"epsChange,omitempty"`
	RevChangeYear       float64 `json:"revChangeYear,omitempty"`
	RevChangeTTM        float64 `json:"revChangeTTM,omitempty"`
	RevChangeIn         float64 `json:"revChangeIn,omitempty"`
	SharesOutstanding   float64 `json:"sharesOutstanding,omitempty"`
	MarketCapFloat      float64 `json:"marketCapFloat,omitempty"`
	MarketCap           float64 `json:"marketCap,omitempty"`
	BookValuePerShare   float64 `json:"bookValuePerShare,omitempty"`
	ShortIntToFloat     float64 `json:"shortIntToFloat,omitempty"`
	ShortIntDayToCover  float64 `json:"shortIntDayToCover,omitempty"`
	DivGrowthRate3Year  float64 `json:"divGrowthRate3Year,omitempty"`
	DividendPayAmount   float64 `json:"dividendPayAmount,omitempty"`
	DividendPayDate     string  `json:"dividendPayDate,omitempty"`
	Beta                float64 `json:"beta,omitempty"`
	Vol1DayAvg          float64 `json:"vol1DayAvg,omitempty"`
	Vol10DayAvg         float64 `json:"vol10DayAvg,omitempty"`
	Vol3MonthAvg        float64 `json:"vol3MonthAvg,omitempty"`
}
