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

/*
Package beater provides the implementation of the libbeat Beater interface for
Winlogbeat. The main event loop is implemented in this package.
*/
package beater

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/sheng855174/elastic/libbeat/common"
	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/libbeat/paths"

	"github.com/sheng855174/elastic/winlogbeat/checkpoint"
	"github.com/sheng855174/elastic/winlogbeat/config"
	"github.com/sheng855174/elastic/winlogbeat/eventlog"
)

// Time the application was started.
var startTime = time.Now().UTC()

// Winlogbeat is used to conform to the beat interface
type Winlogbeat struct {
	beat       *beat.Beat              // Common beat information.
	config     config.WinlogbeatConfig // Configuration settings.
	eventLogs  []*eventLogger          // List of all event logs being monitored.
	done       chan struct{}           // Channel to initiate shutdown of main event loop.
	pipeline   beat.Pipeline           // Interface to publish event.
	checkpoint *checkpoint.Checkpoint  // Persists event log state to disk.
	log        *logp.Logger
}

// New returns a new Winlogbeat.
func New(b *beat.Beat, _ *common.Config) (beat.Beater, error) {
	// Read configuration.
	config := config.DefaultSettings
	if err := b.BeatConfig.Unpack(&config); err != nil {
		return nil, fmt.Errorf("error reading configuration file: %w", err)
	}

	log := logp.NewLogger("winlogbeat")

	// resolve registry file path
	config.RegistryFile = paths.Resolve(paths.Data, config.RegistryFile)
	log.Infof("State will be read from and persisted to %s",
		config.RegistryFile)

	eb := &Winlogbeat{
		beat:   b,
		config: config,
		done:   make(chan struct{}),
		log:    log,
	}

	if err := eb.init(b); err != nil {
		return nil, err
	}

	return eb, nil
}

func (eb *Winlogbeat) init(b *beat.Beat) error {
	config := &eb.config

	// Create the event logs. This will validate the event log specific
	// configuration.
	eb.eventLogs = make([]*eventLogger, 0, len(config.EventLogs))
	for _, config := range config.EventLogs {
		eventLog, err := eventlog.New(config)
		if err != nil {
			return fmt.Errorf("failed to create new event log: %w", err)
		}
		eb.log.Debugf("initialized WinEventLog[%s]", eventLog.Name())

		logger, err := newEventLogger(b.Info, eventLog, config, eb.log)
		if err != nil {
			return fmt.Errorf("failed to create new event log: %w", err)
		}

		eb.eventLogs = append(eb.eventLogs, logger)
	}

	return nil
}

// Setup uses the loaded config and creates necessary markers and environment
// settings to allow the beat to be used.
func (eb *Winlogbeat) setup(b *beat.Beat) error {
	config := &eb.config

	var err error
	eb.checkpoint, err = checkpoint.NewCheckpoint(config.RegistryFile, config.RegistryFlush)
	if err != nil {
		return err
	}

	eb.pipeline = b.Publisher
	return nil
}

// Run is used within the beats interface to execute the Winlogbeat workers.
func (eb *Winlogbeat) Run(b *beat.Beat) error {
	eb.log.Info("sheng Run(b *beat.Beat) start...")

	if err := eb.setup(b); err != nil {
		return err
	}

	acker := newEventACKer(eb.checkpoint)
	eb.log.Info("sheng eb.checkpoint.States() start...")
	persistedState := eb.checkpoint.States()
	eb.log.Info("sheng eb.checkpoint.States() end...")

	// Initialize metrics.
	initMetrics("total")

	var wg sync.WaitGroup
	for _, log := range eb.eventLogs {
		state, _ := persistedState[log.source.Name()]

		// Start a goroutine for each event log.
		wg.Add(1)
		eb.processEventLog(&wg, log, state, acker)
	}

	eb.log.Info("sheng range eb.eventLogs wait...")
	wg.Wait()
	eb.log.Info("sheng range eb.eventLogs end...")

	defer eb.checkpoint.Shutdown()

	if eb.config.ShutdownTimeout > 0 {
		eb.log.Infof("Shutdown will wait max %v for the remaining %v events to publish.",
			eb.config.ShutdownTimeout, acker.Active())
		ctx, cancel := context.WithTimeout(context.Background(), eb.config.ShutdownTimeout)
		defer cancel()
		acker.Wait(ctx)
	}

	return nil
}

// Stop is used to tell the winlogbeat that it should cease executing.
func (eb *Winlogbeat) Stop() {
	eb.log.Info("Stopping Winlogbeat")
	if eb.done != nil {
		close(eb.done)
	}
}

func (eb *Winlogbeat) processEventLog(
	wg *sync.WaitGroup,
	logger *eventLogger,
	state checkpoint.EventLogState,
	acker *eventACKer,
) {
	defer wg.Done()
	logger.run(eb.done, eb.pipeline, state, acker)
}
