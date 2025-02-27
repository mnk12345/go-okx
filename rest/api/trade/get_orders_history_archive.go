package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetOrdersHistoryArchive(param *GetOrdersHistoryArchiveQueryParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/trade/orders-history-archive",
		Method: api.MethodGet,
		Param:  param,
	}, &GetOrderResponse{}
}

type GetOrdersHistoryArchiveQueryParam struct {
	api.PagingParam
	InstType string `url:"instType,omitempty"`
	Uly      string `url:"uly,omitempty"`
	InstId   string `url:"instId,omitempty"`
	OrdType  string `url:"ordType,omitempty"`
	State    string `url:"state,omitempty"`
	Category string `url:"category,omitempty"`
	Begin    string `url:"begin,omitempty"`
	End      string `url:"end,omitempty"`
}
