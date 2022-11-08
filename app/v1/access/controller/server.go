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

package v1

import (
	"github.com/ZTDBP/ZAManager/app/base/controller"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mparam"
	"github.com/ZTDBP/ZAManager/app/v1/access/service"
	"github.com/ZTDBP/ZAManager/pkg/response"

	"github.com/gin-gonic/gin"
)

// @Summary ServerList
// @Description 获取ZTA的server
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/server [get]
func ServerList(c *gin.Context) {
	param := mparam.ServerList{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.ServerList(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddServer
// @Description 新增ZTA的server
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Server body mparam.AddServer true "新增ZTA的server"
// @Success 200 {object} controller.Res
// @Router /access/server [post]
func AddServer(c *gin.Context) {
	param := &mparam.AddServer{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.AddServer(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditServer
// @Description 修改ZTA的server
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Server body mparam.EditServer true "修改ZTA的server"
// @Success 200 {object} controller.Res
// @Router /access/server [put]
func EditServer(c *gin.Context) {
	param := &mparam.EditServer{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditServer(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelServer
// @Description 删除ZTA的server
// @Tags ZTA
// @Produce  json
// @Param uuid path string true "uuid"
// @Success 200 {object} controller.Res
// @Router /access/server/{uuid} [delete]
func DelServer(c *gin.Context) {
	code := service.DelServer(c, c.Param("uuid"))
	response.UtilResponseReturnJson(c, code, nil)
}
