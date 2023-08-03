// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"os"

	"github.com/sheng855174/elastic/libbeat/cmd"
	"github.com/sheng855174/elastic/libbeat/mock"
	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
)

// RootCmd to test libbeat
var RootCmd = cmd.GenRootCmdWithSettings(mock.New, mock.Settings)

func main() {
	if err := RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
