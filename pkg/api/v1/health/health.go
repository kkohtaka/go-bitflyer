// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package health

import (
	"net/http"

	"github.com/google/go-querystring/query"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
)

type Request struct {
	ProductCode markets.ProductCode `json:"product_code" url:"product_code"`
}

type Response struct {
	Status Status `json:"status"`
}

type Status string

const (
	APIPath string = "/v1/gethealth"

	Normal    Status = "NORMAL"
	Busy      Status = "BUSY"
	VeryBusy  Status = "VERY BUSY"
	SuperBusy Status = "SUPER BUSY"
	Stop      Status = "STOP"
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
