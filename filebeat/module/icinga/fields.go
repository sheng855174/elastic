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

package icinga

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "icinga", asset.ModuleFieldsPri, AssetIcinga); err != nil {
		panic(err)
	}
}

// AssetIcinga returns asset data.
// This is the base64 encoded zlib format compressed contents of module/icinga.
func AssetIcinga() string {
	return "eJzsksFqwzAQRO/+iiH35AN06KWnHnrqF2yttbJEloS0TvDfFztO6hinUGihhexxBs08wWxx4N5AagmOKkBFPRtsXkZhUwGWS50lqcRg8FQBwNnEa7Sd5wpohL0tZvS2CNTyLHE47RMbuBy7NCkrqbc58yzL7527qmtxdyPP9xyDkoQyNaCJGbrny0fGfPjoym72bEkzJ2qoFi/a35gXsAP3p5jtwvsCb7i3xLU0wgWnPSnq2KYYOChic8H00Tm2I3jLpZDjT9xVysJHzvcoyQuVhZNI92bo2Xk+sl+4rbhMZ37NHa9WTmDfbFx/da/v2kUSfnEVQ/xjFP9tFEUp680CfnoXU8NjGn99Gh8BAAD//xYdyXI="
}
