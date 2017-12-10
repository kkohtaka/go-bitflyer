// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package permissions

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct{}

type Response []Permission

type Permission string

type Status string

const (
	APIPath string = "/v1/me/getpermissions"
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
