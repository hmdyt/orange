package main

import (
	"log"

	"github.com/hmdyt/orange/infrastructure"
)

func main() {
	if err := infrastructure.RunServer(); err != nil {
		log.Fatal(err)
	}
}
