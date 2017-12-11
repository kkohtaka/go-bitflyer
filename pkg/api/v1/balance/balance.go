// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package balance

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/types"
)

type Request struct{}

type Response []Asset

type Asset struct {
	CurrencyCode types.CurrencyCode `json:"currency_code"`
	Amount       float64            `json:"amount"`
	Available    float64            `json:"available"`
}

const (
	APIPath string = "/v1/me/getbalance"
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
