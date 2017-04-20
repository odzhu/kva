package checkset

import (
	"flag"
	"fmt"
	"os"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

//Checkset is main struct
type Checkset struct {
	Checks    []Check
	clientset *kubernetes.Clientset
}

// NewСheckset is constructor for Checkset
func NewСheckset() *Checkset {
	//Default config location
	kubeconfpath := os.Getenv("HOME") + "/.kube/config"

	kubeconfig := flag.String("kubeconfig", kubeconfpath, "absolute path to the kubeconfig file")

	flag.Parse()
	// uses the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	var c Checkset
	c.clientset, _ = kubernetes.NewForConfig(config)

	return &c
}

//Pods method which returns
func (c *Checkset) Pods() (check *Check) {
	check = newCheck()
	check.checkset = c
	return
}

// Print Reults
func (c *Checkset) Print() {
	c.printDeviations()
}

func (c *Checkset) printDeviations() {
	//Results := runResults()
	//var Result Result
	for _, Result := range c.Checks {
		fmt.Printf("Code: %v Category: %v Description: %v Result: %v Resources: %v \n", Result.code, Result.category, Result.description, Result.Result, Result.resources)
	}
}

// Check struct
type Check struct {
	checkset    *Checkset
	code        int
	category    string
	description string
	Result      bool
	resources   []string
}

func newCheck() (c *Check) {
	c = new(Check)
	return c
}

//Standalone for a pods running without any supervising
func (c *Check) Standalone() (Check, error) {
	pods, err := c.checkset.clientset.CoreV1().Pods("default").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	//fmt.Printf("%+v", pods)
	var spods []string
	for _, podName := range pods.Items {
		if len(podName.ObjectMeta.OwnerReferences) < 1 {
			spods = append(spods, podName.Name)
		}
	}
	if len(spods) > 0 {
		//fmt.Printf("Nasty pods: %v\n", spods)
		return Check{
			checkset:    c.checkset,
			code:        101,
			category:    "pods",
			description: "Result standalone pods",
			Result:      true,
			resources:   spods,
		}, nil
	}
	return Check{
		checkset:    c.checkset,
		code:        101,
		category:    "pods",
		description: "Result standalone pods",
		Result:      false,
		resources:   spods,
	}, nil
}
