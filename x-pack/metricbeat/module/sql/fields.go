// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package sql

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "sql", asset.ModuleFieldsPri, AssetSql); err != nil {
		panic(err)
	}
}

// AssetSql returns asset data.
// This is the base64 encoded zlib format compressed contents of module/sql.
func AssetSql() string {
	return "eJy0k0Fu8jAQhfc5xRPLX4IDZPEvqi4REuIAlWM/wMXJgD2mze0rEoJSpa1AVb18npnvy0Se48C2RDqFAlCvgSVmm/VyVgCRgSaxREU1BeCYbPRH9dKU+F8AwGa9RC0uB2JLtXsm1NTobcI2Sg3TVTijpjKJBbD1DC6VXfMcjak5wC9H2yNL7KLk4zUZ1497XPRnxls8tB7Yvkl0o/wL6eE8dzOQEx1UwHfarITuiVNmbBcTahf/Drq+jBhYHddKCLQ6LG5KHS6aXDN6u/g3MZDqlVZHcR+89LdOchV4n96qZ9z+4lWO7nutpNE3u4etHtraSpr59fNxNiHzLrNKJNA0f6v21EN+1Pr8kj4CAAD//07X+ko="
}
