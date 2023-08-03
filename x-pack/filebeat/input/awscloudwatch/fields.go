// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package awscloudwatch

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "awscloudwatch", asset.ModuleFieldsPri, AssetAwscloudwatch); err != nil {
		panic(err)
	}
}

// AssetAwscloudwatch returns asset data.
// This is the base64 encoded zlib format compressed contents of input/awscloudwatch.
func AssetAwscloudwatch() string {
	return "eJyskUtq7DAQRedexaXn7gVo8OARyAYS6KFRrLIlWlYZqRzh3QfJTdpOMkhCalifc49QiyutCjqntve8mKyltw0gTjwpnI6DUwMYSn10szgOCv8aAHh05E3CEHnC/8sTHsrBpRzA85jODTDUFVXXWwQ90RehpWSdSWGMvMy3jqFBL166ilAYtE/0Pvrk8k2fUnunvZfnsdsL3LWutGaOZtc/GDxbqgjwALFUQNtLIIxsXW8h1iXQKwXBC3kOd5tjfpJIevoLgY30MwMXRkoF2Ymb6DcW5a4qbElZpxuUDFz48Cvn5i0AAP//m93HSw=="
}
