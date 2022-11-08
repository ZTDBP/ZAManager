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

// @Summary RelayList
// @Description 获取ZTA的relay
// @Tags ZTA
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /access/relay [get]
func RelayList(c *gin.Context) {
	param := mparam.RelayList{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.RelayList(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddRelay
// @Description 新增ZTA的relay
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Relay body mparam.AddRelay true "新增ZTA的relay"
// @Success 200 {object} controller.Res
// @Router /access/relay [post]
func AddRelay(c *gin.Context) {
	param := &mparam.AddRelay{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code, data := service.AddRelay(c, param)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary EditRelay
// @Description 修改ZTA的relay
// @Tags ZTA
// @Accept  json
// @Produce  json
// @Param Relay body mparam.EditRelay true "修改ZTA的relay"
// @Success 200 {object} controller.Res
// @Router /access/relay [put]
func EditRelay(c *gin.Context) {
	param := &mparam.EditRelay{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditRelay(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelRelay
// @Description 删除ZTA的relay
// @Tags ZTA
// @Produce  json
// @Param uuid path string true "uuid"
// @Success 200 {object} controller.Res
// @Router /access/relay/{uuid} [delete]
func DelRelay(c *gin.Context) {
	code := service.DelRelay(c, c.Param("uuid"))
	response.UtilResponseReturnJson(c, code, nil)
}
