package myhashkeyapi

type PublicRestExchangeInfoReq struct {
	Symbol *string `json:"symbol"` //symbol	STRING		Symbol Name. e.g: "BTCUSD", "ETHUSDC"
}

type PublicRestExchangeInfoAPI struct {
	client *PublicRestClient
	req    *PublicRestExchangeInfoReq
}

// symbol	STRING		Symbol Name. e.g: "BTCUSD", "ETHUSDC"
func (api *PublicRestExchangeInfoAPI) Symbol(symbol string) *PublicRestExchangeInfoAPI {
	api.req.Symbol = GetPointer(symbol)
	return api
}
