// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// skipping tests on windows 32 bit versions, not supported
//go:build !integration && !windows && !386
// +build !integration,!windows,!386

package proxy

import (
	"os"
	"testing"

	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/metricbeat/mb"
	mbtest "github.com/sheng855174/elastic/metricbeat/mb/testing"

	// Register input module and metricset
	_ "github.com/sheng855174/elastic/x-pack/metricbeat/module/prometheus"
	_ "github.com/sheng855174/elastic/x-pack/metricbeat/module/prometheus/collector"
)

func init() {
	// To be moved to some kind of helper
	os.Setenv("BEAT_STRICT_PERMS", "false")
	mb.Registry.SetSecondarySource(mb.NewLightModulesSource("../../../module"))
}

func TestEventMapping(t *testing.T) {
	logp.TestingSetup()

	mbtest.TestDataFiles(t, "istio", "proxy")
}
