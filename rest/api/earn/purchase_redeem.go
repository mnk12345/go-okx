package trade

import "github.com/iaping/go-okx/rest/api"

func NewPurchaseRedeem(param *PurchaseRedeemParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/finance/savings/purchase-redempt",
		Method: api.MethodPost,
		Param:  param,
	}, &PurchaseRedeemResponse{}
}

type PurchaseRedeemParam struct {
	Ccy  string `json:"ccy"`
	Amt  string `json:"amt"`
	Side string `json:"side"`
	Rate string `json:"rate"`
}

type PurchaseRedeemResponse struct {
	api.Response
	Data []Lend `json:"data"`
}

type Lend struct {
	Ccy  string `json:"ccy"`
	Amt  string `json:"amt"`
	Side string `json:"side"`
	Rate string `json:"rate"`
}
