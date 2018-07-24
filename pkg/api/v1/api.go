// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package v1

import (
	"fmt"
	"net/url"

	"github.com/pkg/errors"
)

type API struct {
	url string
}

func NewAPI(c *Client, apiPath string) *API {
	return &API{
		url: fmt.Sprintf("%s%s", c.APIHost(), apiPath),
	}
}

func (api *API) BaseURL() (*url.URL, error) {
	u, err := url.ParseRequestURI(api.url)
	if err != nil {
		return nil, errors.Wrapf(err, "parse URI: %s", api.url)
	}
	return u, nil
}
