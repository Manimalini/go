// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin freebsd linux netbsd openbsd

package net

import (
	"runtime"
	"testing"
)

func TestTCPLookup(t *testing.T) {
	if testing.Short() || !*testExternal {
		t.Skip("skipping test to avoid external network")
	}
	c, err := Dial("tcp", "8.8.8.8:53")
	defer c.Close()
	if err != nil {
		t.Fatalf("Dial failed: %v", err)
	}
	cfg := &dnsConfig{timeout: 10, attempts: 3}
	_, err = exchange(cfg, c, "com.", dnsTypeALL)
	if err != nil {
		t.Fatalf("exchange failed: %v", err)
	}
}