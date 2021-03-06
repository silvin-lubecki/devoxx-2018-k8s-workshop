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
	v1alpha1 "github.com/simonferquel/devoxx-2018-k8s-workshop/pkg/apis/etcdaas/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ETCDInstanceLister helps list ETCDInstances.
type ETCDInstanceLister interface {
	// List lists all ETCDInstances in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ETCDInstance, err error)
	// ETCDInstances returns an object that can list and get ETCDInstances.
	ETCDInstances(namespace string) ETCDInstanceNamespaceLister
	ETCDInstanceListerExpansion
}

// eTCDInstanceLister implements the ETCDInstanceLister interface.
type eTCDInstanceLister struct {
	indexer cache.Indexer
}

// NewETCDInstanceLister returns a new ETCDInstanceLister.
func NewETCDInstanceLister(indexer cache.Indexer) ETCDInstanceLister {
	return &eTCDInstanceLister{indexer: indexer}
}

// List lists all ETCDInstances in the indexer.
func (s *eTCDInstanceLister) List(selector labels.Selector) (ret []*v1alpha1.ETCDInstance, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ETCDInstance))
	})
	return ret, err
}

// ETCDInstances returns an object that can list and get ETCDInstances.
func (s *eTCDInstanceLister) ETCDInstances(namespace string) ETCDInstanceNamespaceLister {
	return eTCDInstanceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ETCDInstanceNamespaceLister helps list and get ETCDInstances.
type ETCDInstanceNamespaceLister interface {
	// List lists all ETCDInstances in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ETCDInstance, err error)
	// Get retrieves the ETCDInstance from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ETCDInstance, error)
	ETCDInstanceNamespaceListerExpansion
}

// eTCDInstanceNamespaceLister implements the ETCDInstanceNamespaceLister
// interface.
type eTCDInstanceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ETCDInstances in the indexer for a given namespace.
func (s eTCDInstanceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ETCDInstance, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ETCDInstance))
	})
	return ret, err
}

// Get retrieves the ETCDInstance from the indexer for a given namespace and name.
func (s eTCDInstanceNamespaceLister) Get(name string) (*v1alpha1.ETCDInstance, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("etcdinstance"), name)
	}
	return obj.(*v1alpha1.ETCDInstance), nil
}
