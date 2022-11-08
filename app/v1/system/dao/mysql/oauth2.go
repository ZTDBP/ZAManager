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

	"github.com/ZTDBP/ZAManager/app/v1/system/model/mmysql"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/ZTDBP/ZAManager/pkg/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Oauth2 struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewOauth2(c *gin.Context) *Oauth2 {
	return &Oauth2{
		DaoMysql: mysql.DaoMysql{TableName: "zta_oauth2"},
		c:        c,
	}
}

func (p *Oauth2) ListOauth2() (list []mmysql.Oauth2, err error) {
	orm := p.GetOrm().DB
	err = orm.Table(p.TableName).Find(&list).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "ListOauth2 err : %v", err)
	}
	return
}

func (p *Oauth2) GetOauth2ByID(id int64) (info *mmysql.Oauth2, err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ?", id).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetOauth2ById err : %v", err)
	}
	return
}

func (p *Oauth2) GetOauth2ByCompany(company string) (info *mmysql.Oauth2, err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("company = ?", company).First(&info).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	if err != nil {
		logger.Errorf(p.c, "GetOauth2ByCompany err : %v", err)
	}
	return
}

func (p *Oauth2) AddOauth2(data *mmysql.Oauth2) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Create(&data).Error
	if err != nil {
		logger.Errorf(p.c, "AddOauth2 err : %v", err)
	}
	return
}

func (p *Oauth2) EditOauth2(data *mmysql.Oauth2) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Save(&data).Error
	if err != nil {
		logger.Errorf(p.c, "EditOauth2 err : %v", err)
	}
	return
}

func (p *Oauth2) DelOauth2(id uint64) (err error) {
	orm := p.GetOrm()
	err = orm.Table(p.TableName).Where("id = ?", id).Delete(&mmysql.Oauth2{}).Unscoped().Error
	if err != nil {
		logger.Errorf(p.c, "DelOauth2 err : %v", err)
	}
	return
}
