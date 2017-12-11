// Copyright (C) 2017 Kazumasa Kohtaka <kkohtaka@gmail.com> All right reserved
// This file is available under the MIT license.

package types

type CurrencyCode string

type Pagination struct {
	Count  int `json:"count,omitempty" url:"count,omitempty"`
	Before int `json:"before,omitempty" url:"before,omitempty"`
	After  int `json:"after,omitempty" url:"after,omitempty"`
}
