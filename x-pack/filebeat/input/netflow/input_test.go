// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

//go:build !integration
// +build !integration

package netflow

import (
	"testing"

	"github.com/sheng855174/elastic/filebeat/input/inputtest"
	"github.com/sheng855174/elastic/libbeat/common"
)

func TestNewInputDone(t *testing.T) {
	config := common.MapStr{}
	inputtest.AssertNotStartedInputCanBeDone(t, NewInput, &config)
}
