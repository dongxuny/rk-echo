// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"github.com/rookie-ninja/rk-echo/boot"
	"github.com/rookie-ninja/rk-entry/entry"
)

func main() {
	// Bootstrap basic entries from boot config.
	rkentry.RegisterInternalEntriesFromConfig("example/boot/simple/boot.yaml")

	// Bootstrap echo entry from boot config
	res := rkecho.RegisterEchoEntriesWithConfig("example/boot/simple/boot.yaml")

	// Bootstrap echo entry
	res["greeter"].Bootstrap(context.Background())

	// Wait for shutdown signal
	rkentry.GlobalAppCtx.WaitForShutdownSig()

	// Interrupt echo entry
	res["greeter"].Interrupt(context.Background())
}
