/*
Copyright 2024 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v1

import (
	"context"
	"time"

	"github.com/rancher/lasso/pkg/client"
	"github.com/rancher/lasso/pkg/controller"
	v1 "github.com/rancher/rancher/pkg/apis/rke.cattle.io/v1"
	"github.com/rancher/wrangler/pkg/apply"
	"github.com/rancher/wrangler/pkg/condition"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/kv"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

type RKEBootstrapHandler func(string, *v1.RKEBootstrap) (*v1.RKEBootstrap, error)

type RKEBootstrapController interface {
	generic.ControllerMeta
	RKEBootstrapClient

	OnChange(ctx context.Context, name string, sync RKEBootstrapHandler)
	OnRemove(ctx context.Context, name string, sync RKEBootstrapHandler)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, duration time.Duration)

	Cache() RKEBootstrapCache
}

type RKEBootstrapClient interface {
	Create(*v1.RKEBootstrap) (*v1.RKEBootstrap, error)
	Update(*v1.RKEBootstrap) (*v1.RKEBootstrap, error)
	UpdateStatus(*v1.RKEBootstrap) (*v1.RKEBootstrap, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	Get(namespace, name string, options metav1.GetOptions) (*v1.RKEBootstrap, error)
	List(namespace string, opts metav1.ListOptions) (*v1.RKEBootstrapList, error)
	Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error)
	Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.RKEBootstrap, err error)
}

type RKEBootstrapCache interface {
	Get(namespace, name string) (*v1.RKEBootstrap, error)
	List(namespace string, selector labels.Selector) ([]*v1.RKEBootstrap, error)

	AddIndexer(indexName string, indexer RKEBootstrapIndexer)
	GetByIndex(indexName, key string) ([]*v1.RKEBootstrap, error)
}

type RKEBootstrapIndexer func(obj *v1.RKEBootstrap) ([]string, error)

type rKEBootstrapController struct {
	controller    controller.SharedController
	client        *client.Client
	gvk           schema.GroupVersionKind
	groupResource schema.GroupResource
}

func NewRKEBootstrapController(gvk schema.GroupVersionKind, resource string, namespaced bool, controller controller.SharedControllerFactory) RKEBootstrapController {
	c := controller.ForResourceKind(gvk.GroupVersion().WithResource(resource), gvk.Kind, namespaced)
	return &rKEBootstrapController{
		controller: c,
		client:     c.Client(),
		gvk:        gvk,
		groupResource: schema.GroupResource{
			Group:    gvk.Group,
			Resource: resource,
		},
	}
}

func FromRKEBootstrapHandlerToHandler(sync RKEBootstrapHandler) generic.Handler {
	return func(key string, obj runtime.Object) (ret runtime.Object, err error) {
		var v *v1.RKEBootstrap
		if obj == nil {
			v, err = sync(key, nil)
		} else {
			v, err = sync(key, obj.(*v1.RKEBootstrap))
		}
		if v == nil {
			return nil, err
		}
		return v, err
	}
}

func (c *rKEBootstrapController) Updater() generic.Updater {
	return func(obj runtime.Object) (runtime.Object, error) {
		newObj, err := c.Update(obj.(*v1.RKEBootstrap))
		if newObj == nil {
			return nil, err
		}
		return newObj, err
	}
}

func UpdateRKEBootstrapDeepCopyOnChange(client RKEBootstrapClient, obj *v1.RKEBootstrap, handler func(obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error)) (*v1.RKEBootstrap, error) {
	if obj == nil {
		return obj, nil
	}

	copyObj := obj.DeepCopy()
	newObj, err := handler(copyObj)
	if newObj != nil {
		copyObj = newObj
	}
	if obj.ResourceVersion == copyObj.ResourceVersion && !equality.Semantic.DeepEqual(obj, copyObj) {
		return client.Update(copyObj)
	}

	return copyObj, err
}

func (c *rKEBootstrapController) AddGenericHandler(ctx context.Context, name string, handler generic.Handler) {
	c.controller.RegisterHandler(ctx, name, controller.SharedControllerHandlerFunc(handler))
}

func (c *rKEBootstrapController) AddGenericRemoveHandler(ctx context.Context, name string, handler generic.Handler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), handler))
}

func (c *rKEBootstrapController) OnChange(ctx context.Context, name string, sync RKEBootstrapHandler) {
	c.AddGenericHandler(ctx, name, FromRKEBootstrapHandlerToHandler(sync))
}

func (c *rKEBootstrapController) OnRemove(ctx context.Context, name string, sync RKEBootstrapHandler) {
	c.AddGenericHandler(ctx, name, generic.NewRemoveHandler(name, c.Updater(), FromRKEBootstrapHandlerToHandler(sync)))
}

func (c *rKEBootstrapController) Enqueue(namespace, name string) {
	c.controller.Enqueue(namespace, name)
}

func (c *rKEBootstrapController) EnqueueAfter(namespace, name string, duration time.Duration) {
	c.controller.EnqueueAfter(namespace, name, duration)
}

func (c *rKEBootstrapController) Informer() cache.SharedIndexInformer {
	return c.controller.Informer()
}

func (c *rKEBootstrapController) GroupVersionKind() schema.GroupVersionKind {
	return c.gvk
}

func (c *rKEBootstrapController) Cache() RKEBootstrapCache {
	return &rKEBootstrapCache{
		indexer:  c.Informer().GetIndexer(),
		resource: c.groupResource,
	}
}

func (c *rKEBootstrapController) Create(obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error) {
	result := &v1.RKEBootstrap{}
	return result, c.client.Create(context.TODO(), obj.Namespace, obj, result, metav1.CreateOptions{})
}

func (c *rKEBootstrapController) Update(obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error) {
	result := &v1.RKEBootstrap{}
	return result, c.client.Update(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *rKEBootstrapController) UpdateStatus(obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error) {
	result := &v1.RKEBootstrap{}
	return result, c.client.UpdateStatus(context.TODO(), obj.Namespace, obj, result, metav1.UpdateOptions{})
}

func (c *rKEBootstrapController) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	if options == nil {
		options = &metav1.DeleteOptions{}
	}
	return c.client.Delete(context.TODO(), namespace, name, *options)
}

func (c *rKEBootstrapController) Get(namespace, name string, options metav1.GetOptions) (*v1.RKEBootstrap, error) {
	result := &v1.RKEBootstrap{}
	return result, c.client.Get(context.TODO(), namespace, name, result, options)
}

func (c *rKEBootstrapController) List(namespace string, opts metav1.ListOptions) (*v1.RKEBootstrapList, error) {
	result := &v1.RKEBootstrapList{}
	return result, c.client.List(context.TODO(), namespace, result, opts)
}

func (c *rKEBootstrapController) Watch(namespace string, opts metav1.ListOptions) (watch.Interface, error) {
	return c.client.Watch(context.TODO(), namespace, opts)
}

func (c *rKEBootstrapController) Patch(namespace, name string, pt types.PatchType, data []byte, subresources ...string) (*v1.RKEBootstrap, error) {
	result := &v1.RKEBootstrap{}
	return result, c.client.Patch(context.TODO(), namespace, name, pt, data, result, metav1.PatchOptions{}, subresources...)
}

type rKEBootstrapCache struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (c *rKEBootstrapCache) Get(namespace, name string) (*v1.RKEBootstrap, error) {
	obj, exists, err := c.indexer.GetByKey(namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(c.resource, name)
	}
	return obj.(*v1.RKEBootstrap), nil
}

func (c *rKEBootstrapCache) List(namespace string, selector labels.Selector) (ret []*v1.RKEBootstrap, err error) {

	err = cache.ListAllByNamespace(c.indexer, namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.RKEBootstrap))
	})

	return ret, err
}

func (c *rKEBootstrapCache) AddIndexer(indexName string, indexer RKEBootstrapIndexer) {
	utilruntime.Must(c.indexer.AddIndexers(map[string]cache.IndexFunc{
		indexName: func(obj interface{}) (strings []string, e error) {
			return indexer(obj.(*v1.RKEBootstrap))
		},
	}))
}

func (c *rKEBootstrapCache) GetByIndex(indexName, key string) (result []*v1.RKEBootstrap, err error) {
	objs, err := c.indexer.ByIndex(indexName, key)
	if err != nil {
		return nil, err
	}
	result = make([]*v1.RKEBootstrap, 0, len(objs))
	for _, obj := range objs {
		result = append(result, obj.(*v1.RKEBootstrap))
	}
	return result, nil
}

type RKEBootstrapStatusHandler func(obj *v1.RKEBootstrap, status v1.RKEBootstrapStatus) (v1.RKEBootstrapStatus, error)

type RKEBootstrapGeneratingHandler func(obj *v1.RKEBootstrap, status v1.RKEBootstrapStatus) ([]runtime.Object, v1.RKEBootstrapStatus, error)

func RegisterRKEBootstrapStatusHandler(ctx context.Context, controller RKEBootstrapController, condition condition.Cond, name string, handler RKEBootstrapStatusHandler) {
	statusHandler := &rKEBootstrapStatusHandler{
		client:    controller,
		condition: condition,
		handler:   handler,
	}
	controller.AddGenericHandler(ctx, name, FromRKEBootstrapHandlerToHandler(statusHandler.sync))
}

func RegisterRKEBootstrapGeneratingHandler(ctx context.Context, controller RKEBootstrapController, apply apply.Apply,
	condition condition.Cond, name string, handler RKEBootstrapGeneratingHandler, opts *generic.GeneratingHandlerOptions) {
	statusHandler := &rKEBootstrapGeneratingHandler{
		RKEBootstrapGeneratingHandler: handler,
		apply:                         apply,
		name:                          name,
		gvk:                           controller.GroupVersionKind(),
	}
	if opts != nil {
		statusHandler.opts = *opts
	}
	controller.OnChange(ctx, name, statusHandler.Remove)
	RegisterRKEBootstrapStatusHandler(ctx, controller, condition, name, statusHandler.Handle)
}

type rKEBootstrapStatusHandler struct {
	client    RKEBootstrapClient
	condition condition.Cond
	handler   RKEBootstrapStatusHandler
}

func (a *rKEBootstrapStatusHandler) sync(key string, obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error) {
	if obj == nil {
		return obj, nil
	}

	origStatus := obj.Status.DeepCopy()
	obj = obj.DeepCopy()
	newStatus, err := a.handler(obj, obj.Status)
	if err != nil {
		// Revert to old status on error
		newStatus = *origStatus.DeepCopy()
	}

	if a.condition != "" {
		if errors.IsConflict(err) {
			a.condition.SetError(&newStatus, "", nil)
		} else {
			a.condition.SetError(&newStatus, "", err)
		}
	}
	if !equality.Semantic.DeepEqual(origStatus, &newStatus) {
		if a.condition != "" {
			// Since status has changed, update the lastUpdatedTime
			a.condition.LastUpdated(&newStatus, time.Now().UTC().Format(time.RFC3339))
		}

		var newErr error
		obj.Status = newStatus
		newObj, newErr := a.client.UpdateStatus(obj)
		if err == nil {
			err = newErr
		}
		if newErr == nil {
			obj = newObj
		}
	}
	return obj, err
}

type rKEBootstrapGeneratingHandler struct {
	RKEBootstrapGeneratingHandler
	apply apply.Apply
	opts  generic.GeneratingHandlerOptions
	gvk   schema.GroupVersionKind
	name  string
}

func (a *rKEBootstrapGeneratingHandler) Remove(key string, obj *v1.RKEBootstrap) (*v1.RKEBootstrap, error) {
	if obj != nil {
		return obj, nil
	}

	obj = &v1.RKEBootstrap{}
	obj.Namespace, obj.Name = kv.RSplit(key, "/")
	obj.SetGroupVersionKind(a.gvk)

	return nil, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects()
}

func (a *rKEBootstrapGeneratingHandler) Handle(obj *v1.RKEBootstrap, status v1.RKEBootstrapStatus) (v1.RKEBootstrapStatus, error) {
	if !obj.DeletionTimestamp.IsZero() {
		return status, nil
	}

	objs, newStatus, err := a.RKEBootstrapGeneratingHandler(obj, status)
	if err != nil {
		return newStatus, err
	}

	return newStatus, generic.ConfigureApplyForObject(a.apply, obj, &a.opts).
		WithOwner(obj).
		WithSetID(a.name).
		ApplyObjects(objs...)
}
