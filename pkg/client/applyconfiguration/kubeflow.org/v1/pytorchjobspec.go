// Copyright 2024 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	kubefloworgv1 "github.com/kubeflow/training-operator/pkg/apis/kubeflow.org/v1"
)

// PyTorchJobSpecApplyConfiguration represents a declarative configuration of the PyTorchJobSpec type for use
// with apply.
type PyTorchJobSpecApplyConfiguration struct {
	RunPolicy           *RunPolicyApplyConfiguration                             `json:"runPolicy,omitempty"`
	ElasticPolicy       *ElasticPolicyApplyConfiguration                         `json:"elasticPolicy,omitempty"`
	PyTorchReplicaSpecs map[kubefloworgv1.ReplicaType]*kubefloworgv1.ReplicaSpec `json:"pytorchReplicaSpecs,omitempty"`
	NprocPerNode        *string                                                  `json:"nprocPerNode,omitempty"`
}

// PyTorchJobSpecApplyConfiguration constructs a declarative configuration of the PyTorchJobSpec type for use with
// apply.
func PyTorchJobSpec() *PyTorchJobSpecApplyConfiguration {
	return &PyTorchJobSpecApplyConfiguration{}
}

// WithRunPolicy sets the RunPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RunPolicy field is set to the value of the last call.
func (b *PyTorchJobSpecApplyConfiguration) WithRunPolicy(value *RunPolicyApplyConfiguration) *PyTorchJobSpecApplyConfiguration {
	b.RunPolicy = value
	return b
}

// WithElasticPolicy sets the ElasticPolicy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ElasticPolicy field is set to the value of the last call.
func (b *PyTorchJobSpecApplyConfiguration) WithElasticPolicy(value *ElasticPolicyApplyConfiguration) *PyTorchJobSpecApplyConfiguration {
	b.ElasticPolicy = value
	return b
}

// WithPyTorchReplicaSpecs puts the entries into the PyTorchReplicaSpecs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the PyTorchReplicaSpecs field,
// overwriting an existing map entries in PyTorchReplicaSpecs field with the same key.
func (b *PyTorchJobSpecApplyConfiguration) WithPyTorchReplicaSpecs(entries map[kubefloworgv1.ReplicaType]*kubefloworgv1.ReplicaSpec) *PyTorchJobSpecApplyConfiguration {
	if b.PyTorchReplicaSpecs == nil && len(entries) > 0 {
		b.PyTorchReplicaSpecs = make(map[kubefloworgv1.ReplicaType]*kubefloworgv1.ReplicaSpec, len(entries))
	}
	for k, v := range entries {
		b.PyTorchReplicaSpecs[k] = v
	}
	return b
}

// WithNprocPerNode sets the NprocPerNode field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NprocPerNode field is set to the value of the last call.
func (b *PyTorchJobSpecApplyConfiguration) WithNprocPerNode(value string) *PyTorchJobSpecApplyConfiguration {
	b.NprocPerNode = &value
	return b
}