package main

import (
	"github.com/SobolevWladimir/miladb/core"
)

func main() {
	err := core.Start()
	if err != nil {
		panic(err)
	}
}
