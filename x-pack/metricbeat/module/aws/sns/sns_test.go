// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package sns

import (
	"os"

	"github.com/sheng855174/elastic/metricbeat/mb"

	// Register input module and metricset
	_ "github.com/sheng855174/elastic/x-pack/metricbeat/module/aws"
	_ "github.com/sheng855174/elastic/x-pack/metricbeat/module/aws/cloudwatch"
)

func init() {
	// To be moved to some kind of helper
	os.Setenv("BEAT_STRICT_PERMS", "false")
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}
