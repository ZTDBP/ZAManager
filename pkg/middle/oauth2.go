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

package middle

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Oauth2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		if user := session.Get("user"); user != nil {
			ctx.Set("user", user)
			ctx.Next()
		} else {
			//if confer.ConfigEnvGet() == "dev" {
			//	userBytes, _ := json.Marshal(&mmysql.User{
			//		Email:     "nisainan@github.com",
			//		AvatarUrl: "https://avatars.githubusercontent.com/u/25074107?v=4",
			//		UUID:      "3933d404-2025-4851-bfe3-1c07c5280c72",
			//	})
			//	ctx.Set("user", userBytes)
			//	ctx.Next()
			//} else {
			//	ctx.AbortWithStatus(http.StatusUnauthorized)
			//}
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
