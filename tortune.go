package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	db = [...]string{
		"Zero-Trust means everybody gets access to Ring-0",
		"Localhost?! More like loserhost!",
		"alias yolo=\"sudo\"",
	}
)

func main() {
	rand.Seed(time.Now().Unix())

	fmt.Println(db[rand.Intn(len(db))])
}
