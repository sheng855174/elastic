// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package v2

import (
	v2 "github.com/sheng855174/elastic/filebeat/input/v2"
	inputcursor "github.com/sheng855174/elastic/filebeat/input/v2/input-cursor"
	"github.com/sheng855174/elastic/libbeat/common"
)

type cursorInput struct{}

func (cursorInput) Name() string {
	return "httpjson-cursor"
}

type source struct {
	config config
}

func (src source) Name() string {
	return src.config.Request.URL.String()
}

func cursorConfigure(cfg *common.Config) ([]inputcursor.Source, inputcursor.Input, error) {
	conf := defaultConfig()
	if err := cfg.Unpack(&conf); err != nil {
		return nil, nil, err
	}
	sources, inp := newCursorInput(conf)
	return sources, inp, nil
}

func newCursorInput(config config) ([]inputcursor.Source, inputcursor.Input) {
	// we only allow one url per config, if we wanted to allow more than one
	// each source should hold only one url
	return []inputcursor.Source{&source{config: config}}, &cursorInput{}
}

func (in *cursorInput) Test(src inputcursor.Source, _ v2.TestContext) error {
	return test((src.(*source)).config.Request.URL.URL)
}

// Run starts the input and blocks until it ends the execution.
// It will return on context cancellation, any other error will be retried.
func (in *cursorInput) Run(
	ctx v2.Context,
	src inputcursor.Source,
	cursor inputcursor.Cursor,
	publisher inputcursor.Publisher,
) error {
	s := src.(*source)
	return run(ctx, s.config, publisher, &cursor)
}
