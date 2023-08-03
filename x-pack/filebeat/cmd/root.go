// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	fbcmd "github.com/sheng855174/elastic/filebeat/cmd"
	cmd "github.com/sheng855174/elastic/libbeat/cmd"

	// Register the includes.
	_ "github.com/sheng855174/elastic/x-pack/filebeat/include"
	inputs "github.com/sheng855174/elastic/x-pack/filebeat/input/default-inputs"
	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
)

const Name = fbcmd.Name

// Filebeat build the beat root command for executing filebeat and it's subcommands.
func Filebeat() *cmd.BeatsRootCmd {
	settings := fbcmd.FilebeatSettings()
	settings.ElasticLicensed = true
	command := fbcmd.Filebeat(inputs.Init, settings)
	return command
}
