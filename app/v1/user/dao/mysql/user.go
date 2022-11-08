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
	"github.com/ZTDBP/ZAManager/app/v1/user/model/mmysql"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/ZTDBP/ZAManager/pkg/mysql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	c *gin.Context
	mysql.DaoMysql
}

func NewUser(c *gin.Context) *User {
	return &User{
		DaoMysql: mysql.DaoMysql{TableName: "zta_user"},
		c:        c,
	}
}

func (p *User) FirstOrCreateUser(data *mmysql.User) (err error) {
	orm := p.GetOrm()
	err = orm.Where(mmysql.User{Email: data.Email}).Attrs(mmysql.User{UUID: uuid.NewString(), AvatarUrl: data.AvatarUrl}).FirstOrCreate(&data).Error
	if err != nil {
		logger.Errorf(p.c, "FirstOrCreateUser err : %v", err)
	}
	return
}
