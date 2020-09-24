package routers

import (
	"LNPS/server/api"
	"github.com/gin-gonic/gin"
)

func InitAutoCodeRouter(Router *gin.RouterGroup) {


	AutoCodeRouter := Router.Group("autoCode").
		Use()
	{
		AutoCodeRouter.POST("GetAuth", api.GetAuth)
	}
}

