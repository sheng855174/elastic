// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	"github.com/sheng855174/elastic/x-pack/osquerybeat/beater"
	"github.com/sheng855174/elastic/x-pack/osquerybeat/internal/install"

	cmd "github.com/sheng855174/elastic/libbeat/cmd"
	"github.com/sheng855174/elastic/libbeat/cmd/instance"
	"github.com/sheng855174/elastic/libbeat/common"
	"github.com/sheng855174/elastic/libbeat/common/cli"
	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/libbeat/publisher/processing"

	"github.com/spf13/cobra"

	_ "github.com/sheng855174/elastic/x-pack/libbeat/include"
)

// Name of this beat
const (
	Name = "osquerybeat"

	// ecsVersion specifies the version of ECS that this beat is implementing.
	ecsVersion = "1.12.0"
)

// withECSVersion is a modifier that adds ecs.version to events.
var withECSVersion = processing.WithFields(common.MapStr{
	"ecs": common.MapStr{
		"version": ecsVersion,
	},
})

var RootCmd = Osquerybeat()

func Osquerybeat() *cmd.BeatsRootCmd {
	settings := instance.Settings{
		Name:            Name,
		Processing:      processing.MakeDefaultSupport(true, withECSVersion, processing.WithAgentMeta()),
		ElasticLicensed: true,
	}
	command := cmd.GenRootCmdWithSettings(beater.New, settings)

	// Add verify command
	command.AddCommand(genVerifyCmd(settings))

	return command
}

func genVerifyCmd(settings instance.Settings) *cobra.Command {
	return &cobra.Command{
		Use:   "verify",
		Short: "Verify installation",
		Run: cli.RunWith(
			func(_ *cobra.Command, args []string) error {
				log := logp.NewLogger("osquerybeat")
				err := install.VerifyWithExecutableDirectory(log)
				if err != nil {
					return err
				}
				return nil
			}),
	}
}
