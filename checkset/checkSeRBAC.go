package checkset

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//Rbac not enabled security check
func (c *Check) RbacDisabled() (*Check, error) {
	croles, err := c.checkset.Clientset.RbacV1beta1().ClusterRoles().List(metav1.ListOptions{})
	if err != nil {
		c.err = err.Error()
		// panic(err.Error())
	}
	c.code = 203
	c.category = "security"
	c.description = "RBAC not enabled"

	var items []string
	for _, item := range croles.Items {
		if len(item.ObjectMeta.OwnerReferences) < 1 {
			items = append(items, item.Name)
		}
	}
	c.resources = items

	if len(items) > 0 {
		c.Result = false
	} else {
		c.Result = true
	}

	return c, nil
}
