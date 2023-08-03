// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"os"

	"github.com/spf13/cobra"

	cmd "github.com/sheng855174/elastic/libbeat/cmd"
	"github.com/sheng855174/elastic/libbeat/cmd/instance"
	"github.com/sheng855174/elastic/x-pack/functionbeat/config"
	"github.com/sheng855174/elastic/x-pack/functionbeat/manager/beater"
)

// Name of this beat
var Name = "functionbeat"

// RootCmd to handle functionbeat
var RootCmd *cmd.BeatsRootCmd

func init() {
	RootCmd = cmd.GenRootCmdWithSettings(beater.New, instance.Settings{
		Name:            Name,
		HasDashboards:   false,
		ConfigOverrides: config.Overrides,
		ElasticLicensed: true,
	})

	RootCmd.RemoveCommand(RootCmd.RunCmd)
	RootCmd.Run = func(_ *cobra.Command, _ []string) {
		RootCmd.Usage()
		os.Exit(1)
	}

	RootCmd.AddCommand(genDeployCmd())
	RootCmd.AddCommand(genUpdateCmd())
	RootCmd.AddCommand(genRemoveCmd())
	RootCmd.AddCommand(genPackageCmd())

	addBeatSpecificSubcommands()
}

func addBeatSpecificSubcommands() {
	RootCmd.ExportCmd.Short = "Export current config, index template or function"
	RootCmd.ExportCmd.AddCommand(genExportFunctionCmd())
}
