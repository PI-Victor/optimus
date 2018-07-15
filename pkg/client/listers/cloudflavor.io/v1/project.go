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

package v1

import (
	v1 "github.com/PI-Victor/pipelines/pkg/apis/cloudflavor.io/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ProjectLister helps list Projects.
type ProjectLister interface {
	// List lists all Projects in the indexer.
	List(selector labels.Selector) (ret []*v1.Project, err error)
	// Projects returns an object that can list and get Projects.
	Projects(namespace string) ProjectNamespaceLister
	ProjectListerExpansion
}

// projectLister implements the ProjectLister interface.
type projectLister struct {
	indexer cache.Indexer
}

// NewProjectLister returns a new ProjectLister.
func NewProjectLister(indexer cache.Indexer) ProjectLister {
	return &projectLister{indexer: indexer}
}

// List lists all Projects in the indexer.
func (s *projectLister) List(selector labels.Selector) (ret []*v1.Project, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Project))
	})
	return ret, err
}

// Projects returns an object that can list and get Projects.
func (s *projectLister) Projects(namespace string) ProjectNamespaceLister {
	return projectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ProjectNamespaceLister helps list and get Projects.
type ProjectNamespaceLister interface {
	// List lists all Projects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.Project, err error)
	// Get retrieves the Project from the indexer for a given namespace and name.
	Get(name string) (*v1.Project, error)
	ProjectNamespaceListerExpansion
}

// projectNamespaceLister implements the ProjectNamespaceLister
// interface.
type projectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Projects in the indexer for a given namespace.
func (s projectNamespaceLister) List(selector labels.Selector) (ret []*v1.Project, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Project))
	})
	return ret, err
}

// Get retrieves the Project from the indexer for a given namespace and name.
func (s projectNamespaceLister) Get(name string) (*v1.Project, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("project"), name)
	}
	return obj.(*v1.Project), nil
}
