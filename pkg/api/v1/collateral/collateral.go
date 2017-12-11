// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package collateral

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct{}

type Response struct {
	Collateral        float64 `json:"collateral"`         // This is the amount of deposited in Japanese Yen.
	OpenPositionPNL   float64 `json:"open_position_pnl"`  // This is the profit or loss from valuation.
	RequireCollateral float64 `json:"require_collateral"` // This is the current required margin.
	KeepRate          float64 `json:"keep_rate"`          // This is the current maintenance margin.
}

const (
	APIPath string = "/v1/me/getcollateral"
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
