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

import "github.com/ZTDBP/ZAManager/app/base/mdb"

type RelayList struct {
	mdb.Paginate
	Name string `json:"name" form:"name"`
}

type AddRelay struct {
	Name    string `json:"name" form:"name" binding:"required"`
	Host    string `json:"host" form:"host" binding:"required"`         // api.github.com
	Port    int    `json:"port" form:"port" binding:"required"`         // 443
	OutPort int    `json:"out_port" form:"out_port" binding:"required"` // 443
}

type EditRelay struct {
	ID      uint64 `json:"id" form:"id" binding:"required"`
	Name    string `json:"name" form:"name" binding:"required"`
	Host    string `json:"host" form:"host" binding:"required"`         // api.github.com
	Port    int    `json:"port" form:"port" binding:"required"`         // 443
	OutPort int    `json:"out_port" form:"out_port" binding:"required"` // 443
}
