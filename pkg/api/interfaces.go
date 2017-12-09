// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package api

import "net/url"

type Client interface {
	APIHost() string
}

type API interface {
	BaseURL() url.URL
}

type Request interface {
	Method() string
	Query() string
	Payload() []byte
}
