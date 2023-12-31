// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package agent

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable"
	ctesting "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/testing"
)

func TestContextProvider(t *testing.T) {
	builder, _ := composable.Providers.GetContextProvider("agent")
	provider, err := builder(nil, nil)
	require.NoError(t, err)

	comm := ctesting.NewContextComm(context.Background())
	err = provider.Run(comm)
	require.NoError(t, err)

	current := comm.Current()
	_, hasID := current["id"]
	assert.True(t, hasID, "missing id")
	_, hasVersion := current["version"]
	assert.True(t, hasVersion, "missing version")
}
