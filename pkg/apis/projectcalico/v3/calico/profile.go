// Copyright (c) 2017,2021 Tigera, Inc. All rights reserved.

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
	KindProfile     = "Profile"
	KindProfileList = "ProfileList"
)

// ProfileSpec contains the specification for a security Profile resource.
type ProfileSpec struct {
	// The ordered set of ingress rules.  Each rule contains a set of packet match criteria and
	// a corresponding action to apply.
	Ingress []Rule `json:"ingress,omitempty" validate:"omitempty,dive"`
	// The ordered set of egress rules.  Each rule contains a set of packet match criteria and
	// a corresponding action to apply.
	Egress []Rule `json:"egress,omitempty" validate:"omitempty,dive"`
	// An option set of labels to apply to each endpoint (in addition to their own labels)
	// referencing this profile.  If labels configured on the endpoint have keys matching those
	// labels inherited from the profile, the endpoint label values take precedence.
	LabelsToApply map[string]string `json:"labelsToApply,omitempty" validate:"omitempty,labels"`
}
