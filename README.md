# go-bitflyer

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
