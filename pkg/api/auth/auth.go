// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"time"

	"github.com/kkohtaka/go-bitflyer/pkg/api"
	"github.com/pkg/errors"
)

type AuthConfig struct {
	APIKey    string
	APISecret string
}

func GenerateAuthHeaders(config *AuthConfig, now time.Time, api api.API, req api.Request) (*http.Header, error) {
	url, err := api.BaseURL()
	if err != nil {
		return nil, errors.Wrapf(err, "set base URI")
	}
	url.RawQuery = req.Query()

	timestamp := fmt.Sprintf("%d", now.Unix())
	method := req.Method()
	path := url.Path
	payload := req.Payload()

	mac := hmac.New(sha256.New, []byte(config.APISecret))
	mac.Write([]byte(timestamp))
	mac.Write([]byte(method))
	mac.Write([]byte(path))
	if len(payload) > 0 {
		mac.Write(payload)
	}
	sign := hex.EncodeToString(mac.Sum(nil))

	header := http.Header{}
	header.Set("ACCESS-KEY", config.APIKey)
	header.Set("ACCESS-TIMESTAMP", timestamp)
	header.Set("ACCESS-SIGN", sign)

	return &header, nil
}
