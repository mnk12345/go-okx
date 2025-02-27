package trade

import (
	"github.com/iaping/go-okx/rest/api"
)

func NewGetInstruments(param *GetInstrumentsParam) (api.IRequest, api.IResponse) {
	return &api.Request{
		Path:   "/api/v5/account/instruments",
		Method: api.MethodGet,
		Param:  param,
	}, &GetInstrumentsResponse{}
}

type GetInstrumentsParam struct {
	InstType   string `url:"instType"`
	Uly        string `url:"uly,omitempty"`
	InstId     string `url:"instId,omitempty"`
	InstFamily string `url:"instFamily,omitempty"`
}

type GetInstrumentsResponse struct {
	api.Response
	Data []Instrument `json:"data"`
}

type Instrument struct {
	InstType       string `json:"instType"`
	InstId         string `json:"instId"`
	Uly            string `url:"uly"`
	InstFamily     string `url:"instFamily"`
	BaseCcy        string `json:"baseCcy"`
	QuoteCcy       string `json:"quoteCcy"`
	SettleCcy      string `json:"settleCcy"`
	CtVal          string `json:"ctVal"`
	CtMult         string `json:"ctMult"`
	CtValCcy       string `json:"ctValCcy"`
	OptType        string `json:"optType"`
	Stk            string `json:"stk"`
	ListTime       string `json:"listTime"`
	AuctionEndTime string `json:"auctionEndTime"`
	ExpTime        string `json:"expTime"`
	Lever          string `json:"lever"`
	TickSz         string `json:"tickSz"`
	LotSz          string `json:"lotSz"`
	MinSz          string `json:"minSz"`
	CtType         string `json:"ctType"`
	State          string `json:"state"`
	RuleType       string `json:"ruleType"`
	MaxLmtSz       string `json:"maxLmtSz"`
	MaxMktSz       string `json:"maxMktSz"`
	MaxLmtAmt      string `json:"maxLmtAmt"`
	MaxMktAmt      string `json:"maxMktAmt"`
	MaxTwapSz      string `json:"maxTwapSz"`
	MaxIcebergSz   string `json:"maxIcebergSz"`
	MaxTriggerSz   string `json:"maxTriggerSz"`
	MaxStopSz      string `json:"maxStopSz"`
}
