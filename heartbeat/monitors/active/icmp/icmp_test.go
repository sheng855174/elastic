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

package icmp

import (
	"net"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/sheng855174/elastic/heartbeat/hbtest"
	"github.com/sheng855174/elastic/heartbeat/look"
	"github.com/sheng855174/elastic/heartbeat/monitors"
	"github.com/sheng855174/elastic/heartbeat/monitors/stdfields"
	"github.com/sheng855174/elastic/heartbeat/monitors/wrappers"
	"github.com/sheng855174/elastic/heartbeat/scheduler/schedule"
	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/elastic/go-lookslike"
	"github.com/elastic/go-lookslike/testslike"
)

func TestICMPFields(t *testing.T) {
	host := "localhost"
	hostURL := &url.URL{Scheme: "icmp", Host: host}
	ip := "127.0.0.1"
	cfg := Config{
		Hosts: []string{host},
		Mode:  monitors.IPSettings{IPv4: true, IPv6: false, Mode: monitors.PingAny},
	}
	testMockLoop, e := execTestICMPCheck(t, cfg)

	validator := lookslike.Strict(
		lookslike.Compose(
			hbtest.BaseChecks(ip, "up", "icmp"),
			hbtest.SummaryChecks(1, 0),
			hbtest.URLChecks(t, hostURL),
			hbtest.ResolveChecks(ip),
			lookslike.MustCompile(map[string]interface{}{
				"icmp.requests": 1,
				"icmp.rtt":      look.RTT(testMockLoop.pingRtt),
			}),
		),
	)
	testslike.Test(t, validator, e.Fields)
}

func execTestICMPCheck(t *testing.T, cfg Config) (mockLoop, *beat.Event) {
	tl := mockLoop{pingRtt: time.Microsecond * 1000, pingRequests: 1}
	jf, err := newJobFactory(cfg, monitors.NewStdResolver(), tl)
	require.NoError(t, err)
	p, err := jf.makePlugin()
	require.Len(t, p.Jobs, 1)
	require.Equal(t, 1, p.Endpoints)
	e := &beat.Event{}
	sched, _ := schedule.Parse("@every 1s")
	wrapped := wrappers.WrapCommon(p.Jobs, stdfields.StdMonitorFields{ID: "test", Type: "icmp", Schedule: sched, Timeout: 1})
	wrapped[0](e)
	return tl, e
}

type mockLoop struct {
	pingRtt             time.Duration
	pingRequests        int
	pingErr             error
	checkNetworkModeErr error
}

func (t mockLoop) checkNetworkMode(mode string) error {
	return t.checkNetworkModeErr
}

func (t mockLoop) ping(addr *net.IPAddr, timeout time.Duration, interval time.Duration) (time.Duration, int, error) {
	return t.pingRtt, t.pingRequests, t.pingErr
}
