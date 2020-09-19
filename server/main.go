package main

import (

	"github.com/micro/go-micro/web"
)


func main() {

		ginRouter := routers.Routers()
		httpServer := web.NewService(
			web.Name("httpprodservice"),
			web.Handler(ginRouter),
			//注册进consul服务中的端口,也是这里我们gin的server地址
			web.Address(":8000"),
		)

		httpServer.Init() //加了这句就可以使用命令行的形式去设置我们一些启动的配置
		httpServer.Run()

	}

