// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package inputs

import (
	"github.com/sheng855174/elastic/filebeat/beater"
	ossinputs "github.com/sheng855174/elastic/filebeat/input/default-inputs"
	v2 "github.com/sheng855174/elastic/filebeat/input/v2"
	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/awscloudwatch"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/awss3"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/cloudfoundry"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/http_endpoint"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/httpjson"
	"github.com/sheng855174/elastic/x-pack/filebeat/input/o365audit"
)

func Init(info beat.Info, log *logp.Logger, store beater.StateStore) []v2.Plugin {
	return append(
		xpackInputs(info, log, store),
		ossinputs.Init(info, log, store)...,
	)
}

func xpackInputs(info beat.Info, log *logp.Logger, store beater.StateStore) []v2.Plugin {
	return []v2.Plugin{
		cloudfoundry.Plugin(),
		http_endpoint.Plugin(),
		httpjson.Plugin(log, store),
		o365audit.Plugin(log, store),
		awss3.Plugin(store),
		awscloudwatch.Plugin(store),
	}
}
