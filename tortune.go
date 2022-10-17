package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

const version = "0.0.2"

var (
	db = [...]string{
		"Zero-Trust means everybody gets access to Ring-0",
		"Localhost?! More like loserhost!",
		"alias yolo=\"sudo\"",
		"Friends don't let friends compile locally.",
	}
)

func main() {
	versionToggle := flag.Bool("version", false, "show version info")
	flag.Parse()

	if *versionToggle == true {
		fmt.Printf("tortune version %s\n", version)
		return
	}

	rand.Seed(time.Now().Unix())
	fmt.Println(db[rand.Intn(len(db))])
}
