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
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/ZTDBP/ZAManager/app/v1/controlplane/dao/redis"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/ZTDBP/ZAManager/app/v1/system/dao/mysql"
	"github.com/ZTDBP/ZAManager/app/v1/system/service"
	"github.com/ZTDBP/ZAManager/app/v1/user/dao/api"
	userDao "github.com/ZTDBP/ZAManager/app/v1/user/dao/mysql"
	"github.com/ZTDBP/ZAManager/app/v1/user/model/mmysql"
	"github.com/ZTDBP/ZAManager/pconst"
	oauth2Help "github.com/ZTDBP/ZAManager/pkg/oauth2"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/google"
)

func GetRedirectURL(c *gin.Context, company string) (redirectURL string, code int) {
	info, err := mysql.NewOauth2(c).GetOauth2ByCompany(company)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	config, err := service.Oauth2Config(info)
	if err != nil {
		code = pconst.CODE_API_BAD_REQUEST
		return
	}
	redirectURL, err = oauth2Help.GetOauth2RedirectURL(c, config)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func Oauth2Callback(c *gin.Context, session sessions.Session, company, oauth2Code string) {
	var user *mmysql.User
	// 查询对应的配置
	info, err := mysql.NewOauth2(c).GetOauth2ByCompany(company)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
		return
	}
	if info.ID == 0 {
		c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
		return
	}
	config, err := service.Oauth2Config(info)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
		return
	}
	switch company {
	case "github":
		githubUser, err := api.GetGithubUser(c, config, oauth2Code)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
			return
		}
		user = &mmysql.User{Email: fmt.Sprintf("%s@github.com", *githubUser.Login), AvatarUrl: *githubUser.AvatarURL}
		if err = userDao.NewUser(c).FirstOrCreateUser(user); err != nil {
			c.AbortWithError(http.StatusInternalServerError, errors.New("oauth error"))
			return
		}
	case "google":
		config.Endpoint = google.Endpoint
	case "facebook":
		config.Endpoint = facebook.Endpoint
	default:

	}
	userBytes, _ := json.Marshal(user)
	session.Set("user", userBytes)
	session.Save()
	// 判断是否有机器鉴权
	if machine := session.Get("machine"); machine != nil {
		// 删除machine这个session
		session.Delete("machine")
		session.Save()
		// 给当前请求授权
		cookie, _ := c.Cookie("zta")
		redis.NewMachine(c).PubMachineCookie(machine.(string), cookie)
		c.String(http.StatusOK, "Auth Success!")
	} else {
		c.Redirect(http.StatusSeeOther, "/")
	}
	return
}
