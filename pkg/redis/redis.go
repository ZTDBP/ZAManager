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

	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/go-redis/redis/v8"
)

var Client *redis.Client

func Init(cfg *confer.Redis) (err error) {
	Client = redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})
	if err = Client.Ping(context.Background()).Err(); err != nil {
		return
	}
	return
}
