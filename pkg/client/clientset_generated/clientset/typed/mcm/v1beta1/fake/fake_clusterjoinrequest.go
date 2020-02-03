// Licensed Materials - Property of IBM
// (c) Copyright IBM Corporation 2018. All Rights Reserved.
// Note to U.S. Government Users Restricted Rights:
// Use, duplication or disclosure restricted by GSA ADP Schedule
// Contract with IBM Corp.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "github.ibm.com/IBMPrivateCloud/multicloud-operators-foundation/pkg/apis/mcm/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterJoinRequests implements ClusterJoinRequestInterface
type FakeClusterJoinRequests struct {
	Fake *FakeMcmV1beta1
}

var clusterjoinrequestsResource = schema.GroupVersionResource{Group: "mcm.ibm.com", Version: "v1beta1", Resource: "clusterjoinrequests"}

var clusterjoinrequestsKind = schema.GroupVersionKind{Group: "mcm.ibm.com", Version: "v1beta1", Kind: "ClusterJoinRequest"}

// Get takes name of the clusterJoinRequest, and returns the corresponding clusterJoinRequest object, and an error if there is any.
func (c *FakeClusterJoinRequests) Get(name string, options v1.GetOptions) (result *v1beta1.ClusterJoinRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterjoinrequestsResource, name), &v1beta1.ClusterJoinRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterJoinRequest), err
}

// List takes label and field selectors, and returns the list of ClusterJoinRequests that match those selectors.
func (c *FakeClusterJoinRequests) List(opts v1.ListOptions) (result *v1beta1.ClusterJoinRequestList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterjoinrequestsResource, clusterjoinrequestsKind, opts), &v1beta1.ClusterJoinRequestList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterJoinRequestList{ListMeta: obj.(*v1beta1.ClusterJoinRequestList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterJoinRequestList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterJoinRequests.
func (c *FakeClusterJoinRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterjoinrequestsResource, opts))
}

// Create takes the representation of a clusterJoinRequest and creates it.  Returns the server's representation of the clusterJoinRequest, and an error, if there is any.
func (c *FakeClusterJoinRequests) Create(clusterJoinRequest *v1beta1.ClusterJoinRequest) (result *v1beta1.ClusterJoinRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterjoinrequestsResource, clusterJoinRequest), &v1beta1.ClusterJoinRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterJoinRequest), err
}

// Update takes the representation of a clusterJoinRequest and updates it. Returns the server's representation of the clusterJoinRequest, and an error, if there is any.
func (c *FakeClusterJoinRequests) Update(clusterJoinRequest *v1beta1.ClusterJoinRequest) (result *v1beta1.ClusterJoinRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterjoinrequestsResource, clusterJoinRequest), &v1beta1.ClusterJoinRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterJoinRequest), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterJoinRequests) UpdateStatus(clusterJoinRequest *v1beta1.ClusterJoinRequest) (*v1beta1.ClusterJoinRequest, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterjoinrequestsResource, "status", clusterJoinRequest), &v1beta1.ClusterJoinRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterJoinRequest), err
}

// Delete takes name of the clusterJoinRequest and deletes it. Returns an error if one occurs.
func (c *FakeClusterJoinRequests) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterjoinrequestsResource, name), &v1beta1.ClusterJoinRequest{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterJoinRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterjoinrequestsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterJoinRequestList{})
	return err
}

// Patch applies the patch and returns the patched clusterJoinRequest.
func (c *FakeClusterJoinRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1beta1.ClusterJoinRequest, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterjoinrequestsResource, name, pt, data, subresources...), &v1beta1.ClusterJoinRequest{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterJoinRequest), err
}
