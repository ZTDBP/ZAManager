package route

import (
	"github.com/ZTDBP/ZAManager/app/base/controller"
	v1 "github.com/ZTDBP/ZAManager/app/v1/controlplane/controller"
	"github.com/ZTDBP/ZAManager/pconst"
	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/ZTDBP/ZAManager/route/access"
	"github.com/ZTDBP/ZAManager/route/controlplane"
	"github.com/ZTDBP/ZAManager/route/system"
	"github.com/ZTDBP/ZAManager/route/user"

	"github.com/gin-gonic/gin"
)

// Home 主页
func Home(engine *gin.Engine) {
	engine.GET("", controller.Welcome)
}

func Api(engine *gin.Engine) {
	engine.GET("/a/:hash", v1.MachineOauth)
	prefix := confer.ConfigAppGetString("UrlPrefix", "")
	RouteV1 := engine.Group(prefix + pconst.APIAPIV1URL)
	{
		access.APIAccess(RouteV1)
		controlplane.APIControlPlane(RouteV1)
		system.APISystem(RouteV1)
		user.APIUser(RouteV1)
	}
}

func NotFound(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.String(404, "404 Not Found")
	})
}
