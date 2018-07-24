// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package v1

import (
	"github.com/kkohtaka/go-bitflyer/pkg/api/auth"
	"github.com/kkohtaka/go-bitflyer/pkg/api/httpclient"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/addresses"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/balance"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/bankaccounts"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/board"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/chats"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/coinins"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/coinouts"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/collateral"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/collateralaccounts"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/executions"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/health"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/permissions"
	"github.com/kkohtaka/go-bitflyer/pkg/api/v1/ticker"
	"github.com/pkg/errors"
)

const (
	APIHost string = "https://api.bitflyer.jp"
)

type Client struct {
	Host string

	AuthConfig *auth.AuthConfig
}

type ClientOpts struct {
	AuthConfig *auth.AuthConfig
}

func NewClient(opts *ClientOpts) *Client {
	return &Client{
		Host:       APIHost,
		AuthConfig: opts.AuthConfig,
	}
}

func (c *Client) APIHost() string {
	return c.Host
}

// Public APIs

func (c *Client) Markets(req *markets.Request) (*markets.Response, error) {
	var resp markets.Response
	err := httpclient.New().Request(NewAPI(c, markets.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Board(req *board.Request) (*board.Response, error) {
	var resp board.Response
	err := httpclient.New().Request(NewAPI(c, board.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Ticker(req *ticker.Request) (*ticker.Response, error) {
	var resp ticker.Response
	err := httpclient.New().Request(NewAPI(c, ticker.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Executions(req *executions.Request) (*executions.Response, error) {
	var resp executions.Response
	err := httpclient.New().Request(NewAPI(c, executions.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Health(req *health.Request) (*health.Response, error) {
	var resp health.Response
	err := httpclient.New().Request(NewAPI(c, health.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Chats(req *chats.Request) (*chats.Response, error) {
	var resp chats.Response
	err := httpclient.New().Request(NewAPI(c, chats.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

// Private APIs

func (c *Client) Permissions(req *permissions.Request) (*permissions.Response, error) {
	var resp permissions.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, permissions.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Balance(req *balance.Request) (*balance.Response, error) {
	var resp balance.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, balance.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Collateral(req *collateral.Request) (*collateral.Response, error) {
	var resp collateral.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, collateral.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) CollateralAccounts(req *collateralaccounts.Request) (*collateralaccounts.Response, error) {
	var resp collateralaccounts.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, collateralaccounts.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Addresses(req *addresses.Request) (*addresses.Response, error) {
	var resp addresses.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, addresses.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Coinins(req *coinins.Request) (*coinins.Response, error) {
	var resp coinins.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, coinins.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) Coinouts(req *coinouts.Request) (*coinouts.Response, error) {
	var resp coinouts.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, coinouts.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}

func (c *Client) BankAccounts(req *bankaccounts.Request) (*bankaccounts.Response, error) {
	var resp bankaccounts.Response
	err := httpclient.New().Auth(c.AuthConfig).Request(NewAPI(c, bankaccounts.APIPath), req, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "send HTTP request")
	}
	return &resp, nil
}
