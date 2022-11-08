package controlplane

import (
	v1 "github.com/ZTDBP/ZAManager/app/v1/controlplane/controller"
	"github.com/gin-gonic/gin"
)

func APIControlPlane(parentRoute gin.IRouter) {
	controlplane := parentRoute.Group("controlplane")
	{
		controlplane.GET("/machine/:machine_id", v1.LoginUrl)
		controlplane.GET("/machine/auth/poll", v1.MachineLongPoll)
	}
}
