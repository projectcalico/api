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
	KindIPAMBlock     = "IPAMBlock"
	KindIPAMBlockList = "IPAMBlockList"
)

// IPAMBlockSpec contains the specification for an IPAMBlock resource.
type IPAMBlockSpec struct {
	CIDR           string                `json:"cidr"`
	Affinity       *string               `json:"affinity,omitempty"`
	StrictAffinity bool                  `json:"strictAffinity"`
	Allocations    []*int                `json:"allocations"`
	Unallocated    []int                 `json:"unallocated"`
	Attributes     []AllocationAttribute `json:"attributes"`

	// +optional
	Deleted bool `json:"deleted"`
}

type AllocationAttribute struct {
	AttrPrimary   *string           `json:"handle_id,omitempty"`
	AttrSecondary map[string]string `json:"secondary,omitempty"`
}
