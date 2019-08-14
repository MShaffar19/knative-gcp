/*
Copyright 2019 Google LLC

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
	v1alpha1 "github.com/google/knative-gcp/pkg/apis/messaging/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DecoratorLister helps list Decorators.
type DecoratorLister interface {
	// List lists all Decorators in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Decorator, err error)
	// Decorators returns an object that can list and get Decorators.
	Decorators(namespace string) DecoratorNamespaceLister
	DecoratorListerExpansion
}

// decoratorLister implements the DecoratorLister interface.
type decoratorLister struct {
	indexer cache.Indexer
}

// NewDecoratorLister returns a new DecoratorLister.
func NewDecoratorLister(indexer cache.Indexer) DecoratorLister {
	return &decoratorLister{indexer: indexer}
}

// List lists all Decorators in the indexer.
func (s *decoratorLister) List(selector labels.Selector) (ret []*v1alpha1.Decorator, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Decorator))
	})
	return ret, err
}

// Decorators returns an object that can list and get Decorators.
func (s *decoratorLister) Decorators(namespace string) DecoratorNamespaceLister {
	return decoratorNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// DecoratorNamespaceLister helps list and get Decorators.
type DecoratorNamespaceLister interface {
	// List lists all Decorators in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Decorator, err error)
	// Get retrieves the Decorator from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Decorator, error)
	DecoratorNamespaceListerExpansion
}

// decoratorNamespaceLister implements the DecoratorNamespaceLister
// interface.
type decoratorNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Decorators in the indexer for a given namespace.
func (s decoratorNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Decorator, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Decorator))
	})
	return ret, err
}

// Get retrieves the Decorator from the indexer for a given namespace and name.
func (s decoratorNamespaceLister) Get(name string) (*v1alpha1.Decorator, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("decorator"), name)
	}
	return obj.(*v1alpha1.Decorator), nil
}