package system

import (
	v1 "github.com/ZTDBP/ZAManager/app/v1/system/controller"
	"github.com/gin-gonic/gin"
)

func APISystem(parentRoute gin.IRouter) {
	system := parentRoute.Group("system")
	{
		oauth2 := system.Group("oauth2")
		{
			oauth2.GET("", v1.ListOauth2)
			oauth2.POST("", v1.AddOauth2)
			oauth2.PUT("", v1.EditOauth2)
			oauth2.DELETE("/:id", v1.DelOauth2)
		}
	}
}
