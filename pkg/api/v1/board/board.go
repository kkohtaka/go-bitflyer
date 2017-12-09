// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package board

import (
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
)

type Book struct {
	Price float64 `json:"price"`
	Size  float64 `json:"size"`
}

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	MidPrice float64 `json:"mid_price"`
	Bids     []Book  `json:"bids"`
	Asks     []Book  `json:"asks"`
}

const (
	APIPath string = "/v1/getboard"
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

func (b Book) String() string {
	return fmt.Sprintf("%g (x %g)", b.Price, b.Size)
}
