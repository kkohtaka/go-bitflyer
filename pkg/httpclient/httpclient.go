// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/kkohtaka/go-bitflyer/pkg/api"
)

func Get(api api.API, req api.Request, result interface{}) error {
	u := api.BaseURL()
	u.RawQuery = req.Query()

	c := &http.Client{}
	resp, err := c.Get(u.String())
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
