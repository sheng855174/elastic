// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration
// +build integration

package tablespace

import (
	"testing"

	_ "github.com/godror/godror"

	"github.com/sheng855174/elastic/libbeat/tests/compose"
	mbtest "github.com/sheng855174/elastic/metricbeat/mb/testing"
	"github.com/sheng855174/elastic/x-pack/metricbeat/module/oracle"
)

func TestData(t *testing.T) {
	t.Skip("Skip until a proper Docker image is setup for Metricbeat")
	r := compose.EnsureUp(t, "oracle")

	f := mbtest.NewReportingMetricSetV2WithContext(t, getConfig(r.Host()))

	if err := mbtest.WriteEventsReporterV2WithContext(f, t, ""); err != nil {
		t.Fatal("write", err)
	}
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "oracle",
		"metricsets": []string{"tablespace"},
		"hosts":      []string{oracle.GetOracleConnectionDetails(host)},
	}
}
