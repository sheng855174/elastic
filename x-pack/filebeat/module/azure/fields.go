// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package azure

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "azure", asset.ModuleFieldsPri, AssetAzure); err != nil {
		panic(err)
	}
}

// AssetAzure returns asset data.
// This is the base64 encoded zlib format compressed contents of module/azure.
func AssetAzure() string {
	return "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"
}
