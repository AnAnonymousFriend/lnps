package routers

import (
	"LNPS/server/middleware"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "LNPS/server/docs"
)

// 多路由注册
func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.Use(middleware.Cors())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	ApiGroup := Router.Group("/api/v1")
	// 注册用户路由
	InitAutoCodeRouter(ApiGroup)
	return Router

}