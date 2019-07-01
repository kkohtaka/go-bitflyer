package childorders

import (
	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/sendchildorder"
	"net/http"
	"time"
)

type Request struct {
	ProductCode     markets.ProductCode
	Count           int64
	Before          int64
	After           int64
	ChildOrderState OrderState
	ChildOrderId    string
	ParentOrderId   string
}

type OrderState string

const (
	ACTIVE    OrderState = "ACTIVE"
	COMPLETED OrderState = "COMPLETED"
	CANCELED  OrderState = "CANCELED"
	EXPIRED   OrderState = "EXPIRED"
	REJECTED  OrderState = "REJECTED"
)

type Response []Order

type Order struct {
	Id                     int64                    `json:"id"`
	ChildOrderId           string                   `json:"child_order_id"`
	ProductCode            markets.ProductCode      `json:"product_code"`
	Side                   sendchildorder.Side      `json:"side"`
	ChildOrderType         sendchildorder.OrderType `json:"child_order_type"`
	Price                  float64                  `json:"price"`
	AveragePrice           float64                  `json:"average_price"`
	Size                   float64                  `json:"size"`
	ChildOrderState        OrderState               `json:"child_order_state"`
	ExpireDate             time.Time                `json:"expire_date"`
	ChildOrderDate         time.Time                `json:"child_order_date"`
	ChildOrderAcceptanceId string                   `json:"child_order_acceptance_id"`
	OutstandingSize        float64                  `json:"outstanding_size"`
	CancelSize             float64                  `json:"cancel_size"`
	ExecutedSize           float64                  `json:"executed_size"`
	TotalCommission        float64                  `json:"total_commission"`
}

const (
	APIPath = "/v1/me/getchildorders"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
