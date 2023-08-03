// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package zookeeper

import (
	"github.com/sheng855174/elastic/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "zookeeper", asset.ModuleFieldsPri, AssetZookeeper); err != nil {
		panic(err)
	}
}

// AssetZookeeper returns asset data.
// This is the base64 encoded zlib format compressed contents of module/zookeeper.
func AssetZookeeper() string {
	return "eJysVE2L2zAUvOdXDDntwm727kMh5NgWlu6eeikv0iQWUSSjJyV4f32xG+erbkvB72TmSTPzxvZ7xo5thY8Yd2TDNAOyy54V5t9j/Nxj8xmQ6CnKCmtmmQGWapJrsouhwqcZAJzP42u0xXMGbBy91apvPyPInrdSXeW2YYVtiqU5ISPct1TXdFKsy2d0jO6PlL/qYnvZUcHHrS6uTtxPPtS9oWtTSlUXw01vsLZje4zJ3vX+YrCrlXcMeeCFs6OyHyFaTif6KrlG3CDXHKG+Ef3RyUyn/N427JR7brgAI9oDJlG6S4gNU/806kmMn87MW04ubJHYJCpDPhkY3C1XX+DdjjA2HeWhd8gnWHpmPiWKfcIxuQ4Tu3fhcYH32imcdp/alhYx+BabmKDMS+P/MVuiFp+nG+9bzze857P2Aq9R1a09cRBfqJBEPGgxhqovG3G+JL64cIg72sfFwDM/IfNuwKK0p8nSgektx+aigDWNFCW0gy9xrLmJiWDQ0gefa8knAojJRbxv+zsN7WI0oqLn7TJBQKu43wuUjSTJtPBO+7g6FcWxjn0yohqNk0wcXa4hMDe/7G97y8ftNFvrf/fVzwAAAP//jcqrBQ=="
}
