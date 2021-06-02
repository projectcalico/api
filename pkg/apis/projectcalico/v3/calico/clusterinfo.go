// Copyright (c) 2017, 20201 Tigera, Inc. All rights reserved.

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
	KindClusterInformation     = "ClusterInformation"
	KindClusterInformationList = "ClusterInformationList"
)

// ClusterInformationSpec contains the values of describing the cluster.
type ClusterInformationSpec struct {
	// ClusterGUID is the GUID of the cluster
	ClusterGUID string `json:"clusterGUID,omitempty" validate:"omitempty"`
	// ClusterType describes the type of the cluster
	ClusterType string `json:"clusterType,omitempty" validate:"omitempty"`
	// CalicoVersion is the version of Calico that the cluster is running
	CalicoVersion string `json:"calicoVersion,omitempty" validate:"omitempty"`
	// DatastoreReady is used during significant datastore migrations to signal to components
	// such as Felix that it should wait before accessing the datastore.
	DatastoreReady *bool `json:"datastoreReady,omitempty"`
	// Variant declares which variant of Calico should be active.
	Variant string `json:"variant,omitempty"`
}
