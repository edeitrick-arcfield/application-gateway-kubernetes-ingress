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

package fake

import (
	"context"

	v1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/multicluster/multiclusterservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMultiClusterServices implements MultiClusterServiceInterface
type FakeMultiClusterServices struct {
	Fake *FakeMulticlusterservicesV1alpha1
	ns   string
}

var multiclusterservicesResource = v1alpha1.SchemeGroupVersion.WithResource("multiclusterservices")

var multiclusterservicesKind = v1alpha1.SchemeGroupVersion.WithKind("MultiClusterService")

// Get takes name of the multiClusterService, and returns the corresponding multiClusterService object, and an error if there is any.
func (c *FakeMultiClusterServices) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterservicesResource, c.ns, name), &v1alpha1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterService), err
}

// List takes label and field selectors, and returns the list of MultiClusterServices that match those selectors.
func (c *FakeMultiClusterServices) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.MultiClusterServiceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterservicesResource, multiclusterservicesKind, c.ns, opts), &v1alpha1.MultiClusterServiceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.MultiClusterServiceList{ListMeta: obj.(*v1alpha1.MultiClusterServiceList).ListMeta}
	for _, item := range obj.(*v1alpha1.MultiClusterServiceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterServices.
func (c *FakeMultiClusterServices) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterservicesResource, c.ns, opts))

}

// Create takes the representation of a multiClusterService and creates it.  Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *FakeMultiClusterServices) Create(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.CreateOptions) (result *v1alpha1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterservicesResource, c.ns, multiClusterService), &v1alpha1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterService), err
}

// Update takes the representation of a multiClusterService and updates it. Returns the server's representation of the multiClusterService, and an error, if there is any.
func (c *FakeMultiClusterServices) Update(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (result *v1alpha1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterservicesResource, c.ns, multiClusterService), &v1alpha1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterService), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMultiClusterServices) UpdateStatus(ctx context.Context, multiClusterService *v1alpha1.MultiClusterService, opts v1.UpdateOptions) (*v1alpha1.MultiClusterService, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(multiclusterservicesResource, "status", c.ns, multiClusterService), &v1alpha1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterService), err
}

// Delete takes name of the multiClusterService and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterServices) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(multiclusterservicesResource, c.ns, name, opts), &v1alpha1.MultiClusterService{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterServices) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterservicesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.MultiClusterServiceList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterService.
func (c *FakeMultiClusterServices) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.MultiClusterService, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterservicesResource, c.ns, name, pt, data, subresources...), &v1alpha1.MultiClusterService{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.MultiClusterService), err
}
