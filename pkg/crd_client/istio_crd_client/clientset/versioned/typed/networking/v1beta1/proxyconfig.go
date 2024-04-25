/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	"time"

	scheme "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/istio_crd_client/clientset/versioned/scheme"
	v1beta1 "istio.io/client-go/pkg/apis/networking/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ProxyConfigsGetter has a method to return a ProxyConfigInterface.
// A group's client should implement this interface.
type ProxyConfigsGetter interface {
	ProxyConfigs(namespace string) ProxyConfigInterface
}

// ProxyConfigInterface has methods to work with ProxyConfig resources.
type ProxyConfigInterface interface {
	Create(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.CreateOptions) (*v1beta1.ProxyConfig, error)
	Update(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (*v1beta1.ProxyConfig, error)
	UpdateStatus(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (*v1beta1.ProxyConfig, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ProxyConfig, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ProxyConfigList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ProxyConfig, err error)
	ProxyConfigExpansion
}

// proxyConfigs implements ProxyConfigInterface
type proxyConfigs struct {
	client rest.Interface
	ns     string
}

// newProxyConfigs returns a ProxyConfigs
func newProxyConfigs(c *NetworkingV1beta1Client, namespace string) *proxyConfigs {
	return &proxyConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the proxyConfig, and returns the corresponding proxyConfig object, and an error if there is any.
func (c *proxyConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ProxyConfig, err error) {
	result = &v1beta1.ProxyConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxyconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ProxyConfigs that match those selectors.
func (c *proxyConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ProxyConfigList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ProxyConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("proxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested proxyConfigs.
func (c *proxyConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("proxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a proxyConfig and creates it.  Returns the server's representation of the proxyConfig, and an error, if there is any.
func (c *proxyConfigs) Create(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.CreateOptions) (result *v1beta1.ProxyConfig, err error) {
	result = &v1beta1.ProxyConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("proxyconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(proxyConfig).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a proxyConfig and updates it. Returns the server's representation of the proxyConfig, and an error, if there is any.
func (c *proxyConfigs) Update(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (result *v1beta1.ProxyConfig, err error) {
	result = &v1beta1.ProxyConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxyconfigs").
		Name(proxyConfig.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(proxyConfig).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *proxyConfigs) UpdateStatus(ctx context.Context, proxyConfig *v1beta1.ProxyConfig, opts v1.UpdateOptions) (result *v1beta1.ProxyConfig, err error) {
	result = &v1beta1.ProxyConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("proxyconfigs").
		Name(proxyConfig.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(proxyConfig).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the proxyConfig and deletes it. Returns an error if one occurs.
func (c *proxyConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxyconfigs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *proxyConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("proxyconfigs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched proxyConfig.
func (c *proxyConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ProxyConfig, err error) {
	result = &v1beta1.ProxyConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("proxyconfigs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
