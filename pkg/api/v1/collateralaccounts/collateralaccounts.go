// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package collateralaccounts

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct{}

type Response []Account

type Account struct {
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

const (
	APIPath string = "/v1/me/getcollateralaccounts"
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
