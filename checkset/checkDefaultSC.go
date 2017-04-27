package checkset

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//DefaultSC checks storage class availabiltity
func (c *Check) DefaultSC() (*Check, error) {
	c.code = 501
	c.category = "functional"
	c.description = "Default storage class missing"

	if storageclass, err := c.checkset.Clientset.StorageV1beta1Client.StorageClasses().List(metav1.ListOptions{}); err == nil {
		if len(storageclass.Items) == 0 {
			c.Result = true
		}
		for _, sc := range storageclass.Items {
			c.resources = append(c.resources, sc.Name)
		}
		for _, sc := range storageclass.Items {
			if sc.ObjectMeta.GetAnnotations()["storageclass.beta.kubernetes.io/is-default-class"] != "true" {
				c.Result = true
			}
		}

	}
	return c, nil
}
