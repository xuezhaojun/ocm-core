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
	workv1 "open-cluster-management.io/api/work/v1"
)

// FakeManifestWorks implements ManifestWorkInterface
type FakeManifestWorks struct {
	Fake *FakeWorkV1
	ns   string
}

var manifestworksResource = schema.GroupVersionResource{Group: "work.open-cluster-management.io", Version: "v1", Resource: "manifestworks"}

var manifestworksKind = schema.GroupVersionKind{Group: "work.open-cluster-management.io", Version: "v1", Kind: "ManifestWork"}

// Get takes name of the manifestWork, and returns the corresponding manifestWork object, and an error if there is any.
func (c *FakeManifestWorks) Get(ctx context.Context, name string, options v1.GetOptions) (result *workv1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(manifestworksResource, c.ns, name), &workv1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workv1.ManifestWork), err
}

// List takes label and field selectors, and returns the list of ManifestWorks that match those selectors.
func (c *FakeManifestWorks) List(ctx context.Context, opts v1.ListOptions) (result *workv1.ManifestWorkList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(manifestworksResource, manifestworksKind, c.ns, opts), &workv1.ManifestWorkList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &workv1.ManifestWorkList{ListMeta: obj.(*workv1.ManifestWorkList).ListMeta}
	for _, item := range obj.(*workv1.ManifestWorkList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested manifestWorks.
func (c *FakeManifestWorks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(manifestworksResource, c.ns, opts))

}

// Create takes the representation of a manifestWork and creates it.  Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *FakeManifestWorks) Create(ctx context.Context, manifestWork *workv1.ManifestWork, opts v1.CreateOptions) (result *workv1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(manifestworksResource, c.ns, manifestWork), &workv1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workv1.ManifestWork), err
}

// Update takes the representation of a manifestWork and updates it. Returns the server's representation of the manifestWork, and an error, if there is any.
func (c *FakeManifestWorks) Update(ctx context.Context, manifestWork *workv1.ManifestWork, opts v1.UpdateOptions) (result *workv1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(manifestworksResource, c.ns, manifestWork), &workv1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workv1.ManifestWork), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeManifestWorks) UpdateStatus(ctx context.Context, manifestWork *workv1.ManifestWork, opts v1.UpdateOptions) (*workv1.ManifestWork, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(manifestworksResource, "status", c.ns, manifestWork), &workv1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workv1.ManifestWork), err
}

// Delete takes name of the manifestWork and deletes it. Returns an error if one occurs.
func (c *FakeManifestWorks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(manifestworksResource, c.ns, name, opts), &workv1.ManifestWork{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeManifestWorks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(manifestworksResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &workv1.ManifestWorkList{})
	return err
}

// Patch applies the patch and returns the patched manifestWork.
func (c *FakeManifestWorks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *workv1.ManifestWork, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(manifestworksResource, c.ns, name, pt, data, subresources...), &workv1.ManifestWork{})

	if obj == nil {
		return nil, err
	}
	return obj.(*workv1.ManifestWork), err
}
