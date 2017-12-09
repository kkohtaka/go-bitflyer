// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package chats

import "github.com/google/go-querystring/query"

type Request struct {
	FromDate string `json:"from_date,omitempty"` // TODO: Treat timestamp as time.Time
}

type Response []Chat

type Chat struct {
	Nickname string `json:"nickname"`
	Message  string `json:"message"`
	Date     string `json:"date"` // TODO: Treat timestamp as time.Time
}

const (
	APIPath string = "getchats"
)

func (req *Request) Query() string {
	values, _ := query.Values(req)
	return values.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
