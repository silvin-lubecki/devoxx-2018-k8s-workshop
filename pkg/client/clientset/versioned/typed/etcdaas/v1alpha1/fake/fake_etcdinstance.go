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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/simonferquel/devoxx-2018-k8s-workshop/pkg/apis/etcdaas/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeETCDInstances implements ETCDInstanceInterface
type FakeETCDInstances struct {
	Fake *FakeEtcdaasV1alpha1
	ns   string
}

var etcdinstancesResource = schema.GroupVersionResource{Group: "etcdaas", Version: "v1alpha1", Resource: "etcdinstances"}

var etcdinstancesKind = schema.GroupVersionKind{Group: "etcdaas", Version: "v1alpha1", Kind: "ETCDInstance"}

// Get takes name of the eTCDInstance, and returns the corresponding eTCDInstance object, and an error if there is any.
func (c *FakeETCDInstances) Get(name string, options v1.GetOptions) (result *v1alpha1.ETCDInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(etcdinstancesResource, c.ns, name), &v1alpha1.ETCDInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDInstance), err
}

// List takes label and field selectors, and returns the list of ETCDInstances that match those selectors.
func (c *FakeETCDInstances) List(opts v1.ListOptions) (result *v1alpha1.ETCDInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(etcdinstancesResource, etcdinstancesKind, c.ns, opts), &v1alpha1.ETCDInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ETCDInstanceList{}
	for _, item := range obj.(*v1alpha1.ETCDInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eTCDInstances.
func (c *FakeETCDInstances) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(etcdinstancesResource, c.ns, opts))

}

// Create takes the representation of a eTCDInstance and creates it.  Returns the server's representation of the eTCDInstance, and an error, if there is any.
func (c *FakeETCDInstances) Create(eTCDInstance *v1alpha1.ETCDInstance) (result *v1alpha1.ETCDInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(etcdinstancesResource, c.ns, eTCDInstance), &v1alpha1.ETCDInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDInstance), err
}

// Update takes the representation of a eTCDInstance and updates it. Returns the server's representation of the eTCDInstance, and an error, if there is any.
func (c *FakeETCDInstances) Update(eTCDInstance *v1alpha1.ETCDInstance) (result *v1alpha1.ETCDInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(etcdinstancesResource, c.ns, eTCDInstance), &v1alpha1.ETCDInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeETCDInstances) UpdateStatus(eTCDInstance *v1alpha1.ETCDInstance) (*v1alpha1.ETCDInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(etcdinstancesResource, "status", c.ns, eTCDInstance), &v1alpha1.ETCDInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDInstance), err
}

// Delete takes name of the eTCDInstance and deletes it. Returns an error if one occurs.
func (c *FakeETCDInstances) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(etcdinstancesResource, c.ns, name), &v1alpha1.ETCDInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeETCDInstances) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(etcdinstancesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ETCDInstanceList{})
	return err
}

// Patch applies the patch and returns the patched eTCDInstance.
func (c *FakeETCDInstances) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ETCDInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(etcdinstancesResource, c.ns, name, data, subresources...), &v1alpha1.ETCDInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ETCDInstance), err
}
