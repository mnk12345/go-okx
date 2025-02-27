package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetBorrowInfo(param *GetBorrowInfoParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/finance/savings/lending-rate-summary",
		Method: api.MethodGet,
		Param:  param,
	}, &GetBorrowInfoResponse{}
}

type GetBorrowInfoParam struct {
	Ccy string `url:"ccy,omitempty"`
}

type GetBorrowInfoResponse struct {
	api.Response
	Data []BorrowInfo `json:"data"`
}

type BorrowInfo struct {
	Ccy       string `json:"ccy"`
	AvgAmt    string `json:"avgAmt"`
	AvgAmtUsd string `json:"avgAmtUsd"`
	AvgRate   string `json:"avgRate"`
	PreRate   string `json:"preRate"`
	EstRate   string `json:"estRate"`
}
