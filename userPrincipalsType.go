package gotd

// Preferences https://developer.tdameritrade.com/user-principal/apis
type Preferences struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	ExpressTrading                   bool    `json:"expressTrading,omitempty"`
	DirectOptionsRouting             bool    `json:"directOptionsRouting,omitempty"`
	DirectEquityRouting              bool    `json:"directEquityRouting,omitempty"`
	DefaultEquityOrderLegInstruction string  `json:"defaultEquityOrderLegInstruction,omitempty"` //"'BUY' or 'SELL' or 'BUY_TO_COVER' or 'SELL_SHORT' or 'NONE'"
	DefaultEquityOrderType           string  `json:"defaultEquityOrderType,omitempty"`           //"'MARKET' or 'LIMIT' or 'STOP' or 'STOP_LIMIT' or 'TRAILING_STOP' or 'MARKET_ON_CLOSE' or 'NONE'"
	DefaultEquityOrderPriceLinkType  string  `json:"defaultEquityOrderPriceLinkType,omitempty"`  //"'VALUE' or 'PERCENT' or 'NONE'"
	DefaultEquityOrderDuration       string  `json:"defaultEquityOrderDuration,omitempty"`       //"'DAY' or 'GOOD_TILL_CANCEL' or 'NONE'"
	DefaultEquityOrderMarketSession  string  `json:"defaultEquityOrderMarketSession,omitempty"`  //"'AM' or 'PM' or 'NORMAL' or 'SEAMLESS' or 'NONE'"
	DefaultEquityQuantity            float64 `json:"defaultEquityQuantity,omitempty"`
	MutualFundTaxLotMethod           string  `json:"mutualFundTaxLotMethod,omitempty"`    //"'FIFO' or 'LIFO' or 'HIGH_COST' or 'LOW_COST' or 'MINIMUM_TAX' or 'AVERAGE_COST' or 'NONE'"
	OptionTaxLotMethod               string  `json:"optionTaxLotMethod,omitempty"`        //"'FIFO' or 'LIFO' or 'HIGH_COST' or 'LOW_COST' or 'MINIMUM_TAX' or 'AVERAGE_COST' or 'NONE'"
	EquityTaxLotMethod               string  `json:"equityTaxLotMethod,omitempty"`        //"'FIFO' or 'LIFO' or 'HIGH_COST' or 'LOW_COST' or 'MINIMUM_TAX' or 'AVERAGE_COST' or 'NONE'"
	DefaultAdvancedToolLaunch        string  `json:"defaultAdvancedToolLaunch,omitempty"` //"'TA' or 'N' or 'Y' or 'TOS' or 'NONE' or 'CC2'"
	AuthTokenTimeout                 string  `json:"authTokenTimeout,omitempty"`          //"'FIFTY_FIVE_MINUTES' or 'TWO_HOURS' or 'FOUR_HOURS' or 'EIGHT_HOURS'"
}

// Key https://developer.tdameritrade.com/user-principal/apis
type Key struct {
	Key string `json:"key,omitempty"` //"string"
}

// SubscriptionKey https://developer.tdameritrade.com/user-principal/apis
type SubscriptionKey struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	Keys []*Key `json:"keys,omitempty"`
}

// StreamerInfo https://developer.tdameritrade.com/user-principal/apis
type StreamerInfo struct {
	StreamerBinaryURL string `json:"streamerBinaryUrl,omitempty"` //"string"
	StreamerSocketURL string `json:"streamerSocketUrl,omitempty"` //"string"
	Token             string `json:"token,omitempty"`             //"string"
	TokenTimestamp    string `json:"tokenTimestamp,omitempty"`    //"string"
	UserGroup         string `json:"userGroup,omitempty"`         //"string"
	AccessLevel       string `json:"accessLevel,omitempty"`       //"string"
	ACL               string `json:"acl,omitempty"`               //"string"
	AppID             string `json:"appId,omitempty"`             //"string"
}

// Quotes https://developer.tdameritrade.com/user-principal/apis
type Quotes struct {
	IsNyseDelayed   bool `json:"isNyseDelayed,omitempty"`
	IsNasdaqDelayed bool `json:"isNasdaqDelayed,omitempty"`
	IsOpraDelayed   bool `json:"isOpraDelayed,omitempty"`
	IsAmexDelayed   bool `json:"isAmexDelayed,omitempty"`
	IsCmeDelayed    bool `json:"isCmeDelayed,omitempty"`
	IsIceDelayed    bool `json:"isIceDelayed,omitempty"`
	IsForexDelayed  bool `json:"isForexDelayed,omitempty"`
}

// Authorizations https://developer.tdameritrade.com/user-principal/apis
type Authorizations struct {
	Apex               bool   `json:"apex,omitempty"`
	LevelTwoQuotes     bool   `json:"levelTwoQuotes,omitempty"`
	StockTrading       bool   `json:"stockTrading,omitempty"`
	MarginTrading      bool   `json:"marginTrading,omitempty"`
	StreamingNews      bool   `json:"streamingNews,omitempty"`
	OptionTradingLevel string `json:"optionTradingLevel,omitempty"` //"'COVERED' or 'FULL' or 'LONG' or 'SPREAD' or 'NONE'"
	StreamerAccess     bool   `json:"streamerAccess,omitempty"`
	AdvancedMargin     bool   `json:"advancedMargin,omitempty"`
	ScottradeAccount   bool   `json:"scottradeAccount,omitempty"`
}

// AccountU https://developer.tdameritrade.com/user-principal/apis
type AccountU struct {
	AccountID         string            `json:"accountId,omitempty"`         //"string"
	Description       string            `json:"description,omitempty"`       //"string"
	DisplayName       string            `json:"displayName,omitempty"`       //"string"
	AccountCdDomainID string            `json:"accountCdDomainId,omitempty"` //"string"
	Company           string            `json:"company,omitempty"`           //"string"
	Segment           string            `json:"segment,omitempty"`           //"string"
	SurrogateIds      map[string]string `json:"surrogateIds,omitempty"`      //"object"
	Preferences       Preferences       `json:"preferences,omitempty"`
	ACL               string            `json:"acl,omitempty"` //"string"
	Authorizations    Authorizations    `json:"authorizations,omitempty"`
}

// UserPrincipal https://developer.tdameritrade.com/user-principal/apis
type UserPrincipal struct {
	// ServerResponse contains the HTTP response code and headers from the
	// server.
	ServerResponse `json:"-"`

	AuthToken                string           `json:"authToken,omitempty"`           //"string"
	UserID                   string           `json:"userId,omitempty"`              //"string"
	UserCdDomainID           string           `json:"userCdDomainId,omitempty"`      //"string"
	PrimaryAccountID         string           `json:"primaryAccountId,omitempty"`    //"string"
	LastLoginTime            string           `json:"lastLoginTime,omitempty"`       //"string"
	TokenExpirationTime      string           `json:"tokenExpirationTime,omitempty"` //"string"
	LoginTime                string           `json:"loginTime,omitempty"`           //"string"
	AccessLevel              string           `json:"accessLevel,omitempty"`         //"string"
	StalePassword            bool             `json:"stalePassword,omitempty"`
	StreamerInfo             *StreamerInfo    `json:"streamerInfo,omitempty"`
	ProfessionalStatus       string           `json:"professionalStatus,omitempty"` //"'PROFESSIONAL' or 'NON_PROFESSIONAL' or 'UNKNOWN_STATUS'"
	Quotes                   *Quotes          `json:"quotes,omitempty"`
	StreamerSubscriptionKeys *SubscriptionKey `json:"streamerSubscriptionKeys,omitempty"`
	Accounts                 []*AccountU      `json:"accounts,omitempty"`
}
