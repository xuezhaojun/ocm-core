// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	operatorv1 "open-cluster-management.io/api/operator/v1"
)

// FakeClusterManagers implements ClusterManagerInterface
type FakeClusterManagers struct {
	Fake *FakeOperatorV1
}

var clustermanagersResource = schema.GroupVersionResource{Group: "operator.open-cluster-management.io", Version: "v1", Resource: "clustermanagers"}

var clustermanagersKind = schema.GroupVersionKind{Group: "operator.open-cluster-management.io", Version: "v1", Kind: "ClusterManager"}

// Get takes name of the clusterManager, and returns the corresponding clusterManager object, and an error if there is any.
func (c *FakeClusterManagers) Get(ctx context.Context, name string, options v1.GetOptions) (result *operatorv1.ClusterManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clustermanagersResource, name), &operatorv1.ClusterManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ClusterManager), err
}

// List takes label and field selectors, and returns the list of ClusterManagers that match those selectors.
func (c *FakeClusterManagers) List(ctx context.Context, opts v1.ListOptions) (result *operatorv1.ClusterManagerList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clustermanagersResource, clustermanagersKind, opts), &operatorv1.ClusterManagerList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &operatorv1.ClusterManagerList{ListMeta: obj.(*operatorv1.ClusterManagerList).ListMeta}
	for _, item := range obj.(*operatorv1.ClusterManagerList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterManagers.
func (c *FakeClusterManagers) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clustermanagersResource, opts))
}

// Create takes the representation of a clusterManager and creates it.  Returns the server's representation of the clusterManager, and an error, if there is any.
func (c *FakeClusterManagers) Create(ctx context.Context, clusterManager *operatorv1.ClusterManager, opts v1.CreateOptions) (result *operatorv1.ClusterManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clustermanagersResource, clusterManager), &operatorv1.ClusterManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ClusterManager), err
}

// Update takes the representation of a clusterManager and updates it. Returns the server's representation of the clusterManager, and an error, if there is any.
func (c *FakeClusterManagers) Update(ctx context.Context, clusterManager *operatorv1.ClusterManager, opts v1.UpdateOptions) (result *operatorv1.ClusterManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clustermanagersResource, clusterManager), &operatorv1.ClusterManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ClusterManager), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterManagers) UpdateStatus(ctx context.Context, clusterManager *operatorv1.ClusterManager, opts v1.UpdateOptions) (*operatorv1.ClusterManager, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clustermanagersResource, "status", clusterManager), &operatorv1.ClusterManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ClusterManager), err
}

// Delete takes name of the clusterManager and deletes it. Returns an error if one occurs.
func (c *FakeClusterManagers) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustermanagersResource, name, opts), &operatorv1.ClusterManager{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterManagers) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clustermanagersResource, listOpts)

	_, err := c.Fake.Invokes(action, &operatorv1.ClusterManagerList{})
	return err
}

// Patch applies the patch and returns the patched clusterManager.
func (c *FakeClusterManagers) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *operatorv1.ClusterManager, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clustermanagersResource, name, pt, data, subresources...), &operatorv1.ClusterManager{})
	if obj == nil {
		return nil, err
	}
	return obj.(*operatorv1.ClusterManager), err
}
