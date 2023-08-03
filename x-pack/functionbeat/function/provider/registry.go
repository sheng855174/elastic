// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package provider

import (
	"errors"
	"fmt"

	"github.com/sheng855174/elastic/libbeat/common"
	"github.com/sheng855174/elastic/libbeat/feature"
	"github.com/sheng855174/elastic/libbeat/logp"
	"github.com/sheng855174/elastic/x-pack/functionbeat/config"
	"github.com/sheng855174/elastic/x-pack/functionbeat/function/core"
)

// Errors generated by the registry when we are retrieving providers or functions from the main registry.
var (
	errInvalidProvider     = errors.New("invalid provider name")
	errInvalidFunctionName = errors.New("invalid function name")
	errInvalidType         = errors.New("incomptible type received for the feature")
)

// namespace is the namespace were providers will be registered in the global registry.
const namespace = "functionbeat.provider"

// Factory factory to create a concrete provider for a specific cloud service.
type Factory func(*logp.Logger, *Registry, *common.Config) (Provider, error)

// FunctionFactory factory to create a concrete function.
type FunctionFactory func(Provider, *common.Config) (Function, error)

// Registry is a wrapper around the global feature registry and will take care of returning the
// the right providers and will do the type assertion for the providers, we hide the fact that
// we are actually accessing a global registry.
type Registry struct {
	registry *feature.Registry
}

// NewRegistry return a new registry.
func NewRegistry(registry *feature.Registry) *Registry {
	return &Registry{registry: registry}
}

// Lookup search the registry for the specific provider, normalization is done inside the
// registry to deal with lower case and uppercase.
func (r *Registry) Lookup(name string) (Factory, error) {
	if len(name) == 0 {
		return nil, errInvalidProvider
	}

	f, err := r.registry.Lookup(namespace, name)
	if err != nil {
		return nil, err
	}

	p, ok := f.Factory().(Factory)
	if !ok {
		return nil, errInvalidType
	}

	return p, nil
}

// LookupFunction takes a provider and a function and return the corresponding type or an
// error if the function or the provider is not found.
func (r *Registry) LookupFunction(provider, function string) (FunctionFactory, error) {
	if len(provider) == 0 {
		return nil, errInvalidProvider
	}
	if len(function) == 0 {
		return nil, errInvalidFunctionName
	}

	if _, err := r.Lookup(provider); err != nil {
		return nil, err
	}

	ns := getNamespace(provider)

	f, err := r.registry.Lookup(ns, function)
	if err != nil {
		return nil, err
	}

	fn, ok := f.Factory().(FunctionFactory)
	if !ok {
		return nil, errInvalidType
	}

	return fn, nil
}

// AvailableProviders returns the names of registered providers.
func (r *Registry) AvailableProviders() ([]string, error) {
	providerFeatures, err := r.registry.LookupAll(namespace)
	if err != nil {
		return nil, err
	}

	var providers []string
	for _, f := range providerFeatures {
		providers = append(providers, f.Name())
	}

	return providers, nil
}

// CreateFunctions create runnable function based on the configurations received.
func CreateFunctions(
	registry *Registry,
	provider Provider,
	enabledFunctions []string,
	configs []*common.Config,
	clientFactory clientFactory,
) ([]core.Runner, error) {
	var runners []core.Runner

	for _, cfg := range configs {
		c := config.DefaultFunctionConfig
		err := cfg.Unpack(&c)
		if err != nil {
			return nil, err
		}

		if strInSlice(enabledFunctions, c.Name.String()) == -1 {
			continue
		}

		if !c.Enabled {
			return nil, fmt.Errorf("function '%s' not enabled for provider '%s'", c.Name, provider.Name())
		}

		f, err := registry.LookupFunction(provider.Name(), c.Type)
		if err != nil {
			return nil, err
		}

		fn, err := f(provider, cfg)
		if err != nil {
			return nil, err
		}

		runners = append(runners, &Runnable{config: cfg, makeClient: clientFactory, function: fn})
	}

	if len(runners) == 0 {
		return nil, fmt.Errorf("no function are enabled for selected provider: '%s'", provider.Name())
	}
	return runners, nil
}

func strInSlice(haystack []string, name string) int {
	for idx, s := range haystack {
		if s == name {
			return idx
		}
	}
	return -1
}

// FindFunctionByName returns a function instance identified by a unique name or an error if not found.
func FindFunctionByName(
	registry *Registry,
	provider Provider,
	configs []*common.Config,
	name string,
) (Function, error) {

	for _, cfg := range configs {
		c := config.FunctionConfig{}
		err := cfg.Unpack(&c)
		if err != nil {
			return nil, err
		}

		if c.Name.String() != name {
			continue
		}

		if !c.Enabled {
			return nil, fmt.Errorf("function '%s' not enabled for provider '%s'", name, provider.Name())
		}

		f, err := registry.LookupFunction(provider.Name(), c.Type)
		if err != nil {
			return nil, err
		}

		fn, err := f(provider, cfg)
		if err != nil {
			return nil, err
		}
		return fn, nil
	}

	return nil, fmt.Errorf("no function with name '%s' exists", name)
}

// EnabledFunctions returns the list of enabled functions.
func EnabledFunctions(registry *Registry, provider Provider, configs []*common.Config) ([]string, error) {
	var names []string
	for _, cfg := range configs {
		c := config.FunctionConfig{}
		err := cfg.Unpack(&c)
		if err != nil {
			return names, fmt.Errorf("error while finding enabled functions: %+v", err)
		}

		if c.Enabled {
			names = append(names, c.Name.String())
		}
	}

	return names, nil
}
