// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	heartbeatCmd "github.com/sheng855174/elastic/heartbeat/cmd"
	"github.com/sheng855174/elastic/libbeat/cmd"

	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
)

// RootCmd to handle beats cli
var RootCmd *cmd.BeatsRootCmd

func init() {
	settings := heartbeatCmd.HeartbeatSettings()
	settings.ElasticLicensed = true
	RootCmd = heartbeatCmd.Initialize(settings)
}
