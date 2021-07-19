// Copyright (c) 2021 Tigera, Inc. All rights reserved.

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

package v3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	KindNodeBGPStatus     = "NodeBGPStatus"
	KindNodeBGPStatusList = "NodeBGPStatusList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeBGPStatusList is a list of NodeBGPStatus resources.
type NodeBGPStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []NodeBGPStatus `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeBGPStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Status NodeBGPStatusStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// NodeBGPStatusStatus defines the observed state of NodeBGPStatus
type NodeBGPStatusStatus struct {
	// The total number of established bgp sessions.
	NumEstablished int `json:"numEstablished,omitempty"`

	// The total number of non-established bgp sessions.
	NumNotEstablished int `json:"numNotEstablished,omitempty"`

	// Conditions represents the latest observed set of conditions for this component.
	Conditions []NodeBGPStatusCondition `json:"conditions"`
}

// NodeBGPStatusCondition contains the status for a NodeBGPStatus resource.
// +k8s:deepcopy-gen=true
type NodeBGPStatusCondition struct {
	// IP address of the peer whose condition we are reporting.
	// If port number is given, format should be `[<IPv6>]:port` or `<IPv4>:<port>` for IPv4.
	// If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node
	// with ListenPort set in BGPConfiguration, then we use that port to peer.
	// +optional
	PeerIP string `json:"peerIP,omitempty"`

	// The type is type of bgp session state.
	// +optional
	Type BGPPeerType `json:"type,omitempty"`

	// The state is the bgp session state.
	// +optional
	State string `json:"state,omitempty"`

	// Since is the time since the condition last changed.
	// +optional
	Since string `json:"since,omitempty"`

	// The reason it's in the current state.
	// +optional
	Reason string `json:"info,omitempty"`
}

// NewNodeBGPStatus creates a new (zeroed) NodeBGPStatus struct with the TypeMetadata initialised to the current
// version.
func NewNodeBGPStatus() *NodeBGPStatus {
	return &NodeBGPStatus{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindBGPPeer,
			APIVersion: GroupVersionCurrent,
		},
	}
}

type BGPPeerType string

const (
	BGPPeerTypeNodeMesh   BGPPeerType = "NodeMesh"
	BGPPeerTypeNodePeer               = "NodePeer"
	BGPPeerTypeGlobalPeer             = "GlobalPeer"
)
