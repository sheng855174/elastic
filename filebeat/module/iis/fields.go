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

package iis

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "iis", asset.ModuleFieldsPri, AssetIis); err != nil {
		panic(err)
	}
}

// AssetIis returns asset data.
// This is the base64 encoded zlib format compressed contents of module/iis.
func AssetIis() string {
	return "eJzsmM9ug0YQxu9+ilHODYf25kMvlar6UKlSIvWI1jDgUfAOmd21xdtXu5jUcRbMH1tWpHCywPN9v4XZYYZneMNmDURmBWDJVriGp83m5WkFkKPJhGpLrNfw+woA4G/OXYVQsECtxJAuYbN5gYpLKKhCk6wACsIqN+vw/2fQao+dvj9sU+MaSmFXn85EbPzxZ5CBQnj/1cMfnc9HRGemsgyN+Tgd8xzw9ccfrK0ibU4eYb2eoVX2KB8U5yTnEh2McdvUWGWd+XS5g6pYlxcXBrj88bpD+Ov19R+v3ApDxjkmUfcj6d9+vb3/v6RzPhq46m/IYup/Rs3fsDmy5NP9vWxwAKVzIG2s0hmCdvstSg8JygHlDiwBgwuwniuYAGs47ijbhXNd3gJqKw0clYESNYqymMdRM+Y3ujFlxtqith1oawEmnBIQzJAOmP8CVIDSzf9cUcAt503axSTbxmI8uVRF6vJKrexuDTtr60Tw3aGxiZeLquypFNWuyorDfha/jiUcpmZtcBnIKcGonoiQo7Gkg3qi8lw+l66x7nu0O77Mi0lPIaowxtpJNdHXSZX4XzPM3h1Kkxor9KVujXEN4TNsaxa74LlGwkfdWTNUsQZWalCSSNwYT8E9W5yexoadZLgggwULFEFZksM9GuPs2yKQ+pfZoirSvhRjOmM4vFZ6QDHEeg5GPHSUMxs7I9l8WNITOzrJVYk6vsEuuzbo6bfOJXM8UHbJMrwW+LR5Wp6k1YntpGuLO6fpCZ/IshCCYyYTEdgkhauqWEWahtKnMJ1nKYpQSVpdvr/msPQrjdkEJXJPxZ2T/77ZI43aLrnRp4JeIidX9cbe74ydb4JTMhyrj7PQriqOhas4C39bDjWgNBZGsCTWN3p+w2KjHx7Z5lYJNSA18Q7dLpWuC/ahdUgownKP7w5BeNJnB0FlWKf1TpS58SgZPj60+tDqx8fYd4fuHsN/+BBT1xW1mwxq5ipYXplaH9jZBuMZM8TJeub48NiJ9OT+mMnpcV30txvEezqX7zQz/XRPs9F+uqerPD/d0326p/8CAAD//y6Wud4="
}
