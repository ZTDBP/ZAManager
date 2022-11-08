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

package controller

import (
	"github.com/ZTDBP/ZAManager/pconst"
	"github.com/ZTDBP/ZAManager/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code int      `json:"code"`
	Data struct{} `json:"data"`
	Msg  string   `json:"message"`
}

func BindParams(c *gin.Context, params interface{}) (b bool, code int) {
	err := c.ShouldBind(params)
	if err != nil {
		logger.Warnf(c, "params error: %v", err)
		code = pconst.CODE_COMMON_PARAMS_INCOMPLETE
		return
	}
	b = true
	return
}
