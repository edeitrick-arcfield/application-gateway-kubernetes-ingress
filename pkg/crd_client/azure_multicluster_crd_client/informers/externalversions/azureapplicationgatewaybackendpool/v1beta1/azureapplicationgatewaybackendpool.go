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

	azureapplicationgatewaybackendpoolv1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/apis/azureapplicationgatewaybackendpool/v1beta1"
	versioned "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/clientset/versioned"
	internalinterfaces "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/Azure/application-gateway-kubernetes-ingress/pkg/crd_client/azure_multicluster_crd_client/listers/azureapplicationgatewaybackendpool/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// AzureApplicationGatewayBackendPoolInformer provides access to a shared informer and lister for
// AzureApplicationGatewayBackendPools.
type AzureApplicationGatewayBackendPoolInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.AzureApplicationGatewayBackendPoolLister
}

type azureApplicationGatewayBackendPoolInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewAzureApplicationGatewayBackendPoolInformer constructs a new informer for AzureApplicationGatewayBackendPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewAzureApplicationGatewayBackendPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredAzureApplicationGatewayBackendPoolInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredAzureApplicationGatewayBackendPoolInformer constructs a new informer for AzureApplicationGatewayBackendPool type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredAzureApplicationGatewayBackendPoolInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzureapplicationgatewaybackendpoolsV1beta1().AzureApplicationGatewayBackendPools().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AzureapplicationgatewaybackendpoolsV1beta1().AzureApplicationGatewayBackendPools().Watch(context.TODO(), options)
			},
		},
		&azureapplicationgatewaybackendpoolv1beta1.AzureApplicationGatewayBackendPool{},
		resyncPeriod,
		indexers,
	)
}

func (f *azureApplicationGatewayBackendPoolInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredAzureApplicationGatewayBackendPoolInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *azureApplicationGatewayBackendPoolInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&azureapplicationgatewaybackendpoolv1beta1.AzureApplicationGatewayBackendPool{}, f.defaultInformer)
}

func (f *azureApplicationGatewayBackendPoolInformer) Lister() v1beta1.AzureApplicationGatewayBackendPoolLister {
	return v1beta1.NewAzureApplicationGatewayBackendPoolLister(f.Informer().GetIndexer())
}
