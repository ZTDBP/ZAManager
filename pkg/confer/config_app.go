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

package confer

import (
	"strings"
)

func ConfigAppGetString(key string, defaultConfig string) string {
	config := ConfigAppGet(key)
	if config == nil {
		return defaultConfig
	} else {
		configStr := config.(string)
		if len(strings.TrimSpace(configStr)) == 0 {
			configStr = defaultConfig
		}
		return configStr
	}
}

func ConfigAppGetInt(key string, defaultConfig int) int {
	config := ConfigAppGet(key)
	if config == nil {
		return defaultConfig
	} else {
		configInt := config.(int)
		if configInt == 0 {
			configInt = defaultConfig
		}
		return configInt
	}
}

func ConfigAppGet(key string) interface{} {
	globalConfig.RLock()
	defer globalConfig.RUnlock()
	//将配置文件中的app字段转为map
	config, exists := globalConfig.App[key]
	if !exists {
		return nil
	}
	return config
}

func ConfigEnvGet() string {
	strEnv := ConfigAppGet("env")
	return strEnv.(string)
}

func ConfigEnvIsDev() bool {
	env := ConfigEnvGet()
	if env == "dev" {
		return true
	}
	return false
}

func ConfigEnvIsBeta() bool {
	return ConfigEnvGet() == "beta"
}
