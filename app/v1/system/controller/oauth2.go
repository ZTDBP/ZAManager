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
	"strconv"

	"github.com/ZTDBP/ZAManager/app/base/controller"
	"github.com/ZTDBP/ZAManager/app/v1/system/model/mmysql"
	"github.com/ZTDBP/ZAManager/app/v1/system/model/mparam"
	"github.com/ZTDBP/ZAManager/app/v1/system/service"
	"github.com/ZTDBP/ZAManager/pconst"
	"github.com/ZTDBP/ZAManager/pkg/response"
	"github.com/gin-gonic/gin"
)

// @Summary ListOauth2
// @Description 获取ZTA的Oauth2
// @Tags ZTA Oauth2
// @Produce  json
// @Success 200 {object} controller.Res
// @Router /sysytem/oauth2 [get]
func ListOauth2(c *gin.Context) {
	code, data := service.ListOauth2(c)
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary AddOauth2
// @Description 新增ZTA的Oauth2
// @Tags ZTA Oauth2
// @Accept  json
// @Produce  json
// @Param Oauth2 body mparam.AddOauth2 true "新增ZTA的Oauth2"
// @Success 200 {object} controller.Res
// @Router /sysytem/oauth2 [post]
func AddOauth2(c *gin.Context) {
	param := &mmysql.Oauth2{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.AddOauth2(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary EditOauth2
// @Description 修改ZTA的Oauth2
// @Tags ZTA Oauth2
// @Accept  json
// @Produce  json
// @Param Oauth2 body mparam.EditOauth2 true "修改ZTA的Oauth2"
// @Success 200 {object} controller.Res
// @Router /sysytem/oauth2 [put]
func EditOauth2(c *gin.Context) {
	param := &mparam.EditOauth2{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	code = service.EditOauth2(c, param)
	response.UtilResponseReturnJson(c, code, nil)
}

// @Summary DelOauth2
// @Description 删除ZTA的Oauth2
// @Tags ZTA Oauth2
// @Produce  json
// @Param id path int true "主键ID"
// @Success 200 {object} controller.Res
// @Router /sysytem/oauth2/{id} [delete]
func DelOauth2(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.UtilResponseReturnJsonFailed(c, pconst.CODE_COMMON_PARAMS_INCOMPLETE)
		return
	}
	code := service.DelOauth2(c, uint64(idInt))
	response.UtilResponseReturnJson(c, code, nil)
}
