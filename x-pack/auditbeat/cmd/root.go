// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	auditbeatcmd "github.com/sheng855174/elastic/auditbeat/cmd"
	"github.com/sheng855174/elastic/libbeat/cmd"

	// Register Auditbeat x-pack modules.
	_ "github.com/sheng855174/elastic/x-pack/auditbeat/include"
	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
)

// Name of the beat
var Name = auditbeatcmd.Name

// RootCmd to handle beats CLI.
var RootCmd *cmd.BeatsRootCmd

func init() {
	settings := auditbeatcmd.AuditbeatSettings()
	settings.ElasticLicensed = true
	RootCmd = auditbeatcmd.Initialize(settings)
}
