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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	samplecontrollerv1alpha1 "charlescd/pkg/apis/circle/v1alpha1"
	versioned "charlescd/pkg/client/clientset/versioned"
	internalinterfaces "charlescd/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "charlescd/pkg/client/listers/circle/v1alpha1"
)

// CircleInformer provides access to a shared informer and lister for
// Circles.
type CircleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CircleLister
}

type circleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCircleInformer constructs a new informer for Circle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory circletprint and number of connections to the server.
func NewCircleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCircleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCircleInformer constructs a new informer for Circle type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory circletprint and number of connections to the server.
func NewFilteredCircleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CircleV1alpha1().Circles(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CircleV1alpha1().Circles(namespace).Watch(context.TODO(), options)
			},
		},
		&samplecontrollerv1alpha1.Circle{},
		resyncPeriod,
		indexers,
	)
}

func (f *circleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCircleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *circleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&samplecontrollerv1alpha1.Circle{}, f.defaultInformer)
}

func (f *circleInformer) Lister() v1alpha1.CircleLister {
	return v1alpha1.NewCircleLister(f.Informer().GetIndexer())
}
