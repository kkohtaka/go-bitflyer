// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package board

import (
	"fmt"
	"log"
	"net/url"

	"github.com/kkohtaka/go-bitflyer/pkg/api"
	httpclient "github.com/kkohtaka/go-bitflyer/pkg/httpclient"
)

type API struct {
	url string
}

const (
	APIPath string = "board"
)

func NewAPI(c api.Client) *API {
	return &API{
		url: fmt.Sprintf("%s/%s", c.APIHost(), APIPath),
	}
}

func (api *API) BaseURL() url.URL {
	u, err := url.ParseRequestURI(api.url)
	if err != nil {
		log.Fatalf("Failed parsing a request URI: %+v", err)
	}
	return *u
}

func (api *API) Execute(req *Request) (*Response, error) {
	var resp Response
	err := httpclient.Get(api, req, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
