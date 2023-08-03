// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by dev-tools/cmd/buildfleetcfg/buildfleetcfg.go - DO NOT EDIT.

package application

import "github.com/elastic/beats/v7/x-pack/elastic-agent/pkg/packer"

// DefaultAgentFleetConfig is the content of the default configuration when we enroll a beat, the elastic-agent.yml
// will be replaced with this variables.
var DefaultAgentFleetConfig []byte

func init() {
	// Packed File
	// _meta/elastic-agent.fleet.yml
	unpacked := packer.MustUnpack("eJycVc2S4roV3ucxZp8b/+Cp61RlgdwjWUzjDuYiWdqkLImxDZJxBWyQU3n3lGwaum9Sqam7YCOkc873cz7/68s/zP5S/mWvy/OlkX8uq317+eWH3u8vv1ijv/z1C7bL+O+/eT/324KU00hLA8/45SffPH/fsQUJp9CTFnTSgoMIIlNS5UtDjjglPUfE4m/8zCnxXhMQiXBXsQD2GOWaG93zLbiIMPNwmmmV5p0wakyadcXb1SC2wCvRrtoGumE0qjEkq9cG8EeNlFw5gh6f7ukRI94JtKsU+rViQdxzo1tVrDRO8xPfgpEXuVfSqJXW3Yee68Pc/ygaVAJGVYCrCFcep5uKtccqNzfNi02fmHn+t+pUYUR6RldnTjdfcZpbRXffpzpUadFueo7iUJjbwOjmqzvHyfKA02xQxerAt2DgzYyJFSuvpLxmYT6+JsDyAvplsdLSAi0QHBXSB4xuXRnsKhm6Ppk3YUGw5wnoGfU7YWQlAlYpVGuMVp0w0JaUjI++aNdzuxxUoKe5pPUXby+s3xt9ne+AUYTEsoCQ3fH4FSerk0rzqxxPw2sArULaMJp50kaPmV5NNrwG+aCC6CwCeJQ2bmb9f22eeEGtUDXpMM+Wd9KQg0Kx3W8nj3iM6t5pgdGzluNdIngoA9jyYv0d22WF07yWRmkF46Pjh4X5IFs3Kz7uw6lP7zzACm94cp+PuDodcLKssNOa+lqExMOID7IBRxGqXqC45kj33IIro9k/nX53Th7335rlbX0AH/jUvQzzWqDbD4mgVybgyIvsII1ueLG+Y70NLIBeSeMeo9h8eKN58uwlbXyVJm6lgRdpF9W22FRl4Xx0G4UFvgzI8dl3mttyCo8TZymwIlCWUa+SIdAs0KakWa2QHkS77u84OtFmHqO3czHz9D85K2k0KgTPIsGffCzT1cACMsrxNOuQrCsxaebXIpn3BKPIF3SlZQOMCLGbpVZBNOGUJj7wIhsfOiK/Ezqe3k28viyv6+VDowsr6k4U5IJTMMh2U4mWXJgh9o59lAiGYsoUhz3TMsy0vM7vBYWLIpzefcUv66ur/fRi3glKBlVsKm5ii1HmePRm70Va2U+z9jhVJ04X1QP/nGu9CDeVDMlhOgu4L9Ateq/1nhcK1Z384LcP+++5fccovvJiNevUPPJyLGmup7xE3ypuYCdSYvlmxnbn9IdC+sJp7KuXU7UewZg4fyPYchp5r+1Kq3QVvT09r4WBjUCuZq4VItZpN3k0qLXz6eSnInf76Oa4ijAbS5dr20WVI2JYQc4qmf4zzpN8u7jr+N/59+4PaYgnjdbSjw8iJL3LVInIoaS8c72fmsYGp9PM3cT5N/ctgL5I5x1Y28Wj5s7VbI8uz33hfBGwar0FF05h77RRCQhLejtj5HL+XpO6e5Getd9VjHKPF3g+Q7wu6c2XZve+JxMPLjeLIBsUjZw/P+x87rIvwml25TTruMNnwZHTW+24Ew24SAsaXuQhp6T/v7jG9yy5v//t9OHM9fLrPYzfa3/mjsJFSX1fbP9o7+X1074EM663ZumLzcM3cJ+CQRjSK6Tdnh954bzpvjVknHajyE+OS47ImRXHqbZKV5oVudv72n2/Ptd4akkC7ZXJclyn68d7RqMjfvngndFxsjzjF3Z9TX6Xv7/DxA08y2D37suFRHHv9qGk8PzW3H27Of3ty7//9J8AAAD//zmvIS8=")
	raw, ok := unpacked["_meta/elastic-agent.fleet.yml"]
	if !ok {
		// ensure we have something loaded.
		panic("elastic-agent.fleet.yml is not included in the binary")
	}
	DefaultAgentFleetConfig = raw
}
