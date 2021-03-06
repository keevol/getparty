// getparty
// Copyright (C) 2016-2017 Vladimir Bauer
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/vbauerster/getparty"
)

var (
	version = "dev"
	commit  = "xxxxxxx"
)

func main() {
	cmd := &getparty.Cmd{Out: os.Stdout, Err: os.Stderr}
	os.Exit(cmd.Exit(cmd.Run(
		os.Args[1:],
		fmt.Sprintf("%s (%.7s) (%s)", version, commit, runtime.Version()),
	)))
}
