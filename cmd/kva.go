package main

import checkset "github.com/odzhu/kva/checkset"

func main() {

	cs := checkset.NewСheckset()
	//cs.NewCheck().Apiinsecure()
	cs.Run()
	cs.Print()
}
