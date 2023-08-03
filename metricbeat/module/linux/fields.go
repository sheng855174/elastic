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

package linux

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "linux", asset.ModuleFieldsPri, AssetLinux); err != nil {
		panic(err)
	}
}

// AssetLinux returns asset data.
// This is the base64 encoded zlib format compressed contents of module/linux.
func AssetLinux() string {
	return "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"
}
