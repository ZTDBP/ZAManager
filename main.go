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

package main

import (
	"os"

	_ "github.com/ZTDBP/ZAManager/docs"
	"github.com/ZTDBP/ZAManager/server"
	"github.com/urfave/cli"
	_ "go.uber.org/automaxprocs"
)

// @title ZAManager API
// @version 1.0.0
// @description This is ZAManager api list.
// @host 127.0.0.1:80
// @BasePath /api/v1
func main() {
	app := cli.NewApp()
	app.Name = "ZAManager"
	app.Author = "TS"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "server",
			Value: "http",
			Usage: "run server type: http",
		},
		cli.StringFlag{
			Name:  "c",
			Value: "config.yaml",
			Usage: "config file url",
		},
	}
	app.Before = server.InitService
	app.Action = func(c *cli.Context) error {
		println("RunHttp Server.")
		serverType := c.String("server")
		switch serverType {
		case "http":
			server.RunHTTP()
		default:
			server.RunHTTP()
		}
		return nil
	}
	err := app.Run(os.Args)
	if err != nil {
		panic("app run error:" + err.Error())
	}
}
