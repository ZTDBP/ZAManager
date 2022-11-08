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

package mmysql

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Oauth2 struct {
	gorm.Model
	Company      string `json:"name" binding:"required,oneof= github facebook google"`
	ClientId     string `json:"client_id" binding:"required"`
	ClientSecret string `json:"client_secret" binding:"required"`
	RedirectUrl  string `json:"redirect_url" binding:"required"`
	Scopes       Scopes `json:"scopes" binding:"required"`
	AuthUrl      string `json:"auth_url" binding:"required"`
	TokenUrl     string `json:"token_url" binding:"required"`
}

func (Oauth2) TableName() string {
	return "zta_oauth2"
}

type Scopes []string

func (c Scopes) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Scopes) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
