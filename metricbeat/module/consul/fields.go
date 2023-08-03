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

package consul

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "consul", asset.ModuleFieldsPri, AssetConsul); err != nil {
		panic(err)
	}
}

// AssetConsul returns asset data.
// This is the base64 encoded zlib format compressed contents of module/consul.
func AssetConsul() string {
	return "eJzcVsGO2zgMvecriJw38wE5LLC7h93LzgCzc1sUA1pmbLWyaJDUtP77QnY8sRM7k0ELtKguCSTx8b1HmtAOPlG3B8dRU9gAmLdAe9j+1W9sNwAlqRPfmue4h983AADDITRcpkAbAKFAqLSHggw3AEpmPla6h/+3qmH7G2xrs3b7YQNw8BRK3fc4O4jY0CR7Xta1tIdKOLXHnWnINAwriva6uxSZ1zm5cS3oGtcfGRj+JRPvlAwOZK4mhWbYAR8PLA3mUDgIN4CjJT6qYXQEkmL0sQLUAW2Cfy5nJikZtz6wzU7XpK2hTRFrwmB1d3E+YhbMgTAunM8MenghwRCOcMAHsJogsMMASvJCAi4kNZJFXZKi+YZuVjXL/TgE50qiUTmW4Z1GaHcecZ3CNcwpbtEZLSGf0APHauXCTOd9agqS7G2Pmf801LB0wIWhj1QO3Zadf/jvblVpgyGwe3ac4nkfvUlqRugfwha4+EjOFHrQ7P+1Vmufj9e/Ke/DeUrg2KvOGQBjCV4BoaKYm3I0qRVSTULgY+kdGssdPNVeocEOiiRqR/tyLxkPv0Uy0JpTKEHIksR8gqBGWHaghkbwgiHRut0VCyfzcbELbhf9eJwYJ7hLpYGx/GE6+2r85J/Qn/2Hc+qbouv75jidW2FHqldKiVJgRc+OQyBnLO+WO2Pz9wAHr3Aro+tWnyTF72PTJTFjwwD0hVzKV5bTjDxaTHo+y2+z54LJU9+kNVpfp+qCV406pCv7C9iuoV7zb8rdJRFanIu3K7gl3zTnatXmKa9Ub1xvVLF3avjgfYSIkZUcx3I9/0ixr/+vYMr9SfTwrmyGMaDG7c5q2n1mCeV5p+WXXG+egvr8fjtODDUUo/Ju8zUAAP//ku7rqQ=="
}
