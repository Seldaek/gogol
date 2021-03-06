// Copyright 2013 Marc Weistroff. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package gogol

import (
	"runtime"
	"testing"
)

func TestRuntimeProcessor(t *testing.T) {
	r := newRecord(DEBUG, "test", "foobar")
	RuntimeProcessor.Process(r)
	if r.Extra["go.num_cpu"] != runtime.NumCPU() {
		t.Error()
	}
	if r.Extra["go.version"] != runtime.Version() {
		t.Error()
	}
	if r.Extra["go.num_goroutines"] != runtime.NumGoroutine() {
		t.Error()
	}
}
