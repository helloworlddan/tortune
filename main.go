package main

import (
	"flag"
	"fmt"

	"github.com/helloworlddan/tortune/tortune"
)

func main() {
	versionToggle := flag.Bool("version", false, "show version info")
	flag.Parse()

	if *versionToggle == true {
		fmt.Printf("tortune version %s\n", tortune.Version)
		return
	}

	fmt.Println(tortune.HitMe())
}
