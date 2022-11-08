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

package server

import (
	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/ZTDBP/ZAManager/pkg/mysql"
	"github.com/ZTDBP/ZAManager/pkg/redis"
	"github.com/urfave/cli"
)

func InitService(c *cli.Context) (err error) {
	if err = confer.Init(c.String("c")); err != nil {
		return
	}
	cfg := confer.GlobalConfig()
	logger.Init(&logger.Config{
		Level:       logger.LogLevel(),
		Filename:    logger.LogFile(),
		SendToFile:  logger.SendLogToFile(),
		Development: confer.ConfigEnvIsDev(),
	})
	if err = redis.Init(&cfg.Redis); err != nil {
		return
	}
	if err = mysql.Init(&cfg.Mysql); err != nil {
		return
	}
	if err = mysql.SqlMigrate(); err != nil {
		return
	}
	return
}
