/*
Copyright 2019 Gravitational, Inc.

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
	time "time"

	wormholegravitationaliov1beta1 "github.com/gravitational/wormhole/pkg/apis/wormhole.gravitational.io/v1beta1"
	versioned "github.com/gravitational/wormhole/pkg/client/clientset/versioned"
	internalinterfaces "github.com/gravitational/wormhole/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "github.com/gravitational/wormhole/pkg/client/listers/wormhole.gravitational.io/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// WgnodeInformer provides access to a shared informer and lister for
// Wgnodes.
type WgnodeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.WgnodeLister
}

type wgnodeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewWgnodeInformer constructs a new informer for Wgnode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewWgnodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredWgnodeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredWgnodeInformer constructs a new informer for Wgnode type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredWgnodeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WormholeV1beta1().Wgnodes(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.WormholeV1beta1().Wgnodes(namespace).Watch(options)
			},
		},
		&wormholegravitationaliov1beta1.Wgnode{},
		resyncPeriod,
		indexers,
	)
}

func (f *wgnodeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredWgnodeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *wgnodeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&wormholegravitationaliov1beta1.Wgnode{}, f.defaultInformer)
}

func (f *wgnodeInformer) Lister() v1beta1.WgnodeLister {
	return v1beta1.NewWgnodeLister(f.Informer().GetIndexer())
}