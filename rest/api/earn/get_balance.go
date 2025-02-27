package earn

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetBalance(param *GetBalanceParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/finance/savings/balance",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBalanceResponse{}
}

type GetBalanceParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetBalanceResponse struct {
	api.Response
	Data []EarnBalance `json:"data"`
}

type EarnBalance struct {
	Ccy        string `json:"ccy"`
	Amt        string `json:"amt"`
	Earnings   string `json:"earnings"`
	Rate       string `json:"rate"`
	LoanAmt    string `json:"loanAmt"`
	PendingAmt string `json:"pendingAmt"`
}
