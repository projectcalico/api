// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by informer-gen. DO NOT EDIT.

package externalversions

import (
	"fmt"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	cache "k8s.io/client-go/tools/cache"
)

// GenericInformer is type of SharedIndexInformer which will locate and delegate to other
// sharedInformers based on type
type GenericInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() cache.GenericLister
}

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (GenericInformer, error) {
	switch resource {
	// Group=projectcalico.org, Version=v3
	case v3.SchemeGroupVersion.WithResource("bgpconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().BGPConfigurations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("bgpfilters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().BGPFilters().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("bgppeers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().BGPPeers().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("blockaffinities"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().BlockAffinities().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("caliconodestatuses"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().CalicoNodeStatuses().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("clusterinformations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().ClusterInformations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("felixconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().FelixConfigurations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("globalnetworkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().GlobalNetworkPolicies().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("globalnetworksets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().GlobalNetworkSets().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("hostendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().HostEndpoints().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("ipamconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().IPAMConfigurations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("ippools"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().IPPools().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("ipreservations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().IPReservations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("kubecontrollersconfigurations"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().KubeControllersConfigurations().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("networkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().NetworkPolicies().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("networksets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().NetworkSets().Informer()}, nil
	case v3.SchemeGroupVersion.WithResource("profiles"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Projectcalico().V3().Profiles().Informer()}, nil

	}

	return nil, fmt.Errorf("no informer found for %v", resource)
}
