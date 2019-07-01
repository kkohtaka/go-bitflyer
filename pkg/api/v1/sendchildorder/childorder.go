package sendchildorder

import (
	"encoding/json"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/pkg/errors"
	"net/http"
)

type Request struct {
	ProductCode    markets.ProductCode `json:"product_code"`
	ChildOrderType OrderType           `json:"child_order_type"`
	Side           Side                `json:"side"`
	Price          float64             `json:"price"`
	Size           float64             `json:"size"`
	MinuteToExpire int64               `json:"minute_to_expire"`
	TimeInForce    ExecutiveCondition  `json:"time_in_force"`
}

type Response struct {
	ChildOrderAcceptanceId string `json:"child_order_acceptance_id"`
}

type OrderType string

const (
	Limit  OrderType = "LIMIT"
	Market OrderType = "MARKET"
)

type Side string

const (
	Buy  Side = "BUY"
	Sell Side = "SELL"
)

type ExecutiveCondition string

const (
	GTC ExecutiveCondition = "GTC"
	IoC ExecutiveCondition = "IOC"
	FoK ExecutiveCondition = "FOK"
)

const (
	APIPath string = "/v1/me/sendchildorder"
)

func (req *Request) Method() string {
	return http.MethodPost
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	body, err := json.Marshal(*req)
	if err != nil {
		panic(errors.Wrap(err, "json serialization error"))
	}
	return body
}
