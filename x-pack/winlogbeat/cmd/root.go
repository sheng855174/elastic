// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/sheng855174/elastic/libbeat/cmd"
	winlogbeatCmd "github.com/sheng855174/elastic/winlogbeat/cmd"

	// Register fields.
	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
	_ "github.com/sheng855174/elastic/x-pack/winlogbeat/include"
)

// Name of this beat.
var Name = winlogbeatCmd.Name

// RootCmd to handle beats cli
var RootCmd *cmd.BeatsRootCmd

func init() {
	settings := winlogbeatCmd.WinlogbeatSettings()
	settings.ElasticLicensed = true
	RootCmd = winlogbeatCmd.Initialize(settings)
}
