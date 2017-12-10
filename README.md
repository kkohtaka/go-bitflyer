# go-bitflyer

[![Build Status](https://travis-ci.org/kkohtaka/go-bitflyer.svg?branch=master)](https://travis-ci.org/kkohtaka/go-bitflyer)
[![Coverage Status](https://coveralls.io/repos/github/kkohtaka/go-bitflyer/badge.svg?branch=master)](https://coveralls.io/github/kkohtaka/go-bitflyer?branch=master)
[![GoDoc](https://godoc.org/github.com/kkohtaka/go-bitflyer?status.svg)](https://godoc.org/github.com/kkohtaka/go-bitflyer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

go-bitflyer is a Go bindings for [bitFlyer Lightning API](https://lightning.bitflyer.jp/docs?lang=en).

## Usage

```golang
package main

import (
  "log"

  "github.com/kkohtaka/go-bitflyer/pkg/api/auth"
  "github.com/kkohtaka/go-bitflyer/pkg/api/v1"
  "github.com/kkohtaka/go-bitflyer/pkg/api/v1/markets"
  "github.com/kkohtaka/go-bitflyer/pkg/api/v1/permissions"
)

func main() {
  client := v1.NewClient(&v1.ClientOpts{
    AuthConfig: &auth.AuthConfig{
      APIKey:    "**********************",
      APISecret: "********************************************",
    },
  })

  if resp, err := client.Permissions(&permissions.Request{}); err != nil {
    log.Fatalln(err)
  } else {
    log.Println(resp)
  }

  if resp, err := client.Markets(&markets.Request{}); err != nil {
    log.Fatalln(err)
  } else {
    log.Println(resp)
  }
}

```
