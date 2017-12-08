// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package v1

import "github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"

const (
	APIHost string = "https://api.bitflyer.jp/v1"
)

type Client struct {
	Host string
}

func NewClient() *Client {
	return &Client{
		Host: APIHost,
	}
}

func (c *Client) APIHost() string {
	return c.Host
}

func (c *Client) Markets(req *markets.Request) (*markets.Response, error) {
	return markets.NewAPI(c).Execute(req)
}
