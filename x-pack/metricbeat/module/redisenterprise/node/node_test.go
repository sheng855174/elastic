// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !integration
// +build !integration

package node

import (
	"os"
	"testing"

	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/metricbeat/mb"
	mbtest "github.com/sheng855174/elastic/metricbeat/mb/testing"

	// Register input module and metricset
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus"
	_ "github.com/sheng855174/elastic/metricbeat/module/prometheus/collector"
)

func init() {
	// To be moved to some kind of helper
	os.Setenv("BEAT_STRICT_PERMS", "false")
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}

func TestEventMapping(t *testing.T) {
	logp.TestingSetup()

	mbtest.TestDataFiles(t, "redisenterprise", "node")
}
