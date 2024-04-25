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

// Code generated by informer-gen. DO NOT EDIT.

package v1beta1

import (
	"context"
	time "time"

	loaddistributionpolicyv1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/agic/loaddistributionpolicy/v1beta1"
	versioned "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/clientset/versioned"
	internalinterfaces "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/agic_crd_client/listers/loaddistributionpolicy/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// LoadDistributionPolicyInformer provides access to a shared informer and lister for
// LoadDistributionPolicies.
type LoadDistributionPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.LoadDistributionPolicyLister
}

type loadDistributionPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewLoadDistributionPolicyInformer constructs a new informer for LoadDistributionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewLoadDistributionPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredLoadDistributionPolicyInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredLoadDistributionPolicyInformer constructs a new informer for LoadDistributionPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredLoadDistributionPolicyInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoaddistributionpoliciesV1beta1().LoadDistributionPolicies(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LoaddistributionpoliciesV1beta1().LoadDistributionPolicies(namespace).Watch(context.TODO(), options)
			},
		},
		&loaddistributionpolicyv1beta1.LoadDistributionPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *loadDistributionPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredLoadDistributionPolicyInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *loadDistributionPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&loaddistributionpolicyv1beta1.LoadDistributionPolicy{}, f.defaultInformer)
}

func (f *loadDistributionPolicyInformer) Lister() v1beta1.LoadDistributionPolicyLister {
	return v1beta1.NewLoadDistributionPolicyLister(f.Informer().GetIndexer())
}
