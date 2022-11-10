package main

import (

	"apidemo/handler"
	pb "apidemo/proto"

	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"

)

var (
	service = "apidemo"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterApidemoHandler(srv.Server(), new(handler.Apidemo)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
