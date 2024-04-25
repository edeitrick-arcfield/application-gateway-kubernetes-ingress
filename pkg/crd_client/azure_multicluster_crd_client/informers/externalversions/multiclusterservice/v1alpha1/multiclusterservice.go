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

package v1alpha1

import (
	"context"
	time "time"

	multiclusterservicev1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/multicluster/multiclusterservice/v1alpha1"
	versioned "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/clientset/versioned"
	internalinterfaces "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/listers/multiclusterservice/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// MultiClusterServiceInformer provides access to a shared informer and lister for
// MultiClusterServices.
type MultiClusterServiceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MultiClusterServiceLister
}

type multiClusterServiceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewMultiClusterServiceInformer constructs a new informer for MultiClusterService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMultiClusterServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredMultiClusterServiceInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredMultiClusterServiceInformer constructs a new informer for MultiClusterService type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredMultiClusterServiceInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterservicesV1alpha1().MultiClusterServices(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.MulticlusterservicesV1alpha1().MultiClusterServices(namespace).Watch(context.TODO(), options)
			},
		},
		&multiclusterservicev1alpha1.MultiClusterService{},
		resyncPeriod,
		indexers,
	)
}

func (f *multiClusterServiceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredMultiClusterServiceInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *multiClusterServiceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&multiclusterservicev1alpha1.MultiClusterService{}, f.defaultInformer)
}

func (f *multiClusterServiceInformer) Lister() v1alpha1.MultiClusterServiceLister {
	return v1alpha1.NewMultiClusterServiceLister(f.Informer().GetIndexer())
}
