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

package linux

import (
	"time"

	"github.com/sheng855174/elastic/libbeat/paths"
	"github.com/sheng855174/elastic/metricbeat/mb"
)

func init() {
	// Register the ModuleFactory function for the "system" module.
	if err := mb.Registry.AddModule("linux", NewModule); err != nil {
		panic(err)
	}
}

type LinuxModule interface {
	GetHostFS() string
}

// Module defines the base module config used in `linux`
type Module struct {
	mb.BaseModule
	HostFS string `config:"hostfs"`
	Period time.Duration
}

// NewModule initializes a new module
func NewModule(base mb.BaseModule) (mb.Module, error) {
	// This only needs to be configured once for all system modules.

	config := struct {
		Hostfs string        `config:"hostfs"`
		Period time.Duration `config:"period"`
	}{}

	if err := base.UnpackConfig(&config); err != nil {
		return nil, err
	}

	dir := config.Hostfs
	if dir == "" {
		dir = "/"
	}

	// Steer towards system.hostfs, since the two behave fundamentally the same, and system.hostfs has a CLI flag that many users may default to.
	if len(paths.Paths.Hostfs) > 2 {
		dir = paths.Paths.Hostfs
	}

	return &Module{BaseModule: base, HostFS: dir, Period: config.Period}, nil
}

func (m Module) GetHostFS() string {
	return m.HostFS
}
