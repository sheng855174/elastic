// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"os"

	"github.com/sheng855174/elastic/x-pack/functionbeat/provider/local/cmd"
	_ "github.com/sheng855174/elastic/x-pack/functionbeat/provider/local/include" // imports features
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
