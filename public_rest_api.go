package myhashkeyapi

type PublicRest int

const (
	PublicRestExchangeInfo PublicRest = iota //GET 获取交易规范
)

var PublicRestAPIMap = map[PublicRest]string{
	PublicRestExchangeInfo: "/api/v1/exchangeInfo", //GET 获取交易规范
}

// hashkey PublicRestExchangeInfo PublicRest接口 GET 获取交易规范
func (client *PublicRestClient) NewExchangeInfo() *PublicRestExchangeInfoAPI {
	return &PublicRestExchangeInfoAPI{
		client: client,
		req:    &PublicRestExchangeInfoReq{},
	}
}
func (api *PublicRestExchangeInfoAPI) Do() (*PublicRestExchangeInfoRes, error) {
	url := hashkeyHandlerRequestAPIWithPathQueryParam(REST, api.req, PublicRestAPIMap[PublicRestExchangeInfo])
	if api.req.Symbol != nil && *api.req.Symbol != "" {
		res, err := hashkeyCallAPI[PublicRestExchangeInfoResRow](api.client.c, url, NIL_REQBODY, GET)
		if err != nil {
			return nil, err
		}
		return &PublicRestExchangeInfoRes{Symbols: []PublicRestExchangeInfoResRow{*res}}, nil
	} else {
		return hashkeyCallAPI[PublicRestExchangeInfoRes](api.client.c, url, NIL_REQBODY, GET)
	}
}
