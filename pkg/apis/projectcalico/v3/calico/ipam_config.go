// Copyright (c) 2019, 2021 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package calico

const (
	KindIPAMConfig     = "IPAMConfig"
	KindIPAMConfigList = "IPAMConfigList"
)

// IPAMConfigSpec contains the specification for an IPAMConfig resource.
type IPAMConfigSpec struct {
	StrictAffinity     bool `json:"strictAffinity"`
	AutoAllocateBlocks bool `json:"autoAllocateBlocks"`

	// MaxBlocksPerHost, if non-zero, is the max number of blocks that can be
	// affine to each host.
	// +optional
	MaxBlocksPerHost int `json:"maxBlocksPerHost,omitempty"`
}
