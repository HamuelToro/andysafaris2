package main

import (
	"log"

	"github.com/EmilioCliff/car-rental/api"
)

func main() {
	server := api.NewServer()
	log.Println("starting server at local host: ", "8080")
	if err := server.Start("0.0.0.0:8080"); err != nil {
		log.Println(err)
	}
}