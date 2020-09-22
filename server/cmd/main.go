package main

import (
	"LNPS/server/routers"
	"github.com/micro/go-micro/web"

)



func main() {

	//client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))

		ginRouter := routers.Routers()
		httpServer := web.NewService(
			web.Name("httpprodservice"),
			web.Handler(ginRouter),
			web.Address(":8000"),
		)

		httpServer.Init()
		httpServer.Run()

	}




