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
	KindCalicoNodeStatus     = "CalicoNodeStatus"
	KindCalicoNodeStatusList = "CalicoNodeStatusList"
)

// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CalicoNodeStatusList is a list of CalicoNodeStatus resources.
type CalicoNodeStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Items []CalicoNodeStatus `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CalicoNodeStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   CalicoNodeStatusSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status CalicoNodeStatusStatus `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`
}

// CalicoNodeStatusSpec contains the specification for a CalicoNodeStatus resource.
type CalicoNodeStatusSpec struct {
	// The node name identifies the Calico node instance for node status.
	NodeName string `json:"nodeName"`

	// Classes specify type of information CalicoNodeStatus should contain.
	Classes []NodeStatusClassType `json:"classes,omitempty"`

	// UpdateInterval is the period at which CalicoNodeStatus should be updated.
	// Set to 0 to disable CalicoNodeStatus refresh. [Default: 10s]
	UpdateInterval *metav1.Duration `json:"updateInterval,omitempty" configv1timescale:"seconds"`
}

// CalicoNodeStatusStatus defines the observed state of CalicoNodeStatus.
// No validation needed for status since it is updated by Calico.
type CalicoNodeStatusStatus struct {
	// UpdateTimestamp is a timestamp representing the server time when CalicoNodeStatus object
	// last updated. It is represented in RFC3339 form and is in UTC.
	UpdateTimestamp metav1.Time `json:"updateTimestamp,omitempty"`

	// AdditionalInfo is a a human-readable description of the status of last update.
	AdditionalInfo string `json:"additionalInfo,omitempty"`
}

// CalicoNodeBGPStatus defines the observed state of BGP status on the node.
type CalicoNodeBGPStatus struct {
	// The total number of established bgp sessions.
	NumEstablished int `json:"numEstablished,omitempty"`

	// The total number of non-established bgp sessions.
	NumNotEstablished int `json:"numNotEstablished,omitempty"`

	// BirdStatus represents the latest observed status of bird.
	BirdStatus CalicoNodeBirdStatus `json:"birdStatus,omitempty"`

	// Peers represents BGP peers status on the node.
	Peers []CalicoNodePeers `json:"peers,omitempty"`

	// Routes represents routes on the node.
	Routes []CalicoNodeRoutes `json:"routes,omitempty"`
}

// CalicoNodeBirdStatus defines the observed state of bird.
type CalicoNodeBirdStatus struct {
	// Ready indicates if bird status is ready.
	Ready bool `json:"ready,omitempty"`

	// Bird version.
	Version string `json:"version,omitempty"`

	// Route ID used by bird.
	RouteID string `json:"routeID,omitempty"`

	// ServerTime holds the value of serverTime from birdctl output.
	ServerTime string `json:"serverTime,omitempty"`

	// LastBootTime holds the value of lastBootTime from birdctl output.
	LastBootTime string `json:"lastBootTime,omitempty"`

	// LastReconfigTime holds the value of lastReconfigTime from birdctl output.
	LastReconfigTime string `json:"lastReconfigTime,omitempty"`
}

// CalicoNodePeers contains the status of BGP peers on the node.
type CalicoNodePeers struct {
	// IP address of the peer whose condition we are reporting.
	// If port number is given, format should be `[<IPv6>]:port` or `<IPv4>:<port>` for IPv4.
	// If optional port number is not set, and this peer IP and ASNumber belongs to a calico/node
	// with ListenPort set in BGPConfiguration, then we use that port to peer.
	PeerIP string `json:"peerIP,omitempty"`

	// The type is type of bgp session state.
	Type BGPPeerType `json:"type,omitempty"`

	// The state is the bgp session state.
	State string `json:"state,omitempty"`

	// Since is the time since the condition last changed.
	Since string `json:"since,omitempty"`

	// The reason it's in the current state.
	Reason string `json:"info,omitempty"`
}

// CalicoNodePeers contains the status of BGP routes on the node.
type CalicoNodeRoutes struct {
	// Destination of the route.
	Destination string `json:"destination,omitempty"`

	// Gateway for the destination.
	Gateway string `json:"gateway,omitempty"`

	// Interface for the destination
	Interface string `json:"interface,omitempty"`

	// InstalledBy indicates who installed this route.
	// If it is populated by a BGP peer, this is the name of the BGPPeer object.
	// If it is populated by node mesh, this is the name of the node.
	InstalledBy string `json:"installedBy,omitempty"`
}

// NewCalicoNodeStatus creates a new (zeroed) CalicoNodeStatus struct with the TypeMetadata initialised to the current
// version.
func NewCalicoNodeStatus() *CalicoNodeStatus {
	return &CalicoNodeStatus{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindCalicoNodeStatus,
			APIVersion: GroupVersionCurrent,
		},
	}
}

type NodeStatusClassType string

const (
	NodeStatusClassTypeAgent    NodeStatusClassType = "Agent"
	NodeStatusClassTypeBGPPeers                     = "BGPPeers"
	NodeStatusClassTypeRoutes                       = "Routes"
)

type BGPPeerType string

const (
	BGPPeerTypeNodeMesh   BGPPeerType = "NodeMesh"
	BGPPeerTypeNodePeer               = "NodePeer"
	BGPPeerTypeGlobalPeer             = "GlobalPeer"
)
