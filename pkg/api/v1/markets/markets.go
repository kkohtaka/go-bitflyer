// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package markets

import (
	"fmt"
	"net/http"
)

type Request struct{}

type Response []Market

type Market struct {
	ProductCode ProductCode `json:"product_code"`
	Alias       ProductCode `json:"alias"`
}

type ProductCode string

const (
	APIPath string = "/v1/getmarkets"
)

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	return ""
}

func (req *Request) Payload() []byte {
	return nil
}

func (m Market) String() string {
	if m.Alias == "" {
		return string(m.ProductCode)
	}
	return fmt.Sprintf("%s (%s)", m.ProductCode, m.Alias)
}
