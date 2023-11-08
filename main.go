package main

import (
	"example/komposervice/api"
	"example/komposervice/internal/config"
	"fmt"
	"log"
)

// @title Microservice API Documentation
// @version 1.0.0
// @description This is a documentation for the Microservice API
// @host
// @basePath /
func main() {
	server := api.New()
	if err := server.Run(fmt.Sprintf("%s:%s", config.Host, config.Port)); err != nil {
		log.Fatal(err)
	}
}
