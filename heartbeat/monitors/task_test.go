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

package monitors

import (
	"context"
	"testing"

	"github.com/elastic/go-lookslike/validator"

	"github.com/stretchr/testify/require"

	"github.com/elastic/go-lookslike"
	"github.com/elastic/go-lookslike/testslike"

	"github.com/sheng855174/elastic/heartbeat/eventext"
	"github.com/sheng855174/elastic/heartbeat/monitors/jobs"
	"github.com/sheng855174/elastic/libbeat/beat"
	"github.com/sheng855174/elastic/libbeat/common"
)

func Test_runPublishJob(t *testing.T) {
	defineJob := func(fields common.MapStr) func(event *beat.Event) (j []jobs.Job, e error) {
		return func(event *beat.Event) (j []jobs.Job, e error) {
			eventext.MergeEventFields(event, fields)
			return nil, nil
		}
	}
	simpleJob := defineJob(common.MapStr{"foo": "bar"})

	testCases := []struct {
		name       string
		job        jobs.Job
		validators []validator.Validator
	}{
		{
			"simple",
			simpleJob,
			[]validator.Validator{
				lookslike.MustCompile(map[string]interface{}{"foo": "bar"}),
			},
		},
		{
			"one cont",
			func(event *beat.Event) (j []jobs.Job, e error) {
				simpleJob(event)
				return []jobs.Job{simpleJob}, nil
			},
			[]validator.Validator{
				lookslike.MustCompile(map[string]interface{}{"foo": "bar"}),
				lookslike.MustCompile(map[string]interface{}{"foo": "bar"}),
			},
		},
		{
			"multiple conts",
			func(event *beat.Event) (j []jobs.Job, e error) {
				simpleJob(event)
				return []jobs.Job{
					defineJob(common.MapStr{"baz": "bot"}),
					defineJob(common.MapStr{"blah": "blargh"}),
				}, nil
			},
			[]validator.Validator{
				lookslike.MustCompile(map[string]interface{}{"foo": "bar"}),
				lookslike.MustCompile(map[string]interface{}{"baz": "bot"}),
				lookslike.MustCompile(map[string]interface{}{"blah": "blargh"}),
			},
		},
		{
			"cancelled cont",
			func(event *beat.Event) (j []jobs.Job, e error) {
				eventext.CancelEvent(event)
				return []jobs.Job{simpleJob}, nil
			},
			[]validator.Validator{
				lookslike.MustCompile(map[string]interface{}{"foo": "bar"}),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := &MockBeatClient{}
			queue := runPublishJob(tc.job, &WrappedClient{
				Publish: client.Publish,
				Close:   client.Close,
				wait:    func() {},
			})
			for {
				if len(queue) == 0 {
					break
				}
				tf := queue[0]
				queue = queue[1:]
				conts := tf(context.Background())
				for _, cont := range conts {
					queue = append(queue, cont)
				}
			}
			client.Close()

			require.Len(t, client.publishes, len(tc.validators))
			for idx, event := range client.publishes {
				testslike.Test(t, tc.validators[idx], event.Fields)
			}
		})
	}
}
