package checkset

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

//Standalone for a pods running without any supervising
func (c *Check) Standalone() (*Check, error) {
	pods, err := c.checkset.Clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	c.code = 101
	c.category = "ha"
	c.description = "Result standalone pods"

	//fmt.Printf("%+v", pods)
	var spods []string
	for _, podName := range pods.Items {
		if len(podName.ObjectMeta.OwnerReferences) < 1 {
			spods = append(spods, podName.Name)
		}
	}
	c.resources = spods
	if len(spods) > 0 {
		//fmt.Printf("Nasty pods: %v\n", spods)
		c.Result = true
	}
	c.Result = false

	return c, nil
}
