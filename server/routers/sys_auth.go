package routers

import (
	"LNPS/server/routers/api"
	"github.com/gin-gonic/gin"
	jwt "LNPS/server/middleware"


)

func InitAutoCodeRouter(Router *gin.RouterGroup) {
	
	AutoCodeRouter := Router.Group("autoCode").
		Use(jwt.JWT())
	{
		AutoCodeRouter.POST("GetAuth", api.GetAuth)
	}
}

