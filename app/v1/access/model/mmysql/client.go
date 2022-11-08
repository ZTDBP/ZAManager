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
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	UserUUID string       `json:"user_uuid" gorm:"user_uuid"`
	Name     string       `json:"name"`
	ServerID uint64       `json:"server_id"`
	UUID     string       `json:"uuid" gorm:"column:uuid"`
	Port     int          `json:"port"`
	Expire   int          `json:"expire"` // 过期时间：天
	Relay    Relays       `json:"relay"`
	Server   ServerAttr   `json:"server"`
	Target   ClientTarget `json:"target"`
	CaPem    string       `json:"ca_pem"`
	CertPem  string       `json:"cert_pem"`
	KeyPem   string       `json:"key_pem"`
}

func (Client) TableName() string {
	return "zta_client"
}

type Relays []RelayAttrs

//type Servers ServerAttr

//type Resource []Resource

type ClientTarget struct {
	Host string `json:"host" binding:"required"`
	Port int    `json:"port" binding:"required"`
}

func (c Relays) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *Relays) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

//func (c Resource) Value() (driver.Value, error) {
//	b, err := json.Marshal(c)
//	return string(b), err
//}
//
//func (c *Resource) Scan(input interface{}) error {
//	return json.Unmarshal(input.([]byte), c)
//}

func (c ClientTarget) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ClientTarget) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

type ClientAttrs struct {
	Type   string       `json:"type"`
	Name   string       `json:"name"`
	UUID   string       `json:"uuid"`
	Port   int          `json:"port"`
	Relay  []RelayAttrs `json:"relay"`
	Server ServerAttr   `json:"server"`
	Target ClientTarget `json:"target"`
}
