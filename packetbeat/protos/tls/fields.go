// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package tls

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("packetbeat", "tls", asset.ModuleFieldsPri, AssetTls); err != nil {
		panic(err)
	}
}

// AssetTls returns asset data.
// This is the base64 encoded zlib format compressed contents of protos/tls.
func AssetTls() string {
	return "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"
}
