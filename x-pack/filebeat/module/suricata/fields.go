// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package suricata

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "suricata", asset.ModuleFieldsPri, AssetSuricata); err != nil {
		panic(err)
	}
}

// AssetSuricata returns asset data.
// This is the base64 encoded zlib format compressed contents of module/suricata.
func AssetSuricata() string {
	return "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"
}
