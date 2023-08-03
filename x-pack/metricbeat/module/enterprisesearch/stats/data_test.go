// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !integration
// +build !integration

package stats

import (
	"testing"

	"github.com/sheng855174/elastic/libbeat/logp"
	mbtest "github.com/sheng855174/elastic/metricbeat/mb/testing"

	_ "github.com/sheng855174/elastic/x-pack/metricbeat/module/enterprisesearch"
)

func TestEventMapping(t *testing.T) {
	logp.TestingSetup()
	mbtest.TestDataFiles(t, "enterprisesearch", "stats")
}
