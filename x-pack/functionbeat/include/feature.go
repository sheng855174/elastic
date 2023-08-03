// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package include

import (
	"github.com/sheng855174/elastic/libbeat/feature"
	"github.com/sheng855174/elastic/x-pack/functionbeat/manager/aws"
	"github.com/sheng855174/elastic/x-pack/functionbeat/provider/local/local"
)

// Bundle feature enabled.
var Bundle = feature.MustBundle(
	aws.Bundle,
	local.Bundle,
)

func init() {
	feature.MustRegisterBundle(Bundle)
}
