// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIPAMConfigurations implements IPAMConfigurationInterface
type FakeIPAMConfigurations struct {
	Fake *FakeProjectcalicoV3
}

var ipamconfigurationsResource = v3.SchemeGroupVersion.WithResource("ipamconfigurations")

var ipamconfigurationsKind = v3.SchemeGroupVersion.WithKind("IPAMConfiguration")

// Get takes name of the iPAMConfiguration, and returns the corresponding iPAMConfiguration object, and an error if there is any.
func (c *FakeIPAMConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.IPAMConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ipamconfigurationsResource, name), &v3.IPAMConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.IPAMConfiguration), err
}

// List takes label and field selectors, and returns the list of IPAMConfigurations that match those selectors.
func (c *FakeIPAMConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.IPAMConfigurationList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ipamconfigurationsResource, ipamconfigurationsKind, opts), &v3.IPAMConfigurationList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.IPAMConfigurationList{ListMeta: obj.(*v3.IPAMConfigurationList).ListMeta}
	for _, item := range obj.(*v3.IPAMConfigurationList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAMConfigurations.
func (c *FakeIPAMConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ipamconfigurationsResource, opts))
}

// Create takes the representation of a iPAMConfiguration and creates it.  Returns the server's representation of the iPAMConfiguration, and an error, if there is any.
func (c *FakeIPAMConfigurations) Create(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.CreateOptions) (result *v3.IPAMConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ipamconfigurationsResource, iPAMConfiguration), &v3.IPAMConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.IPAMConfiguration), err
}

// Update takes the representation of a iPAMConfiguration and updates it. Returns the server's representation of the iPAMConfiguration, and an error, if there is any.
func (c *FakeIPAMConfigurations) Update(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.UpdateOptions) (result *v3.IPAMConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ipamconfigurationsResource, iPAMConfiguration), &v3.IPAMConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.IPAMConfiguration), err
}

// Delete takes name of the iPAMConfiguration and deletes it. Returns an error if one occurs.
func (c *FakeIPAMConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ipamconfigurationsResource, name, opts), &v3.IPAMConfiguration{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPAMConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ipamconfigurationsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.IPAMConfigurationList{})
	return err
}

// Patch applies the patch and returns the patched iPAMConfiguration.
func (c *FakeIPAMConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.IPAMConfiguration, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ipamconfigurationsResource, name, pt, data, subresources...), &v3.IPAMConfiguration{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.IPAMConfiguration), err
}
