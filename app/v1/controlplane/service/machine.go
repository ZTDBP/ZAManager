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

package service

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/ZTDBP/ZAManager/app/v1/controlplane/model/mparam"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/ZTDBP/ZAManager/app/v1/controlplane/dao/redis"

	"github.com/ZTDBP/ZAManager/pkg/util"

	"github.com/ZTDBP/ZAManager/pconst"
	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/gin-gonic/gin"
)

func GetLoginUrl(c *gin.Context, machine string) (code int, loginURL string) {
	// 通过machineID和当前时间戳，计算出唯一的hash，作为登陆的地址path
	hash := util.NewMd5(fmt.Sprintf("%s%d", machine, time.Now().UnixNano()))
	// hash 放入redis缓存
	err := redis.NewMachine(c).SetLoginHash(machine, hash)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	domain := confer.ConfigAppGetString("domain", "")
	if len(os.Getenv(domain)) > 0 {
		domain = os.Getenv(domain)
	}
	loginURL = fmt.Sprintf("%s/a/%s", domain, hash)
	return
}

func MachineOauth(c *gin.Context, hash string) {
	// 判断当前hash是否存在或者是否在有消息内
	exist, _, err := redis.NewMachine(c).GetLoginHash(hash)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
		return
	}
	if exist {
		session := sessions.Default(c)
		session.Set("machine", hash)
		session.Save()
		c.Redirect(http.StatusSeeOther, "/api/v1/user/login/github")
	} else {
		// TODO 重定向到404页面
		c.String(http.StatusNotFound, "auth key not exist or expired")
		return
	}

}

func MachineLongPoll(c *gin.Context, param mparam.MachineLongPoll) (data string, code int) {
	result, err := redis.NewMachine(c).SubMachineCookie(param.Category, time.Second*time.Duration(param.Timeout))
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if len(result) >= 2 {
		data = result[1]
		return
	}
	code = pconst.CODE_COMMON_DATA_NOT_EXIST
	return
}
