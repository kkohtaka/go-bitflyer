package cancelchildorder

import (
	"encoding/json"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/pkg/errors"
	"net/http"
)

type Request struct {
	ProductCode            markets.ProductCode `json:"product_code"`
	ChildOrderId           string              `json:"child_order_id"`
	ChildOrderAcceptanceId string              `json:"child_order_acceptance_id"`
}

type Response struct{}

const (
	APIPath = "/v1/me/cancelchildorder"
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
