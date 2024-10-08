// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"context"

	v1 "github.com/openshift/api/operator/v1"
	operatorv1 "github.com/openshift/client-go/operator/applyconfigurations/operator/v1"
	scheme "github.com/openshift/client-go/operator/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// KubeControllerManagersGetter has a method to return a KubeControllerManagerInterface.
// A group's client should implement this interface.
type KubeControllerManagersGetter interface {
	KubeControllerManagers() KubeControllerManagerInterface
}

// KubeControllerManagerInterface has methods to work with KubeControllerManager resources.
type KubeControllerManagerInterface interface {
	Create(ctx context.Context, kubeControllerManager *v1.KubeControllerManager, opts metav1.CreateOptions) (*v1.KubeControllerManager, error)
	Update(ctx context.Context, kubeControllerManager *v1.KubeControllerManager, opts metav1.UpdateOptions) (*v1.KubeControllerManager, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, kubeControllerManager *v1.KubeControllerManager, opts metav1.UpdateOptions) (*v1.KubeControllerManager, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.KubeControllerManager, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.KubeControllerManagerList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.KubeControllerManager, err error)
	Apply(ctx context.Context, kubeControllerManager *operatorv1.KubeControllerManagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeControllerManager, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, kubeControllerManager *operatorv1.KubeControllerManagerApplyConfiguration, opts metav1.ApplyOptions) (result *v1.KubeControllerManager, err error)
	KubeControllerManagerExpansion
}

// kubeControllerManagers implements KubeControllerManagerInterface
type kubeControllerManagers struct {
	*gentype.ClientWithListAndApply[*v1.KubeControllerManager, *v1.KubeControllerManagerList, *operatorv1.KubeControllerManagerApplyConfiguration]
}

// newKubeControllerManagers returns a KubeControllerManagers
func newKubeControllerManagers(c *OperatorV1Client) *kubeControllerManagers {
	return &kubeControllerManagers{
		gentype.NewClientWithListAndApply[*v1.KubeControllerManager, *v1.KubeControllerManagerList, *operatorv1.KubeControllerManagerApplyConfiguration](
			"kubecontrollermanagers",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *v1.KubeControllerManager { return &v1.KubeControllerManager{} },
			func() *v1.KubeControllerManagerList { return &v1.KubeControllerManagerList{} }),
	}
}
