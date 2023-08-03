// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	// include the composable providers
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/agent"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/docker"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/env"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/host"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/kubernetes"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/kubernetesleaderelection"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/kubernetessecrets"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/local"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/localdynamic"
	_ "github.com/sheng855174/elastic/x-pack/elastic-agent/pkg/composable/providers/path"
)
