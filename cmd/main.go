package main

import (
	"log"
	"shecare/internals/di"
)

func main() {
	err := di.InitializeDependency()
	if err != nil {
		log.Fatal(err)
	}
}
