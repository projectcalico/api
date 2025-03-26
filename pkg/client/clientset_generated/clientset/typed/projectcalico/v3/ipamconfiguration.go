// Copyright (c) 2025 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package v3

import (
	"context"
	"time"

	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	scheme "github.com/projectcalico/api/pkg/client/clientset_generated/clientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// IPAMConfigurationsGetter has a method to return a IPAMConfigurationInterface.
// A group's client should implement this interface.
type IPAMConfigurationsGetter interface {
	IPAMConfigurations() IPAMConfigurationInterface
}

// IPAMConfigurationInterface has methods to work with IPAMConfiguration resources.
type IPAMConfigurationInterface interface {
	Create(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.CreateOptions) (*v3.IPAMConfiguration, error)
	Update(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.UpdateOptions) (*v3.IPAMConfiguration, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v3.IPAMConfiguration, error)
	List(ctx context.Context, opts v1.ListOptions) (*v3.IPAMConfigurationList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.IPAMConfiguration, err error)
	IPAMConfigurationExpansion
}

// iPAMConfigurations implements IPAMConfigurationInterface
type iPAMConfigurations struct {
	client rest.Interface
}

// newIPAMConfigurations returns a IPAMConfigurations
func newIPAMConfigurations(c *ProjectcalicoV3Client) *iPAMConfigurations {
	return &iPAMConfigurations{
		client: c.RESTClient(),
	}
}

// Get takes name of the iPAMConfiguration, and returns the corresponding iPAMConfiguration object, and an error if there is any.
func (c *iPAMConfigurations) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.IPAMConfiguration, err error) {
	result = &v3.IPAMConfiguration{}
	err = c.client.Get().
		Resource("ipamconfigurations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of IPAMConfigurations that match those selectors.
func (c *iPAMConfigurations) List(ctx context.Context, opts v1.ListOptions) (result *v3.IPAMConfigurationList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v3.IPAMConfigurationList{}
	err = c.client.Get().
		Resource("ipamconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested iPAMConfigurations.
func (c *iPAMConfigurations) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("ipamconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a iPAMConfiguration and creates it.  Returns the server's representation of the iPAMConfiguration, and an error, if there is any.
func (c *iPAMConfigurations) Create(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.CreateOptions) (result *v3.IPAMConfiguration, err error) {
	result = &v3.IPAMConfiguration{}
	err = c.client.Post().
		Resource("ipamconfigurations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAMConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a iPAMConfiguration and updates it. Returns the server's representation of the iPAMConfiguration, and an error, if there is any.
func (c *iPAMConfigurations) Update(ctx context.Context, iPAMConfiguration *v3.IPAMConfiguration, opts v1.UpdateOptions) (result *v3.IPAMConfiguration, err error) {
	result = &v3.IPAMConfiguration{}
	err = c.client.Put().
		Resource("ipamconfigurations").
		Name(iPAMConfiguration.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(iPAMConfiguration).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the iPAMConfiguration and deletes it. Returns an error if one occurs.
func (c *iPAMConfigurations) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("ipamconfigurations").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *iPAMConfigurations) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("ipamconfigurations").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched iPAMConfiguration.
func (c *iPAMConfigurations) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.IPAMConfiguration, err error) {
	result = &v3.IPAMConfiguration{}
	err = c.client.Patch(pt).
		Resource("ipamconfigurations").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
