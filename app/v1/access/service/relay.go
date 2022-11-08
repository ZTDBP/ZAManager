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
	"time"

	"github.com/ZTDBP/ZAManager/app/v1/access/dao/api"
	"github.com/ZTDBP/ZAManager/app/v1/access/dao/mysql"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mparam"
	"github.com/ZTDBP/ZAManager/pconst"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
)

func RelayList(c *gin.Context, param mparam.RelayList) (code int, RelayList mapi.RelayList) {
	count, list, err := mysql.NewRelay(c).RelayList(param)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	RelayList.List = list
	RelayList.Paginate.Total = count
	RelayList.Paginate.PageSize = param.LimitNum
	RelayList.Paginate.Current = param.Page
	return
}

func AddRelay(c *gin.Context, param *mparam.AddRelay) (code int, data *mmysql.Relay) {
	data = &mmysql.Relay{
		Name:    param.Name,
		Host:    param.Host,
		Port:    param.Port,
		OutPort: param.OutPort,
		UUID:    uuid.NewString(),
	}
	attrs := map[string]interface{}{
		"type":     "relay",
		"name":     data.Name,
		"host":     data.Host,
		"port":     data.Port,
		"out_port": data.OutPort,
		"uuid":     data.UUID,
	}
	sentinelSign, err := api.ApplySign(c, attrs, "zero-access", "zero-access", data.Host, time.Now().AddDate(0, 0, 90))
	if err != nil {
		return pconst.CODE_COMMON_SERVER_BUSY, nil
	}
	data.CaPem = sentinelSign.CaPEM
	data.CertPem = sentinelSign.CertPEM
	data.KeyPem = sentinelSign.KeyPEM
	err = mysql.NewRelay(c).AddRelay(data)
	if err != nil {
		return pconst.CODE_COMMON_SERVER_BUSY, nil
	}
	return
}

func EditRelay(c *gin.Context, param *mparam.EditRelay) (code int) {
	info, err := mysql.NewRelay(c).GetRelayByID(param.ID)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	if info.ID == 0 {
		code = pconst.CODE_COMMON_DATA_NOT_EXIST
		return
	}
	info.Name = param.Name
	info.Host = param.Host
	info.Port = param.Port
	info.OutPort = param.OutPort
	attrs := map[string]interface{}{
		"type":     "relay",
		"name":     info.Name,
		"host":     info.Host,
		"port":     info.Port,
		"out_port": info.OutPort,
		"uuid":     info.UUID,
	}
	sentinelSign, err := api.ApplySign(c, attrs, "zero-access", "zero-access", info.Host, time.Now().AddDate(0, 0, 90))
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	info.CaPem = sentinelSign.CaPEM
	info.CertPem = sentinelSign.CertPEM
	info.KeyPem = sentinelSign.KeyPEM
	err = mysql.NewRelay(c).EditRelay(info)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
		return
	}
	return
}

func DelRelay(c *gin.Context, uuid string) (code int) {
	// TODO check clients
	err := mysql.NewRelay(c).DelRelay(uuid)
	if err != nil {
		code = pconst.CODE_COMMON_SERVER_BUSY
	}
	return
}
