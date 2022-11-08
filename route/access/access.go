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

package access

import (
	v1 "github.com/ZTDBP/ZAManager/app/v1/access/controller"
	"github.com/ZTDBP/ZAManager/pkg/middle"

	"github.com/gin-gonic/gin"
)

func APIAccess(parentRoute gin.IRouter) {
	r := parentRoute.Group("access", middle.Oauth2())
	//r := parentRoute.Group("access")
	{
		resource := r.Group("resource")
		{
			resource.GET("", v1.ResourceList)
			resource.POST("", v1.AddResource)
			resource.PUT("", v1.EditResource)
			resource.DELETE("/:uuid", v1.DelResource)
		}
		relay := r.Group("relay")
		{
			relay.GET("", v1.RelayList)
			relay.POST("", v1.AddRelay)
			relay.PUT("", v1.EditRelay)
			relay.DELETE("/:uuid", v1.DelRelay)
		}
		server := r.Group("server")
		{
			server.GET("", v1.ServerList)
			server.POST("", v1.AddServer)
			server.PUT("", v1.EditServer)
			server.DELETE("/:uuid", v1.DelServer)
		}
		client := r.Group("client")
		{
			client.GET("", v1.ClientList)
			client.POST("", v1.AddClient)
			client.PUT("", v1.EditClient)
			client.DELETE("/:uuid", v1.DelClient)
		}
	}

}
