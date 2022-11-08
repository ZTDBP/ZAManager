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

package mparam

import "github.com/ZTDBP/ZAManager/app/v1/system/model/mmysql"

type EditOauth2 struct {
	ID           int64         `json:"id" binding:"required"`
	Company      string        `json:"name" binding:"required,oneof= github facebook google"`
	ClientId     string        `json:"client_id" binding:"required"`
	ClientSecret string        `json:"client_secret" binding:"required"`
	RedirectUrl  string        `json:"redirect_url" binding:"required"`
	Scopes       mmysql.Scopes `json:"scopes" binding:"required"`
	AuthUrl      string        `json:"auth_url" binding:"required"`
	TokenUrl     string        `json:"token_url" binding:"required"`
}

type AddOauth2 struct {
	Company      string        `json:"name" binding:"required,oneof= github facebook google"`
	ClientId     string        `json:"client_id" binding:"required"`
	ClientSecret string        `json:"client_secret" binding:"required"`
	RedirectUrl  string        `json:"redirect_url" binding:"required"`
	Scopes       mmysql.Scopes `json:"scopes" binding:"required"`
	AuthUrl      string        `json:"auth_url" binding:"required"`
	TokenUrl     string        `json:"token_url" binding:"required"`
}
