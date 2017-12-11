// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package executions

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/types"
)

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`

	Pagination types.Pagination `json:",inline"`
}

type Response []Execution

type Execution struct {
	ID                         int     `json:"id"`
	Side                       string  `json:"side"`
	Price                      float64 `json:"price"`
	Size                       float64 `json:"size"`
	ExecDate                   string  `json:"exec_date"` // TODO: Treat timestamp as time.Time
	BuyChildOrderAcceptanceID  string  `json:"buy_child_order_acceptance_id"`
	SellChildOrderAcceptanceID string  `json:"sell_child_order_acceptance_id"`
}

const (
	APIPath string = "/v1/getexecutions"
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
