package account

import "github.com/iaping/go-okx/rest/api"

func NewGetBills(param *GetBillsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/bills",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBillsResponse{}
}

type GetBillsParam struct {
	api.PagingParam
	InstType string `url:"instType,omitempty"`
	InstId   string `url:"instId,omitempty"`
	Ccy      string `url:"ccy,omitempty"`
	MgnMode  string `url:"mgnMode,omitempty"`
	CtType   string `url:"ctType,omitempty"`
	Type     string `url:"type,omitempty"`
	SubType  string `url:"subType,omitempty"`
	After    string `url:"after,omitempty"`
	Before   string `url:"before,omitempty"`
	Begin    string `url:"begin,omitempty"`
	End      string `url:"end,omitempty"`
	Limit    string `url:"limit,omitempty"`
}

type GetBillsResponse struct {
	api.Response
	Data []Bills `json:"data"`
}

type Bills struct {
	InstType  string `json:"instType"`
	BillId    string `json:"billId"`
	Type      string `json:"type"`
	SubType   string `json:"subType"`
	Ts        int64  `json:"ts,string"`
	BalChg    string `json:"balChg"`
	PosBalChg string `json:"posBalChg"`
	Bal       string `json:"bal"`
	PosBal    string `json:"posBal"`
	Sz        string `json:"sz"`
	Ccy       string `json:"ccy"`
	Pnl       string `json:"pnl"`
	Fee       string `json:"fee"`
	MgnMode   string `json:"mgnMode"`
	InstId    string `json:"instId"`
	OrdId     string `json:"ordId"`
	ExecType  string `json:"execType"`
	From      string `json:"from"`
	To        string `json:"to"`
	Notes     string `json:"notes"`
}
