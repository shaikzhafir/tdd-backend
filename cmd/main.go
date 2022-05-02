package main

import (
	"log"
	"tdd-backend"
)

func main() {
	config := tdd_backend.NewConfig()
	server := tdd_backend.New(config)

	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
