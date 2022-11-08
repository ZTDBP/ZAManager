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

type Server struct {
	gorm.Model
	ResourceID string `json:"resource_id"`
	UserUUID   string `json:"user_uuid" gorm:"user_uuid"`
	Name       string `json:"name"`
	UUID       string `json:"uuid" gorm:"column:uuid"`
	Host       string `json:"host"`
	Port       int    `json:"port"`
	OutPort    int    `json:"out_port"`
	CaPem      string `json:"ca_pem"`
	CertPem    string `json:"cert_pem"`
	KeyPem     string `json:"key_pem"`
}

func (Server) TableName() string {
	return "zta_server"
}

type ServerAttr struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	OutPort int    `json:"out_port"`
}

func (c ServerAttr) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *ServerAttr) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}
