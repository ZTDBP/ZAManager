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

package pconst

const (
	//common
	CODE_ERROR_OK  = 0
	CODE_COMMON_OK = 1001
	//CODE_COMMON_ACCESS_FAIL        = 1002
	CODE_COMMON_SERVER_BUSY        = 1003
	CODE_COMMON_PARAMS_INCOMPLETE  = 1004
	CODE_COMMON_DATA_NOT_EXIST     = 1006
	CODE_COMMON_DATA_ALREADY_EXIST = 1007
	CODE_VICTORIA_METRICS_ERR      = 1008
	CODE_DATA_HAS_RELATION         = 1009
	CODE_DATA_WRONG_STATE          = 1010
	CODE_API_BAD_REQUEST           = 1011
	CODE_NOT_FOUND                 = 1012

	CODE_COMMON_USER_NO_LOGIN       = 1005
	CODE_COMMON_ACCESS_FORBIDDEN    = 1013
	ERROR_AUTH_CHECK_TOKEN_FAIL     = 1014
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT  = 1015
	ERROR_AUTH_TOKEN                = 1016
	ERROR_AUTH                      = 1017
	CODE_COMMON_USER_INVALID        = 1018
	ERROR_AUTH_CHECK_SECRET_EXPIRED = 1019
	ERROR_CHECK_CAPTCHA             = 1020
)

/*
注册中心相关
*/
const (
	STATIC_REGISTER_CENTER_TYPE  = 1
	DYNAMIC_REGISTER_CENTER_TYPE = 2
)

// SERVICE_RULE_EXPIRE 治理规则redis过期时间
const SERVICE_RULE_EXPIRE = 60 * 60 * 25

const SITE = "IDG_SITEUID"
const CLUSTER = "IDG_CLUSTERUID"
const NAMESPACE = "NAMESPACE"
const SENTRY_NAMESPACE = "SENTRY_NAMESPACE"
const DEFAULT_NAMESPACE = "msp"
const DEFAULT_SENTRY_NAMESPACE = "sentry"
const SITE_DOMAIN = "SITE_DOMAIN"
const DEFAULT_SITE_DOMAIN = "gw002.oneitfarm.com"
const BLOCKCHAIN_SYNCING = 1
const BLOCKCHAIN_FINISHED = 2

const (
	HTTPRouteGroup = "HTTPRouteGroup"
	TCPRoute       = "TCPRoute"
	UDPRoute       = "UDPRoute"
)
