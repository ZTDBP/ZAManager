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

import (
	"github.com/ZTDBP/ZAManager/app/base/mdb"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
)

type ClientList struct {
	mdb.Paginate
	Name     string `json:"name" form:"name"`
	ServerID int    `json:"server_id"`
}

type AddClient struct {
	ServerID uint64              `json:"server_id" form:"server_id" binding:"required"`
	Name     string              `json:"name" form:"name" binding:"required"`
	Port     int                 `json:"port" form:"port" binding:"required"`     // 443
	Expire   int                 `json:"expire" form:"expire" binding:"required"` // 过期时间：天
	Target   mmysql.ClientTarget `json:"target" binding:"required"`
}

type EditClient struct {
	ID       uint64              `json:"id" form:"id" binding:"required"`
	ServerID uint64              `json:"server_id" form:"server_id" binding:"required"`
	Name     string              `json:"name" form:"name" binding:"required"`
	Port     int                 `json:"port" form:"port" binding:"required"`     // 443
	Expire   int                 `json:"expire" form:"expire" binding:"required"` // 过期时间：天
	Target   mmysql.ClientTarget `json:"target" binding:"required"`
}
