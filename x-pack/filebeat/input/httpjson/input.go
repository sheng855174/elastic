// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package httpjson

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"go.uber.org/zap"

	inputv2 "github.com/sheng855174/elastic/filebeat/input/v2"
	cursor "github.com/sheng855174/elastic/filebeat/input/v2/input-cursor"
	stateless "github.com/sheng855174/elastic/filebeat/input/v2/input-stateless"
	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/sheng855174/elastic/libbeat/common"
	"github.com/sheng855174/elastic/libbeat/common/transport/httpcommon"
	"github.com/sheng855174/elastic/libbeat/common/useragent"
	"github.com/sheng855174/elastic/libbeat/feature"
	"github.com/sheng855174/elastic/libbeat/logp"
	v2 "github.com/sheng855174/elastic/x-pack/filebeat/input/httpjson/internal/v2"
	"github.com/elastic/go-concert/ctxtool"
	"github.com/elastic/go-concert/timed"
)

const (
	inputName = "httpjson"
)

var (
	userAgent = useragent.UserAgent("Filebeat", false)

	// for testing
	timeNow = time.Now
)

// retryLogger provides a shim for a *logp.Logger to be used by
// go-retryablehttp as a retryablehttp.LeveledLogger.
type retryLogger struct {
	log *logp.Logger
}

func newRetryLogger() *retryLogger {
	return &retryLogger{
		log: logp.NewLogger("httpjson.retryablehttp", zap.AddCallerSkip(1)),
	}
}

func (log *retryLogger) Error(msg string, keysAndValues ...interface{}) {
	log.log.Errorw(msg, keysAndValues...)
}

func (log *retryLogger) Info(msg string, keysAndValues ...interface{}) {
	log.log.Infow(msg, keysAndValues...)
}

func (log *retryLogger) Debug(msg string, keysAndValues ...interface{}) {
	log.log.Debugw(msg, keysAndValues...)
}

func (log *retryLogger) Warn(msg string, keysAndValues ...interface{}) {
	log.log.Warnw(msg, keysAndValues...)
}

func Plugin(log *logp.Logger, store cursor.StateStore) inputv2.Plugin {
	sim := stateless.NewInputManager(statelessConfigure)
	return inputv2.Plugin{
		Name:       inputName,
		Stability:  feature.Stable,
		Deprecated: false,
		Manager: inputManager{
			v2inputManager: v2.NewInputManager(log, store),
			stateless:      &sim,
			cursor: &cursor.InputManager{
				Logger:     log,
				StateStore: store,
				Type:       inputName,
				Configure:  cursorConfigure,
			},
		},
	}
}

func test(url *url.URL) error {
	port := func() string {
		if url.Port() != "" {
			return url.Port()
		}
		switch url.Scheme {
		case "https":
			return "443"
		}
		return "80"
	}()

	_, err := net.DialTimeout("tcp", net.JoinHostPort(url.Hostname(), port), time.Second)
	if err != nil {
		return fmt.Errorf("url %q is unreachable", url)
	}

	return nil
}

func run(
	ctx inputv2.Context,
	config config,
	publisher cursor.Publisher,
	cursor *cursor.Cursor,
) error {
	log := ctx.Logger.With("input_url", config.URL)

	stdCtx := ctxtool.FromCanceller(ctx.Cancelation)

	httpClient, err := newHTTPClient(stdCtx, config)
	if err != nil {
		return err
	}

	dateCursor := newDateCursorFromConfig(config, log)

	rateLimiter := newRateLimiterFromConfig(config, log)

	pagination := newPaginationFromConfig(config)

	requester := newRequester(
		config,
		rateLimiter,
		dateCursor,
		pagination,
		httpClient,
		log,
	)

	requester.loadCursor(cursor, log)

	// TODO: disallow passing interval = 0 as a mean to run once.
	if config.Interval == 0 {
		return requester.processHTTPRequest(stdCtx, publisher)
	}

	err = timed.Periodic(stdCtx, config.Interval, func() error {
		log.Info("Process another repeated request.")
		if err := requester.processHTTPRequest(stdCtx, publisher); err != nil {
			log.Error(err)
		}
		return nil
	})

	log.Infof("Context done: %v", err)

	return nil
}

func newHTTPClient(ctx context.Context, config config) (*http.Client, error) {
	config.Transport.Timeout = config.HTTPClientTimeout

	httpClient, err := config.Transport.Client(
		httpcommon.WithAPMHTTPInstrumentation(),
		config.KeepAlive.Settings(),
		httpcommon.WithHeaderRoundTripper(map[string]string{"User-Agent": userAgent}),
	)
	if err != nil {
		return nil, err
	}

	// Make retryable HTTP client
	client := &retryablehttp.Client{
		HTTPClient:   httpClient,
		Logger:       newRetryLogger(),
		RetryWaitMin: config.RetryWaitMin,
		RetryWaitMax: config.RetryWaitMax,
		RetryMax:     config.RetryMax,
		CheckRetry:   retryablehttp.DefaultRetryPolicy,
		Backoff:      retryablehttp.DefaultBackoff,
	}

	if config.OAuth2.isEnabled() {
		return config.OAuth2.client(ctx, client.StandardClient())
	}

	return client.StandardClient(), nil
}

func makeEvent(body string) beat.Event {
	fields := common.MapStr{
		"event": common.MapStr{
			"created": time.Now().UTC(),
		},
		"message": body,
	}

	return beat.Event{
		Timestamp: time.Now().UTC(),
		Fields:    fields,
	}
}
