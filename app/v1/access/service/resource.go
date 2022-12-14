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
	"github.com/ZTDBP/ZAManager/app/v1/access/dao/mysql"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mparam"
	"github.com/ZTDBP/ZAManager/pconst"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func ResourceList(c *gin.Context, param mparam.ResourceList) (code int, ResourceList mapi.ResourceList) {
	count, list, err := mysql.NewResource(c).ResourceList(param)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	ResourceList.List = list
	ResourceList.Paginate.Total = count
	ResourceList.Paginate.PageSize = param.LimitNum
	ResourceList.Paginate.Current = param.Page
	return
}

func AddResource(c *gin.Context, param *mparam.AddResource) (code int, data *mmysql.Resource) {
	data = &mmysql.Resource{
		Name: param.Name,
		UUID: uuid.NewString(),
		Type: param.Type,
		Host: param.Host,
		Port: param.Port,
	}
	err := mysql.NewResource(c).AddResource(data)
	if err != nil {
		return pconst.CODE_COMMON_SERVER_BUSY, nil
	}
	return
}

func EditResource(c *gin.Context, param *mparam.EditResource) (code int) {
	info, err := mysql.NewResource(c).GetResourceByID(param.ID)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_COMMON_DATA_NOT_EXIST
		return
	}
	info.Name = param.Name
	info.Type = param.Type
	info.Host = param.Host
	info.Port = param.Port
	err = mysql.NewResource(c).EditResource(info)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func DelResource(c *gin.Context, uuid string) (code int) {
	// check if any servers under this resource
	resource, err := mysql.NewResource(c).GetResourceByUUID(uuid)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if resource == nil || resource.ID == 0 {
		code = pconst.CODE_COMMON_DATA_NOT_EXIST
		return
	}
	total, _, err := mysql.NewServer(c).ServerList(mparam.ServerList{
		ResourceID: int(resource.ID),
	})
	if total > 0 {
		code = pconst.CODE_DATA_HAS_RELATION
		c.Set("error", "There are servers under this resource")
		return
	}
	err = mysql.NewResource(c).DelResource(uuid)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
	}
	return
}
