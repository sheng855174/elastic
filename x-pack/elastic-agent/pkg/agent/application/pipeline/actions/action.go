// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package actions

import (
	"context"

	"github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/agent/storage/store"
	"github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/fleetapi"
	"github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/fleetapi/client"
)

// Handler handles action coming from fleet.
type Handler interface {
	Handle(ctx context.Context, a fleetapi.Action, acker store.FleetAcker) error
}

// ClientSetter sets the client for communication.
type ClientSetter interface {
	SetClient(client.Sender)
}
