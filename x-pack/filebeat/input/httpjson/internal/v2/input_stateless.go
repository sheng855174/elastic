// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package v2

import (
	v2 "github.com/sheng855174/elastic/filebeat/input/v2"
	stateless "github.com/sheng855174/elastic/filebeat/input/v2/input-stateless"
	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/sheng855174/elastic/libbeat/common"
)

type statelessInput struct {
	config config
}

func (statelessInput) Name() string {
	return "httpjson-stateless"
}

func statelessConfigure(cfg *common.Config) (stateless.Input, error) {
	conf := defaultConfig()
	if err := cfg.Unpack(&conf); err != nil {
		return nil, err
	}
	return newStatelessInput(conf)
}

func newStatelessInput(config config) (*statelessInput, error) {
	return &statelessInput{config: config}, nil
}

func (in *statelessInput) Test(v2.TestContext) error {
	return test(in.config.Request.URL.URL)
}

type statelessPublisher struct {
	wrapped stateless.Publisher
}

func (pub statelessPublisher) Publish(event beat.Event, _ interface{}) error {
	pub.wrapped.Publish(event)
	return nil
}

// Run starts the input and blocks until it ends the execution.
// It will return on context cancellation, any other error will be retried.
func (in *statelessInput) Run(ctx v2.Context, publisher stateless.Publisher) error {
	pub := statelessPublisher{wrapped: publisher}
	return run(ctx, in.config, pub, nil)
}
