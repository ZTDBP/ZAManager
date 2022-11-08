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

package user

import (
	v1 "github.com/ZTDBP/ZAManager/app/v1/user/controller"
	"github.com/ZTDBP/ZAManager/pkg/middle"
	"github.com/gin-gonic/gin"
)

func APIUser(parentRoute gin.IRouter) {
	user := parentRoute.Group("user")
	{
		user.GET("/login/:company", v1.Login)
		user.GET("/oauth2/callback/:company", v1.Oauth2Callback)
		user.GET("/detail", middle.Oauth2(), v1.UserDetail)
	}
}
