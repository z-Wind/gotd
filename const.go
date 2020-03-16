package gotd

const (
	AccountFieldPositions = "positions"
	AccountFieldOrders    = "orders"

	InstrumentsProjectionSymbolSearch = "symbol-search"
	InstrumentsProjectionSymbolRegex  = "symbol-regex"
	InstrumentsProjectionDescSearch   = "desc-search"
	InstrumentsProjectionDescRegex    = "desc-regex"
	InstrumentsProjectionFundamental  = "fundamental"

	HoursMarketsEQUITY = "EQUITY"
	HoursMarketsOPTION = "OPTION"
	HoursMarketsFUTURE = "FUTURE"
	HoursMarketsBOND   = "BOND"
	HoursMarketsFOREX  = "FOREX"

	MoversIndexCOMPX    = "$COMPX"
	MoversIndexDJI      = "$DJI"
	MoversIndexSPXX     = "$SPX.X"
	MoversDirectionUp   = "up"
	MoversDirectionDown = "down"
	MoversChangePercent = "percent"
	MoversChangeValue   = "value"

	OptionChainContractTypeCALL   = "CALL"
	OptionChainContractTypePUT    = "PUT"
	OptionChainContractTypeALL    = "ALL"
	OptionChainStrategySINGLE     = "SINGLE"
	OptionChainStrategyANALYTICAL = "ANALYTICAL"
	OptionChainStrategyCOVERED    = "COVERED"
	OptionChainStrategyVERTICAL   = "VERTICAL"
	OptionChainStrategyCALENDAR   = "CALENDAR"
	OptionChainStrategySTRANGLE   = "STRANGLE"
	OptionChainStrategySTRADDLE   = "STRADDLE"
	OptionChainStrategyBUTTERFLY  = "BUTTERFLY"
	OptionChainStrategyCONDOR     = "CONDOR"
	OptionChainStrategyDIAGONAL   = "DIAGONAL"
	OptionChainStrategyCOLLAR     = "COLLAR"
	OptionChainStrategyROLL       = "ROLL"
	OptionChainRangeITM           = "ITM"
	OptionChainRangeNTM           = "NTM"
	OptionChainRangeOTM           = "OTM"
	OptionChainRangeSAK           = "SAK"
	OptionChainRangeSBK           = "SBK"
	OptionChainRangeSNK           = "SNK"
	OptionChainRangeALL           = "ALL"
	OptionChainExpMonthJAN        = "JAN"
	OptionChainExpMonthFEB        = "FEB"
	OptionChainExpMonthMAR        = "MAR"
	OptionChainExpMonthAPR        = "APR"
	OptionChainExpMonthMAY        = "MAY"
	OptionChainExpMonthJUN        = "JUN"
	OptionChainExpMonthJULY       = "JULY"
	OptionChainExpMonthAUG        = "AUG"
	OptionChainExpMonthSEP        = "SEP"
	OptionChainExpMonthOCT        = "OCT"
	OptionChainExpMonthNOV        = "NOV"
	OptionChainExpMonthDEC        = "DEC"
	OptionChainExpMonthALL        = "ALL"
	OptionChainTypeS              = "S"
	OptionChainTypeNS             = "NS"
	OptionChainTypeALL            = "ALL"

	PriceHistoryPeriodTypeDay        = "day"
	PriceHistoryPeriodTypeMonth      = "month"
	PriceHistoryPeriodTypeYear       = "year"
	PriceHistoryPeriodTypeYtd        = "ytd"
	PriceHistoryFrequencyTypeMinute  = "minute"
	PriceHistoryFrequencyTypeDaily   = "daily"
	PriceHistoryFrequencyTypeWeekly  = "weekly"
	PriceHistoryFrequencyTypeMonthly = "monthly"

	TransactionsKindALL             = "ALL"
	TransactionsKindTRADE           = "TRADE"
	TransactionsKindBUYONLY         = "BUY_ONLY"
	TransactionsKindSELLONLY        = "SELL_ONLY"
	TransactionsKindCASHInOrCASHOut = "CASH_IN_OR_CASH_OUT"
	TransactionsKindCHECKING        = "CHECKING"
	TransactionsKindDIVIDEND        = "DIVIDEND"
	TransactionsKindINTEREST        = "INTEREST"
	TransactionsKindOTHER           = "OTHER"
	TransactionsKindADVISORFEES     = "ADVISOR_FEES"

	UserPrincipalsFieldsStreamerSubscriptionKeys = "streamerSubscriptionKeys"
	UserPrincipalsFieldsStreamerConnectionInfo   = "streamerConnectionInfo"
	UserPrincipalsFieldsPreferences              = "preferences"
	UserPrincipalsFieldsSurrogateIds             = "surrogateIds"

	OrdersStatusAWAITINGPARENTORDER      = "AWAITING_PARENT_ORDER"
	OrdersStatusAWAITINGCONDITION        = "AWAITING_CONDITION"
	OrdersStatusAWAITINGMANUALREVIEW     = "AWAITING_MANUAL_REVIEW"
	OrdersStatusACCEPTED                 = "ACCEPTED"
	OrdersStatusAWAITINGUROUT            = "AWAITING_UR_OUT"
	OrdersStatusPENDINGACTIVATION        = "PENDING_ACTIVATION"
	OrdersStatusQUEUED                   = "QUEUED"
	OrdersStatusWORKING                  = "WORKING"
	OrdersStatusREJECTED                 = "REJECTED"
	OrdersStatusPENDINGCANCEL            = "PENDING_CANCEL"
	OrdersStatusCANCELED                 = "CANCELED"
	OrdersStatusPENDINGREPLACE           = "PENDING_REPLACE"
	OrdersStatusREPLACED                 = "REPLACED"
	OrdersStatusFILLED                   = "FILLED"
	OrdersStatusEXPIRED                  = "EXPIRED"
	OrderInstructionBuy                  = "BUY"
	OrderInstructionSELL                 = "SELL"
	OrderAssetTypeEQUITY                 = "EQUITY"
	OrderAssetTypeOPTION                 = "OPTION"
	OrderAssetTypeINDEX                  = "INDEX"
	OrderAssetTypeMUTUALFUND             = "MUTUAL_FUND"
	OrderAssetTypeCASHEQUIVALENT         = "CASH_EQUIVALENT"
	OrderAssetTypeFIXEDINCOME            = "FIXED_INCOME"
	OrderAssetTypeCURRENCYOrderAssetType = "CURRENCYOrderAssetType"

	PreferencesTimeoutFIFTYFIVEMINUTES = "FIFTY_FIVE_MINUTES"
	PreferencesTimeoutTWOHOURS         = "TWO_HOURS"
	PreferencesTimeoutFOURHOURS        = "FOUR_HOURS"
	PreferencesTimeoutEIGHTHOURS       = "EIGHT_HOURS"
)
