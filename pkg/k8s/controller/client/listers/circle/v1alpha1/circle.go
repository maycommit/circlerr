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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/maycommit/circlerr/pkg/k8s/controller/apis/circle/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CircleLister helps list Circles.
// All objects returned here must be treated as read-only.
type CircleLister interface {
	// List lists all Circles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Circle, err error)
	// Circles returns an object that can list and get Circles.
	Circles(namespace string) CircleNamespaceLister
	CircleListerExpansion
}

// circleLister implements the CircleLister interface.
type circleLister struct {
	indexer cache.Indexer
}

// NewCircleLister returns a new CircleLister.
func NewCircleLister(indexer cache.Indexer) CircleLister {
	return &circleLister{indexer: indexer}
}

// List lists all Circles in the indexer.
func (s *circleLister) List(selector labels.Selector) (ret []*v1alpha1.Circle, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Circle))
	})
	return ret, err
}

// Circles returns an object that can list and get Circles.
func (s *circleLister) Circles(namespace string) CircleNamespaceLister {
	return circleNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CircleNamespaceLister helps list and get Circles.
// All objects returned here must be treated as read-only.
type CircleNamespaceLister interface {
	// List lists all Circles in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Circle, err error)
	// Get retrieves the Circle from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Circle, error)
	CircleNamespaceListerExpansion
}

// circleNamespaceLister implements the CircleNamespaceLister
// interface.
type circleNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Circles in the indexer for a given namespace.
func (s circleNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Circle, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Circle))
	})
	return ret, err
}

// Get retrieves the Circle from the indexer for a given namespace and name.
func (s circleNamespaceLister) Get(name string) (*v1alpha1.Circle, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("circle"), name)
	}
	return obj.(*v1alpha1.Circle), nil
}
