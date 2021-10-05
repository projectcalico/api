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
	Node string `json:"node,omitempty" validate:"required,name"`

	// Classes declares the types of information to monitor for this calico/node,
	// and allows for selective status reporting about certain subsets of information.
	Classes []NodeStatusClassType `json:"classes,omitempty"`

	// UpdateIntervalInSeconds is the period at which CalicoNodeStatus should be updated.
	// Set to 0 to disable CalicoNodeStatus refresh. [Default: 10]
	UpdateIntervalInSeconds *int `json:"updateIntervalInSeconds,omitempty"`
}

// CalicoNodeStatusStatus defines the observed state of CalicoNodeStatus.
// No validation needed for status since it is updated by Calico.
type CalicoNodeStatusStatus struct {
	// LastUpdated is a timestamp representing the server time when CalicoNodeStatus object
	// last updated. It is represented in RFC3339 form and is in UTC.
	// +nullable
	LastUpdated metav1.Time `json:"lastUpdated,omitempty"`

	// Agent holds agent status on the node.
	Agent CalicoNodeAgentStatus `json:"agent,omitempty"`

	// BGP holds node BGP status.
	BGP CalicoNodeBGPStatus `json:"bgp,omitempty"`

	// Routes reports routes known to the Calico BGP daemon on the node.
	Routes CalicoNodeBGPRouteStatus `json:"routes,omitempty"`
}

// CalicoNodeAgentStatus defines the observed state of agent status on the node.
type CalicoNodeAgentStatus struct {
	// Bird4 represents the latest observed status of bird4.
	Birdv4 CalicoNodeBirdStatus `json:"birdv4,omitempty"`

	// Bird6 represents the latest observed status of bird6.
	Birdv6 CalicoNodeBirdStatus `json:"birdv6,omitempty"`
}

// CalicoNodeBGPStatus defines the observed state of BGP status on the node.
type CalicoNodeBGPStatus struct {
	// The total number of IPv4 established bgp sessions.
	V4NumEstablished int `json:"v4NumEstablished,omitempty"`

	// The total number of IPv4 non-established bgp sessions.
	V4NumNotEstablished int `json:"v4NumNotEstablished,omitempty"`

	// The total number of IPv6 established bgp sessions.
	V6NumEstablished int `json:"v6NumEstablished,omitempty"`

	// The total number of IPv6 non-established bgp sessions.
	V6NumNotEstablished int `json:"v6NumNotEstablished,omitempty"`

	// V4Peers represents IPv4 BGP peers status on the node.
	V4Peers []CalicoNodePeer `json:"v4Peers,omitempty"`

	// V6Peers represents IPv6 BGP peers status on the node.
	V6Peers []CalicoNodePeer `json:"v6Peers,omitempty"`
}

// CalicoNodeBGPRouteStatus defines the observed state of routes status on the node.
type CalicoNodeBGPRouteStatus struct {
	// V4 represents IPv4 routes on the node.
	V4Routes []CalicoNodeRoutes `json:"v4Routes,omitempty"`

	// V6 represents IPv6 routes on the node.
	V6Routes []CalicoNodeRoutes `json:"v6Routes,omitempty"`
}

// CalicoNodeBirdStatus defines the observed state of bird.
type CalicoNodeBirdStatus struct {
	// Ready indicates if bird status is ready.
	Ready bool `json:"ready,omitempty"`

	// Bird version.
	Version string `json:"version,omitempty"`

	// Router ID used by bird.
	RouterID string `json:"routerID,omitempty"`

	// ServerTime holds the value of serverTime from bird.ctl output.
	ServerTime string `json:"serverTime,omitempty"`

	// LastBootTime holds the value of lastBootTime from bird.ctl output.
	LastBootTime string `json:"lastBootTime,omitempty"`

	// LastReconfigTime holds the value of lastReconfigTime from bird.ctl output.
	LastReconfigTime string `json:"lastReconfigTime,omitempty"`
}

// CalicoNodePeer contains the status of BGP peers on the node.
type CalicoNodePeer struct {
	// IP address of the peer whose condition we are reporting.
	PeerIP string `json:"peerIP,omitempty"`

	// Type indicates whether this peer is configured via the node-to-node mesh,
	// or via en explicit global or per-node BGPPeer object.
	Type BGPPeerType `json:"type,omitempty"`

	// The state is the bird bgp session state.
	State string `json:"state,omitempty"`

	// Since the state or reason last changed.
	Since string `json:"since,omitempty"`

	// Extra information from bird on the bgp session.
	Info string `json:"reason,omitempty"`
}

// CalicoNodeRoutes contains the status of BGP routes on the node.
type CalicoNodeRoutes struct {
	// Destination of the route.
	Destination string `json:"destination,omitempty"`

	// Gateway for the destination.
	Gateway string `json:"gateway,omitempty"`

	// Interface for the destination
	Interface string `json:"interface,omitempty"`

	// LearnedFrom indicates who installed this route.
	// If it is populated by a BGP peer, this is the name of the BGPPeer object.
	// If it is populated by node mesh, this is the name of the node.
	// Or it is one of kernel, direct or static.
	LearnedFrom string `json:"learnedFrom,omitempty"`
}

// CalicoNodeRouteLearnedFrom contains the information of the source from which a routes has been learned.
type CalicoNodeRouteLearnedFrom struct {
	// Type of the source where a route is learned from.
	SourceType CalicoNodeRouteSourceType `json:"sourceType,omitempty"`

	// If sourceType is NodeMesh or BGPPeer, IP address of the router that sent us this route.
	PeerIP string `json:"peerIP,omitempty"`

	// If source is a Kubernetes node running Calico, the name of the Kubernetes node that originated the route.
	Node string `json:"node,omitempty" validate:"required,name"`
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

type CalicoNodeRouteSourceType string

const (
	RouteSourceTypeKernel      CalicoNodeRouteSourceType = "Kernel"
	RouteSourceTypeStatic                                = "Static"
	RouteSourceTypeDirect                                = "Direct"
	RouteSourceTypeNodeMesh                              = "NodeMesh"
	RouteSourceTypeNodeBGPPeer                           = "BGPPeer"
)

type NodeStatusClassType string

const (
	NodeStatusClassTypeAgent  NodeStatusClassType = "Agent"
	NodeStatusClassTypeBGP                        = "BGP"
	NodeStatusClassTypeRoutes                     = "Routes"
)

type BGPPeerType string

const (
	BGPPeerTypeNodeMesh   BGPPeerType = "NodeMesh"
	BGPPeerTypeNodePeer               = "NodePeer"
	BGPPeerTypeGlobalPeer             = "GlobalPeer"
)
