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
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email"`
	AvatarUrl string `json:"avatar_url"`
	UUID      string `json:"uuid" gorm:"column:uuid"`
}

func (User) TableName() string {
	return "zta_user"
}
