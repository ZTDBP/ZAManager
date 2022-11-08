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

package server

import (
	"fmt"
	"strconv"

	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/ZTDBP/ZAManager/pkg/gin"
	"github.com/ZTDBP/ZAManager/pkg/middle"
	"github.com/ZTDBP/ZAManager/route"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RunHTTP() {
	engine := gin.NewGin()
	//engine.Use(middle.RequestID())
	// 仅针对开发环境生效的组件
	//if confer.ConfigEnvIsDev() {
	// 跨域中间件
	engine.Use(middle.CorsV2())
	// swagger
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
	engine.Use(middle.Session("zta", &confer.GlobalConfig().Redis))
	route.Home(engine)
	route.Api(engine)
	route.NotFound(engine)
	httpPort := confer.ConfigAppGetInt("port", 80)
	portStr := ":" + strconv.Itoa(httpPort)
	fmt.Println("start", httpPort)
	gin.ListenHTTP(portStr, engine, 10)
}
