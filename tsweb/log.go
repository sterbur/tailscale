// Copyright (c) 2020 Tailscale Inc & AUTHORS All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tsweb

import (
	"encoding/json"
	"strings"
	"time"
)

// Msg is a structured event log entry.
type Msg struct {
	Where    string        `json:"where"`
	When     time.Time     `json:"when"`
	Duration time.Duration `json:"duration,omitempty"`
	Domain   string        `json:"domain,omitempty"`
	Msg      string        `json:"msg,omitempty"`
	Err      error         `json:"err,omitempty"`
	HTTP     *MsgHTTP      `json:"http,omitempty"`
	Data     interface{}   `json:"data,omitempty"`
}

// MsgHTTP contains information about the processing of one HTTP
// request.
type MsgHTTP struct {
	Code       int    `json:"code"`
	Path       string `json:"path"`
	RemoteAddr string `json:"remote_addr"`
	UserAgent  string `json:"user_agent"`
	Referer    string `json:"referer"`
}

// String returns m as a JSON string.
func (m Msg) String() string {
	if m.When.IsZero() {
		m.When = time.Now()
	}
	var buf strings.Builder
	json.NewEncoder(&buf).Encode(m)
	return strings.TrimRight(buf.String(), "\n")
}
