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

package api

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func GetGithubUser(c *gin.Context, config *oauth2.Config, code string) (user *github.User, err error) {
	token, err := config.Exchange(c, code)
	if err != nil {
		return
	}
	client := config.Client(context.TODO(), token)
	response, err := client.Get("https://api.github.com/user")
	if err != nil {
		return
	}
	defer response.Body.Close()
	info, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(info, &user)
	if err != nil {
		return
	}
	return
}
