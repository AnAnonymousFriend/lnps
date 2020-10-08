package main

import (
	"LNPS/server/pkg/setting"
	"LNPS/server/routers"
	"github.com/micro/go-micro/web"
)

func main() {
		setting.Setup()
		print(setting.RedisSetting.MaxIdle)
		print(setting.MongoDBSetting.Host)

		ginRouter := routers.Routers()
		httpServer := web.NewService(
			web.Name("httpprodservice"),
			web.Handler(ginRouter),
			web.Address(":8000"),
		)
		httpServer.Init()
		httpServer.Run()

		}




