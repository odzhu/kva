package checkset

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

//Checkset is main struct
type Checkset struct {
	Checks    []*Check
	Clientset *kubernetes.Clientset
	Config    *rest.Config
}

// NewСheckset is constructor for Checkset
func NewСheckset() *Checkset {
	//Default config location
	kubeconfpath := os.Getenv("HOME") + "/.kube/config"
	// creates the clientset
	c := new(Checkset)
	kubeconfig := flag.String("kubeconfig", kubeconfpath, "absolute path to the kubeconfig file")

	flag.Parse()
	// uses the current context in kubeconfig
	var err error
	c.Config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	c.Clientset, _ = kubernetes.NewForConfig(c.Config)

	return c
}

//Run method is main run dispatcher
func (c *Checkset) Run() *Checkset {

	funcs := make([](func() (*Check, error)), 0)
	funcs = append(funcs,
		c.NewCheck().DefaultSC,
		c.NewCheck().Standalone,
		c.NewCheck().Etcdopened,
		c.NewCheck().Apiinsecure,
		c.NewCheck().RbacDisabled)
	for _, f := range funcs {
		if check, err := f(); err == nil {
			c.Checks = append(c.Checks, check)
		}
	}
	return c
}

// Print Reults
func (c *Checkset) Print() {
	c.printDeviations()
}

func (c *Checkset) printDeviations() {
	//Results := runResults()
	//var Result Result
	w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, ' ', tabwriter.Debug)
	fmt.Fprintln(w, "Code\tCategory\tDescription\tResult\tResources")

	for _, Result := range c.Checks {
		fmt.Fprintf(w, "%v\t%v\t%v\t%v\t%v\n", Result.code, Result.category, Result.description, Result.Result, Result.resources)
	}

	w.Flush()
}

//NewCheck constructs new Check
func (c *Checkset) NewCheck() (check *Check) {
	check = new(Check)
	check.checkset = c
	return check
}
