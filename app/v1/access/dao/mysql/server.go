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

package mysql

import (
	"errors"

	"github.com/ZTDBP/ZAManager/pkg/util"

	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mparam"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/ZTDBP/ZAManager/pkg/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewServer(c *gin.Context) *Server {
	return &Server{
		DaoMysql: mysql.DaoMysql{TableName: "zta_server"},
		c:        c,
	}
}

func (p *Server) ServerList(param mparam.ServerList) (
	total int64, list []mmysql.Server, err error) {
	orm := p.GetOrm().DB
	query := orm.Table(p.TableName)
	if len(param.Name) > 0 {
		query = query.Where("name like ?", "%"+param.Name+"%")
	}
	if param.ResourceID > 0 {
		query = query.Where("find_in_set (?,resource_id)", param.ResourceID)
	}
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
	}
	err = query.Model(&list).Count(&total).Error
	if total > 0 {
		offset := param.GetOffset()
		err = query.Limit(param.LimitNum).Offset(offset).
			Order("created_at desc").
			Find(&list).Error
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "ServerList err : %v", err)
	}
	return
}

func (p *Server) GetServerByID(id uint64) (info *mmysql.Server, err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("id = ?", id)
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
	}
	err = query.First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetServerById err : %v", err)
	}
	return
}

func (p *Server) GetServerByUUID(uuid string) (info *mmysql.Server, err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("uuid = ?", uuid)
	if user := util.User(p.c); user != nil {
		query = query.Where("`user_uuid` = ?", user.UUID)
	}
	err = query.First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetServerByUUID err : %v", err)
	}
	return
}

func (p *Server) AddServer(data *mmysql.Server) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddServer err : %v", err)
	}
	return
}

func (p *Server) EditServer(data *mmysql.Server) (err error) {
	if user := util.User(p.c); user != nil {
		data.UserUUID = user.UUID
	}
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditServer err : %v", err)
	}
	return
}

func (p *Server) DelServer(uuid string) (err error) {
	orm := p.GetOrm()
	query := orm.Table(p.TableName).Where("uuid = ?", uuid)
	if user := util.User(p.c); user != nil {
		query = query.Where("user_uuid = ?", user.UUID)
	}
	err = query.Delete(&mmysql.Server{}).Error
	if err != nil {
		logger.Errorf(p.c, "DelServer err : %v", err)
	}
	return
}
