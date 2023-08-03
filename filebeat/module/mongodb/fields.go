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

package mongodb

import (
	"github.com/sheng855174/elastictictic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "mongodb", asset.ModuleFieldsPri, AssetMongodb); err != nil {
		panic(err)
	}
}

// AssetMongodb returns asset data.
// This is the base64 encoded zlib format compressed contents of module/mongodb.
func AssetMongodb() string {
	return "eJyUk9Gu0zAMhu/7FNa5P0eaBOeiF0iwo0lcFJ4hNG5mLbFD4o6Vp0dpJ2ghbJovf+v397tunuGEUwtB2In91gAoqccWnq7KUwNgMfeJopJwCx8aAIBO7OgRBkkQTcrEDrpiePsEXhwM5DG/NAADobe5nT3PwCbgmlVKp4gtuCRjvCoV3FKHeRgMScKaNoNKrWFroBe3mvIv8Ca01F5YDXG+Ev6foJbiT45eQhRG1s34m2SAw8h96RkPvVF0kuinKQLIAAFzNg43JryYEMsJ91+77uOXt01zWf6E0w9JthqRFS8PBdwvlrtxiEkNW09ZkR8KlfGMiXSqmIwnkzd6NHqcT/7i8Yx+0wvkkln20DRiBVXb4Dap5rjPIfvIJ/7Mig4TJIwJM7KW96ZHhJHp+4hAtmgDYSpXKI3yCLMaxfD37/b7Hu9ed+9fd7vKql7YNb8CAAD//8hOG3I="
}
