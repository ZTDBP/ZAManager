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

package middle

import (
	"github.com/ZTDBP/ZAManager/pkg/confer"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Session(name string, cfg *confer.Redis) gin.HandlerFunc {
	//return sessions.Sessions(name, sessions.NewCookieStore([]byte("secret")))
	store, _ := sessions.NewRedisStore(10, "tcp", cfg.Addr, "", []byte("secret"))
	return sessions.Sessions(name, store)
}
