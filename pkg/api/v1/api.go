// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package v1

import (
	"fmt"
	"log"
	"net/url"
)

type API struct {
	url string
}

func NewAPI(c *Client, apiPath string) *API {
	return &API{
		url: fmt.Sprintf("%s%s", c.APIHost(), apiPath),
	}
}

func (api *API) BaseURL() url.URL {
	u, err := url.ParseRequestURI(api.url)
	if err != nil {
		log.Fatalf("Failed parsing a request URI: %+v", err)
	}
	return *u
}
