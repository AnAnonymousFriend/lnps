package routers

import (
	"LNPS/server/api"
	jwt "LNPS/server/middleware"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	
	AutoCodeRouter := Router.Group("autoCode").
		Use(jwt.JWT())
	{
		AutoCodeRouter.POST("GetAuth", api.GetAuth)
	}
}

