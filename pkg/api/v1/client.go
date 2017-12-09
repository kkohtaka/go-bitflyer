// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package v1

import (
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/board"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	httpclient "github.com/kkohtaka/go-bitflyer/pkg/httpclient"
)

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
	var resp markets.Response
	err := httpclient.Get(NewAPI(c, markets.APIPath), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) Board(req *board.Request) (*board.Response, error) {
	var resp board.Response
	err := httpclient.Get(NewAPI(c, board.APIPath), req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
