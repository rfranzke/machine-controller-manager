package v1alpha1

import (
	v1alpha1 "github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	scheme "github.com/gardener/machine-controller-manager/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// GCPMachineClassesGetter has a method to return a GCPMachineClassInterface.
// A group's client should implement this interface.
type GCPMachineClassesGetter interface {
	GCPMachineClasses(namespace string) GCPMachineClassInterface
}

// GCPMachineClassInterface has methods to work with GCPMachineClass resources.
type GCPMachineClassInterface interface {
	Create(*v1alpha1.GCPMachineClass) (*v1alpha1.GCPMachineClass, error)
	Update(*v1alpha1.GCPMachineClass) (*v1alpha1.GCPMachineClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.GCPMachineClass, error)
	List(opts v1.ListOptions) (*v1alpha1.GCPMachineClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPMachineClass, err error)
	GCPMachineClassExpansion
}

// gCPMachineClasses implements GCPMachineClassInterface
type gCPMachineClasses struct {
	client rest.Interface
	ns     string
}

// newGCPMachineClasses returns a GCPMachineClasses
func newGCPMachineClasses(c *MachineV1alpha1Client, namespace string) *gCPMachineClasses {
	return &gCPMachineClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the gCPMachineClass, and returns the corresponding gCPMachineClass object, and an error if there is any.
func (c *gCPMachineClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.GCPMachineClass, err error) {
	result = &v1alpha1.GCPMachineClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of GCPMachineClasses that match those selectors.
func (c *gCPMachineClasses) List(opts v1.ListOptions) (result *v1alpha1.GCPMachineClassList, err error) {
	result = &v1alpha1.GCPMachineClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested gCPMachineClasses.
func (c *gCPMachineClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a gCPMachineClass and creates it.  Returns the server's representation of the gCPMachineClass, and an error, if there is any.
func (c *gCPMachineClasses) Create(gCPMachineClass *v1alpha1.GCPMachineClass) (result *v1alpha1.GCPMachineClass, err error) {
	result = &v1alpha1.GCPMachineClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		Body(gCPMachineClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a gCPMachineClass and updates it. Returns the server's representation of the gCPMachineClass, and an error, if there is any.
func (c *gCPMachineClasses) Update(gCPMachineClass *v1alpha1.GCPMachineClass) (result *v1alpha1.GCPMachineClass, err error) {
	result = &v1alpha1.GCPMachineClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		Name(gCPMachineClass.Name).
		Body(gCPMachineClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the gCPMachineClass and deletes it. Returns an error if one occurs.
func (c *gCPMachineClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *gCPMachineClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched gCPMachineClass.
func (c *gCPMachineClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.GCPMachineClass, err error) {
	result = &v1alpha1.GCPMachineClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("gcpmachineclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
