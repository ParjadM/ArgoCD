package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1alpha1 "github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
)

// AppProjectLister helps list AppProjects.
type AppProjectLister interface {
	// List lists all AppProjects in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// AppProjects returns an object that can list and get AppProjects.
	AppProjects(namespace string) AppProjectNamespaceLister
	AppProjectListerExpansion
}

// appProjectLister implements the AppProjectLister interface.
type appProjectLister struct {
	indexer cache.Indexer
}

// NewAppProjectLister returns a new AppProjectLister.
func NewAppProjectLister(indexer cache.Indexer) AppProjectLister {
	return &appProjectLister{indexer: indexer}
}

// List lists all AppProjects in the indexer.
func (s *appProjectLister) List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppProject))
	})
	return ret, err
}

// AppProjects returns an object that can list and get AppProjects.
func (s *appProjectLister) AppProjects(namespace string) AppProjectNamespaceLister {
	return appProjectNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AppProjectNamespaceLister helps list and get AppProjects.
type AppProjectNamespaceLister interface {
	// List lists all AppProjects in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error)
	// Get retrieves the AppProject from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.AppProject, error)
	AppProjectNamespaceListerExpansion
}

// appProjectNamespaceLister implements the AppProjectNamespaceLister
// interface.
type appProjectNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AppProjects in the indexer for a given namespace.
func (s appProjectNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.AppProject, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.AppProject))
	})
	return ret, err
}

// Get retrieves the AppProject from the indexer for a given namespace and name.
func (s appProjectNamespaceLister) Get(name string) (*v1alpha1.AppProject, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("appproject"), name)
	}
	return obj.(*v1alpha1.AppProject), nil
}
