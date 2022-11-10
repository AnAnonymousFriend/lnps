package main

import (
	"go-micro.dev/v4"
	"log"
)

var (
	serverName    = "lnps"
	serverVersion = "1.0"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(serverName),
		micro.Version(serverVersion),
	)
	srv.Init()

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
