// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build integration
// +build integration

package channels

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/sheng855174/elastic/libbeat/tests/compose"
	mbtest "github.com/sheng855174/elastic/metricbeat/mb/testing"
)

func TestData(t *testing.T) {
	t.Skip("tests failing: https://github.com/elastic/beats/issues/34844")
	service := compose.EnsureUp(t, "stan")

	m := mbtest.NewFetcher(t, getConfig(service.Host()))
	m.WriteEvents(t, "")
}

func TestFetch(t *testing.T) {
	t.Skip("tests failing: https://github.com/elastic/beats/issues/34844")
	service := compose.EnsureUp(t, "stan")

	m := mbtest.NewFetcher(t, getConfig(service.Host()))
	events, errs := m.FetchEvents()
	if len(errs) > 0 {
		t.Fatalf("Expected 0 error, had %d. %v\n", len(errs), errs)
	}
	assert.NotEmpty(t, events)
}

func getConfig(host string) map[string]interface{} {
	return map[string]interface{}{
		"module":     "stan",
		"metricsets": []string{"channels"},
		"hosts":      []string{host},
	}
}
