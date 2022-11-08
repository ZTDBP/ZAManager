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
