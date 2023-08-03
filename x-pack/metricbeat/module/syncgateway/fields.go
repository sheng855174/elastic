// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package syncgateway

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "syncgateway", asset.ModuleFieldsPri, AssetSyncgateway); err != nil {
		panic(err)
	}
}

// AssetSyncgateway returns asset data.
// This is the base64 encoded zlib format compressed contents of module/syncgateway.
func AssetSyncgateway() string {
	return "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"
}
