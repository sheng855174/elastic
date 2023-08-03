// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package main

import (
	"fmt"
	"os"

	"github.com/sheng855174/elastic/x-pack/osquerybeat/cmd"

	"github.com/sheng855174/elastic/libbeat/common/proc"

	_ "github.com/sheng855174/elastic/x-pack/osquerybeat/include"
)

func main() {
	pj, err := proc.CreateJobObject()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create process JobObject: %v\n", err)
		os.Exit(1)
	}
	defer pj.Close()
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
