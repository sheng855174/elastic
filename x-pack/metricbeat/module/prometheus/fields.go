// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package prometheus

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded zlib format compressed contents of module/prometheus.
func AssetPrometheus() string {
	return "eJzEkktqAzEMhvdzih8vQ5MDeNEzFLosJSi24rjxC0tTmtuXZIYw9EGzKMQ7+/+QPgmvceSTRes1sx54lPVHI3ccAI2a2MI8XSPoqbFHZu3RiRkAz+J6bBprsXgcAOBZSQXiOp3Zfa8ZhEUNLr7VWHQzAJ0Tk7BFoAEQVo0liMWLEUnmAeag2szrAOwjJy/20mGNQpmXzpvV5p3SyJcYF02Luntjp/PTdNlOia/jLvH3ZJuptVjCjJmVmZkfxjyfxVSBxsDzZn6XdHUsyv1+mrPAn6Kd9I7LPHf3N7seomgNnfKNwl/5/3G+Vl36Ts7LX/4ZAAD//xMCFRw="
}
