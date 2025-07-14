package earn

import (
	"encoding/json"

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
	Ccy       json.RawMessage `json:"ccy"`
	AvgAmt    json.RawMessage `json:"avgAmt"`
	AvgAmtUsd json.RawMessage `json:"avgAmtUsd"`
	AvgRate   json.RawMessage `json:"avgRate"`
	PreRate   json.RawMessage `json:"preRate"`
	EstRate   json.RawMessage `json:"estRate"`
}
