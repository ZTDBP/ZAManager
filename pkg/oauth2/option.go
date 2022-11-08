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

package oauth2

import (
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"golang.org/x/oauth2"
)

func GetOauth2RedirectURL(c *gin.Context, config *oauth2.Config) (redirectURL string, err error) {
	state := xid.New().String()
	redirectURL = config.AuthCodeURL(state, oauth2.ApprovalForce)
	session := sessions.Default(c)
	session.Set("state", state)
	err = session.Save()
	return
}
