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

package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	apiv3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"

	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
)

func main() {
	// Create a new config based on kubeconfig file.
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}

	flag.Parse()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Build a clientset based on the provided kubeconfig file.
	cs, err := clientset.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	conditions := []apiv3.NodeBGPStatusCondition{
		{
			PeerIP: "172.16.101.23",
			Type:   "node-to-node-mesh",
			State:  "up",
			Since:  "09:29:58",
			Info:   "Established",
		},
		{
			PeerIP: "172.16.101.24",
			Type:   "node-to-node-mesh",
			State:  "up",
			Since:  "09:29:59",
			Info:   "Established",
		},
		{
			PeerIP: "172.16.101.25",
			Type:   "node-to-node-mesh",
			State:  "up",
			Since:  "09:30:00",
			Info:   "Established",
		},
		{
			PeerIP: "172.16.101.26",
			Type:   "node-to-node-mesh",
			State:  "up",
			Since:  "09:30:01",
			Info:   "Established",
		},
	}

	labels1 := map[string]string{
		"node":        "work-node-1",
		"bgpPeerName": "rack-1",
		"peerIP":      "172.18.0.1",
		"ASNumber":    "65310",
	}

	labels2 := map[string]string{
		"node":        "work-node-2",
		"bgpPeerName": "rack-1",
		"peerIP":      "172.18.0.2",
		"ASNumber":    "65310",
	}

	err = CreateMockNode(cs, "mock-bgp-0", labels1, conditions)
	if err != nil {
		panic(err)
	}

	err = CreateMockNode(cs, "mock-bgp-1", labels2, conditions)
	if err != nil {
		panic(err)
	}
}

func CreateMockNode(cs *clientset.Clientset, name string, labels map[string]string, conditions []apiv3.NodeBGPStatusCondition) error {
	typeMeta := metav1.TypeMeta{Kind: "NodeBGPStatus", APIVersion: "projectcalico.org/v3"}

	status := &apiv3.NodeBGPStatus{
		TypeMeta: typeMeta,
		ObjectMeta: metav1.ObjectMeta{
			Name:   name,
			Labels: labels,
		},
	}

	_, err := cs.ProjectcalicoV3().NodeBGPStatuses().Create(context.Background(), status, metav1.CreateOptions{})
	if err != nil {
		panic(err)
	}

	newStatus, err := cs.ProjectcalicoV3().NodeBGPStatuses().Get(context.Background(), name, metav1.GetOptions{})
	if err != nil {
		return err
	}

	fmt.Printf("%#v\n", newStatus)

	newStatus.TypeMeta = typeMeta
	newStatus.Status = apiv3.NodeBGPStatusStatus{
		Conditions: conditions,
	}

	_, err = cs.ProjectcalicoV3().NodeBGPStatuses().UpdateStatus(context.Background(), newStatus, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

func ListStatus(cs *clientset.Clientset) error {
	// List NodeBGPStatus.
	list, err := cs.ProjectcalicoV3().NodeBGPStatuses().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		return err
	}

	for _, bgp := range list.Items {
		fmt.Printf("%#v\n", bgp)
	}

	return nil
}
