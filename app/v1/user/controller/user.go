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
	"errors"
	"fmt"
	"net/http"

	"github.com/ZTDBP/ZAManager/app/v1/user/service"
	"github.com/ZTDBP/ZAManager/pconst"
	"github.com/ZTDBP/ZAManager/pkg/response"
	"github.com/ZTDBP/ZAManager/pkg/util"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	redirectURL, code := service.GetRedirectURL(c, c.Param("company"))
	if code != pconst.CODE_ERROR_OK {
		// TODO Redirect to BadRequest page
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("company %s not support", c.Param("company")))
		return
	}
	c.Redirect(http.StatusSeeOther, redirectURL)
}

func UserDetail(c *gin.Context) {
	response.UtilResponseReturnJson(c, pconst.CODE_ERROR_OK, util.User(c))
}

func Oauth2Callback(c *gin.Context) {
	session := sessions.Default(c)
	state := session.Get("state")
	if state != c.Query("state") {
		_ = c.AbortWithError(http.StatusUnauthorized, errors.New("state error"))
		return
	}
	if len(c.Query("code")) == 0 {
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("code error"))
	}
	service.Oauth2Callback(c, session, c.Param("company"), c.Query("code"))
}
