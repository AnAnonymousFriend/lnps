package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 多路由注册
func Routers() *gin.Engine {
	var Router = gin.Default()
	//Router.Use(middleware.Cors())
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("/api/v1")
	// 注册用户路由
	InitAutoCodeRouter(ApiGroup)
	return Router

}