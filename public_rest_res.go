package myhashkeyapi

type PublicRestExchangeInfoRes struct {
	Timezone      string                         `json:"timezone"`
	ServerTime    string                         `json:"serverTime"`
	BrokerFilters []interface{}                  `json:"brokerFilters"`
	Symbols       []PublicRestExchangeInfoResRow `json:"symbols"`
}

type PublicRestExchangeInfoResRow struct {
	Symbol             string                   `json:"symbol"`
	SymbolName         string                   `json:"symbolName"`
	Status             string                   `json:"status"`
	BaseAsset          string                   `json:"baseAsset"`
	BaseAssetName      string                   `json:"baseAssetName"`
	BaseAssetPrecision string                   `json:"baseAssetPrecision"`
	QuoteAsset         string                   `json:"quoteAsset"`
	QuoteAssetName     string                   `json:"quoteAssetName"`
	QuotePrecision     string                   `json:"quotePrecision"`
	RetailAllowed      bool                     `json:"retailAllowed"`
	PiAllowed          bool                     `json:"piAllowed"`
	CorporateAllowed   bool                     `json:"corporateAllowed"`
	OmnibusAllowed     bool                     `json:"omnibusAllowed"`
	IcebergAllowed     bool                     `json:"icebergAllowed"`
	IsAggregate        bool                     `json:"isAggregate"`
	AllowMargin        bool                     `json:"allowMargin"`
	Filters            []map[string]interface{} `json:"filters"`
}
