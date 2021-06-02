package main

import (
	"context"
	"flag"
	"fmt"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset"
)

func main() {
	// Create a new config based on kubeconfig file.
	var kubeconfig *string
	kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
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

	// List global network policies.
	list, err := cs.ProjectcalicoV3().GlobalNetworkPolicies().List(context.Background(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, gnp := range list.Items {
		fmt.Printf("%#v\n", gnp)
	}
}
