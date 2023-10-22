// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Khulnasoft

package flow

import (
	flowpb "github.com/khulnasoft/shipyard/api/v1/flow"
	"github.com/khulnasoft/faker"
)

type serviceOptions struct {
	namespace string
	name      string
}

// ServiceOption is an option to use with Service.
type ServiceOption interface {
	apply(o *serviceOptions)
}

type funcServiceOption func(*serviceOptions)

func (fso funcServiceOption) apply(so *serviceOptions) {
	fso(so)
}

// WithServiceNamespace sets the namespace of the service.
func WithServiceNamespace(ns string) ServiceOption {
	return funcServiceOption(func(o *serviceOptions) {
		o.namespace = ns
	})
}

// WithServiceName sets the name of the service.
func WithServiceName(name string) ServiceOption {
	return funcServiceOption(func(o *serviceOptions) {
		o.name = name
	})
}

// Service generates a random Service. Options may be provided to customize the
// service to return.
func Service(options ...ServiceOption) *flowpb.Service {
	opts := serviceOptions{
		namespace: fake.K8sNamespace(),
		name:      fake.Name(),
	}
	for _, opt := range options {
		opt.apply(&opts)
	}

	return &flowpb.Service{
		Namespace: opts.namespace,
		Name:      opts.name,
	}
}
