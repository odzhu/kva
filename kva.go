package main

import (
	"fmt"

	checkset "github.com/odzhu/kva/checkset"
)

func main() {

	cs := checkset.NewСheckset()
	output, _ := cs.Pods().Standalone()
	fmt.Printf("%+v", output)

}
