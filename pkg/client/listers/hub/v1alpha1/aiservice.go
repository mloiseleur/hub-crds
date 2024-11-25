/*
The GNU AFFERO GENERAL PUBLIC LICENSE

Copyright (c) 2020-2024 Traefik Labs

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published
by the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/traefik/hub-crds/pkg/apis/hub/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AIServiceLister helps list AIServices.
// All objects returned here must be treated as read-only.
type AIServiceLister interface {
	// List lists all AIServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AIService, err error)
	// AIServices returns an object that can list and get AIServices.
	AIServices(namespace string) AIServiceNamespaceLister
	AIServiceListerExpansion
}

// aIServiceLister implements the AIServiceLister interface.
type aIServiceLister struct {
	indexer cache.Indexer
}

// NewAIServiceLister returns a new AIServiceLister.
func NewAIServiceLister(indexer cache.Indexer) AIServiceLister {
	return &aIServiceLister{indexer: indexer}
}

// List lists all AIServices in the indexer.
func (s *aIServiceLister) List(selector labels.Selector) (ret []*v1alpha1.AIService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AIService))
	})
	return ret, err
}

// AIServices returns an object that can list and get AIServices.
func (s *aIServiceLister) AIServices(namespace string) AIServiceNamespaceLister {
	return aIServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AIServiceNamespaceLister helps list and get AIServices.
// All objects returned here must be treated as read-only.
type AIServiceNamespaceLister interface {
	// List lists all AIServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.AIService, err error)
	// Get retrieves the AIService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.AIService, error)
	AIServiceNamespaceListerExpansion
}

// aIServiceNamespaceLister implements the AIServiceNamespaceLister
// interface.
type aIServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AIServices in the indexer for a given namespace.
func (s aIServiceNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AIService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AIService))
	})
	return ret, err
}

// Get retrieves the AIService from the indexer for a given namespace and name.
func (s aIServiceNamespaceLister) Get(name string) (*v1alpha1.AIService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("aiservice"), name)
	}
	return obj.(*v1alpha1.AIService), nil
}