// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package memory

import (
	"github.com/sheng855174/elastic/metricbeat/mb"
	"github.com/sheng855174/elastic/x-pack/metricbeat/module/syncgateway"
)

func eventMapping(r mb.ReporterV2, content *syncgateway.SgResponse) {
	delete(content.MemStats, "BySize")
	delete(content.MemStats, "PauseNs")
	delete(content.MemStats, "PauseEnd")

	r.Event(mb.Event{
		MetricSetFields: content.MemStats,
	})
}
