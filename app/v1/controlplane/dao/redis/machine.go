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

package redis

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ZTDBP/ZAManager/pkg/logger"

	redisclient "github.com/ZTDBP/ZAManager/pkg/redis"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Machine struct {
	c      *gin.Context
	ctx    context.Context
	client *redis.Client
	Prefix string
}

func NewMachine(c *gin.Context) *Machine {
	return &Machine{
		c:      c,
		ctx:    context.Background(),
		client: redisclient.Client,
		Prefix: "zta:",
	}
}

func (m *Machine) SetLoginHash(key, hash string) (err error) {
	// 先将之前的machine取出来
	result, err := m.client.Get(m.ctx, fmt.Sprintf("%s%s", m.Prefix, key)).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return
	}
	if len(result) > 0 {
		// 删除鉴权的redis数据
		if err = m.client.Del(m.ctx, result).Err(); err != nil {
			logger.Errorf(m.c, "SetLoginHash err : %v", err)
			return
		}
	}
	if err = m.client.Set(m.ctx,
		fmt.Sprintf("%s%s", m.Prefix, key), fmt.Sprintf("%s%s", m.Prefix, hash), time.Second*60*5).Err(); err != nil {
		logger.Errorf(m.c, "SetLoginHash err : %v", err)
		return
	}
	if err = m.client.Set(m.ctx, fmt.Sprintf("%s%s", m.Prefix, hash), "", time.Second*60*5).Err(); err != nil {
		logger.Errorf(m.c, "SetLoginHash err : %v", err)
		return
	}
	return
}

func (m *Machine) GetLoginHash(hash string) (exist bool, data string, err error) {
	result, err := m.client.Get(context.Background(),
		fmt.Sprintf("%s%s", m.Prefix, hash)).Result()
	if errors.Is(err, redis.Nil) {
		return false, data, nil
	} else if err != nil {
		logger.Errorf(m.c, "GetLoginHash err : %v", err)
		return false, data, err
	} else {
		return true, result, nil
	}
}

func (m *Machine) PubMachineCookie(key, value string) (err error) {
	err = m.client.RPush(m.ctx, fmt.Sprintf("%spubsub:%s", m.Prefix, key), value).Err()
	if err != nil {
		logger.Errorf(m.c, "PubMachineCookie err : %v", err)
		return
	}
	err = m.client.Expire(m.ctx, fmt.Sprintf("%s%s", m.Prefix, key), time.Second*60*5).Err()
	if err != nil {
		logger.Errorf(m.c, "PubMachineCookie err : %v", err)
		return
	}
	return
}

func (m *Machine) SubMachineCookie(key string, timeout time.Duration) (data []string, err error) {
	data, err = m.client.BLPop(m.ctx, timeout, fmt.Sprintf("%spubsub:%s", m.Prefix, key)).Result()
	if errors.Is(err, redis.Nil) {
		return nil, nil
	}
	if err != nil {
		logger.Errorf(m.c, "SubMachineCookie err : %v", err)
		return
	}
	return
}
