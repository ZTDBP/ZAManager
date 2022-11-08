// Copyright 2022-present The ZTDBP Authors.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
