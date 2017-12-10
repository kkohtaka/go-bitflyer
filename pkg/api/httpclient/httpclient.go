// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/kkohtaka/go-bitflyer/pkg/api"
	"github.com/kkohtaka/go-bitflyer/pkg/api/auth"
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
	u := api.BaseURL()
	payload := req.Payload()

	var body io.Reader
	if len(payload) > 0 {
		body = bytes.NewReader(payload)
	}
	rawReq, err := http.NewRequest(req.Method(), u.String(), body)
	if err != nil {
		fmt.Printf("failed creating a POST request from url: %s", u.String())
		return err
	}
	if hc.authConfig != nil {
		rawReq.Header = *(auth.GenerateAuthHeaders(hc.authConfig, time.Now(), api, req))
	}
	if len(payload) > 0 {
		rawReq.Header.Set("Content-Type", "application/json")
	}

	c := &http.Client{}
	resp, err := c.Do(rawReq)
	if err != nil {
		fmt.Printf("failed getting resource from url: %s", u.String())
		return err
	}
	defer resp.Body.Close()

	// TODO: Don't use ioutil.ReadAll()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("failed reading data from url: %s", u.String())
		return err
	}

	err = json.Unmarshal(data, result)
	if err != nil {
		fmt.Printf("failed unmarshalling data: %s", string(data))
	}
	return err
}
