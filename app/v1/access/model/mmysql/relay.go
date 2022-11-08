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
	"gorm.io/gorm"
)

type Relay struct {
	gorm.Model
	UserUUID string `json:"user_uuid" gorm:"user_uuid"`
	Name     string `json:"name"`
	UUID     string `json:"uuid" gorm:"column:uuid"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	OutPort  int    `json:"out_port"`
	CaPem    string `json:"ca_pem"`
	CertPem  string `json:"cert_pem"`
	KeyPem   string `json:"key_pem"`
}

func (Relay) TableName() string {
	return "zta_relay"
}

type RelayAttrs struct {
	Name    string `json:"name"`
	UUID    string `json:"uuid"`
	Host    string `json:"host"`
	Port    int    `json:"port"`
	OutPort int    `json:"out_port"`
	Sort    int    `json:"sort"`
}
