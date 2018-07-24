// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package httpclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/kkohtaka/go-bitflyer/pkg/api"
	"github.com/kkohtaka/go-bitflyer/pkg/api/auth"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

type httpClient struct {
	authConfig *auth.AuthConfig
}

func New() *httpClient {
	return &httpClient{}
}

func (hc *httpClient) Auth(authConfig *auth.AuthConfig) *httpClient {
	hc.authConfig = authConfig
	return hc
}

func (hc *httpClient) Request(api api.API, req api.Request, result interface{}) error {
	u, err := api.BaseURL()
	if err != nil {
		return errors.Wrapf(err, "set base URI")
	}
	payload := req.Payload()

	var body io.Reader
	if len(payload) > 0 {
		body = bytes.NewReader(payload)
	}
	rawReq, err := http.NewRequest(req.Method(), u.String(), body)
	if err != nil {
		return errors.Wrapf(err, "create POST request from url: %s", u.String())
	}
	if hc.authConfig != nil {
		header, err := auth.GenerateAuthHeaders(hc.authConfig, time.Now(), api, req)
		if err != nil {
			return errors.Wrap(err, "generate auth header")
		}
		rawReq.Header = *header
	}
	if len(payload) > 0 {
		rawReq.Header.Set("Content-Type", "application/json")
	}

	c := &http.Client{}
	resp, err := c.Do(rawReq)
	if err != nil {
		return errors.Wrapf(err, "send HTTP request with url: %s", u.String())
	}
	defer resp.Body.Close()

	// TODO: Don't use ioutil.ReadAll()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrapf(err, "read data fetched from url: %s", u.String())
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		return errors.Wrapf(err, "unmarshal data: %s", string(data))
	}
	return nil
}
