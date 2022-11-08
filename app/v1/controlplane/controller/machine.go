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

package controller

import (
	"github.com/ZTDBP/ZAManager/app/base/controller"
	"github.com/ZTDBP/ZAManager/app/v1/controlplane/model/mparam"
	"github.com/ZTDBP/ZAManager/app/v1/controlplane/service"
	"github.com/ZTDBP/ZAManager/pkg/response"
	"github.com/gin-gonic/gin"
)

// @Summary LoginUrl
// @Description 根据机器码获取客户端鉴权的url
// @Tags ZTA ControlPlane
// @Produce  json
// @Param machine_id path string true "machine_id"
// @Success 200 {object} controller.Res
// @Router /controlplane/machine/{machine_id} [get]
func LoginUrl(c *gin.Context) {
	code, data := service.GetLoginUrl(c, c.Param("machine_id"))
	response.UtilResponseReturnJson(c, code, data)
}

// @Summary MachineOauth
// @Description 机器鉴权
// @Tags ZTA ControlPlane
// @Produce  json
// @Param hash path string true "hash"
// @Success 200 {object} controller.Res
// @Router /a/{hash} [get]
func MachineOauth(c *gin.Context) {
	service.MachineOauth(c, c.Param("hash"))
}

// @Summary MachineLongPoll
// @Description 机器鉴权
// @Tags ZTA ControlPlane
// @Produce  json
// @Param category query string true "轮询的主题"
// @Param timeout query int true "超时时间，单位：秒"
// @Success 200 {object} controller.Res
// @Router /machine/auth/poll [get]
func MachineLongPoll(c *gin.Context) {
	param := mparam.MachineLongPoll{}
	b, code := controller.BindParams(c, &param)
	if !b {
		response.UtilResponseReturnJsonFailed(c, code)
		return
	}
	data, code := service.MachineLongPoll(c, param)
	response.UtilResponseReturnJson(c, code, data)
}
