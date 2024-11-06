// Copyright (c) 2024 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"net/http"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"github.com/projectcalico/api/pkg/client/clientset_generated/clientset/scheme"
	rest "k8s.io/client-go/rest"
)

type ProjectcalicoV3Interface interface {
	RESTClient() rest.Interface
	BGPConfigurationsGetter
	BGPFiltersGetter
	BGPPeersGetter
	BlockAffinitiesGetter
	CalicoNodeStatusesGetter
	ClusterInformationsGetter
	FelixConfigurationsGetter
	GlobalNetworkPoliciesGetter
	GlobalNetworkSetsGetter
	HostEndpointsGetter
	IPAMConfigurationsGetter
	IPPoolsGetter
	IPReservationsGetter
	KubeControllersConfigurationsGetter
	NetworkPoliciesGetter
	NetworkSetsGetter
	ProfilesGetter
	TiersGetter
}

// ProjectcalicoV3Client is used to interact with features provided by the projectcalico.org group.
type ProjectcalicoV3Client struct {
	restClient rest.Interface
}

func (c *ProjectcalicoV3Client) BGPConfigurations() BGPConfigurationInterface {
	return newBGPConfigurations(c)
}

func (c *ProjectcalicoV3Client) BGPFilters() BGPFilterInterface {
	return newBGPFilters(c)
}

func (c *ProjectcalicoV3Client) BGPPeers() BGPPeerInterface {
	return newBGPPeers(c)
}

func (c *ProjectcalicoV3Client) BlockAffinities() BlockAffinityInterface {
	return newBlockAffinities(c)
}

func (c *ProjectcalicoV3Client) CalicoNodeStatuses() CalicoNodeStatusInterface {
	return newCalicoNodeStatuses(c)
}

func (c *ProjectcalicoV3Client) ClusterInformations() ClusterInformationInterface {
	return newClusterInformations(c)
}

func (c *ProjectcalicoV3Client) FelixConfigurations() FelixConfigurationInterface {
	return newFelixConfigurations(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkPolicies() GlobalNetworkPolicyInterface {
	return newGlobalNetworkPolicies(c)
}

func (c *ProjectcalicoV3Client) GlobalNetworkSets() GlobalNetworkSetInterface {
	return newGlobalNetworkSets(c)
}

func (c *ProjectcalicoV3Client) HostEndpoints() HostEndpointInterface {
	return newHostEndpoints(c)
}

func (c *ProjectcalicoV3Client) IPAMConfigurations() IPAMConfigurationInterface {
	return newIPAMConfigurations(c)
}

func (c *ProjectcalicoV3Client) IPPools() IPPoolInterface {
	return newIPPools(c)
}

func (c *ProjectcalicoV3Client) IPReservations() IPReservationInterface {
	return newIPReservations(c)
}

func (c *ProjectcalicoV3Client) KubeControllersConfigurations() KubeControllersConfigurationInterface {
	return newKubeControllersConfigurations(c)
}

func (c *ProjectcalicoV3Client) NetworkPolicies(namespace string) NetworkPolicyInterface {
	return newNetworkPolicies(c, namespace)
}

func (c *ProjectcalicoV3Client) NetworkSets(namespace string) NetworkSetInterface {
	return newNetworkSets(c, namespace)
}

func (c *ProjectcalicoV3Client) Profiles() ProfileInterface {
	return newProfiles(c)
}

func (c *ProjectcalicoV3Client) Tiers() TierInterface {
	return newTiers(c)
}

// NewForConfig creates a new ProjectcalicoV3Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*ProjectcalicoV3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new ProjectcalicoV3Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*ProjectcalicoV3Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &ProjectcalicoV3Client{client}, nil
}

// NewForConfigOrDie creates a new ProjectcalicoV3Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *ProjectcalicoV3Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new ProjectcalicoV3Client for the given RESTClient.
func New(c rest.Interface) *ProjectcalicoV3Client {
	return &ProjectcalicoV3Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v3.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *ProjectcalicoV3Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
