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

// Constants for API
const (
	IAPIRoot      = "/i"
	APPAPIRoot    = "/app"
	APIAPIRoot    = "/api"
	ADMINAPIRoot  = "/admin"
	APIV1Version  = "v1"
	APIV2Version  = "v2"
	APIV3Version  = "v3"
	APIV4Version  = "v4"
	IAPIV1URL     = IAPIRoot + "/" + APIV1Version
	IAPIV2URL     = IAPIRoot + "/" + APIV2Version
	IAPIV3URL     = IAPIRoot + "/" + APIV3Version
	IAPIV4URL     = IAPIRoot + "/" + APIV4Version
	APPAPIV1URL   = APPAPIRoot + "/" + APIV1Version
	APIAPIV1URL   = APIAPIRoot + "/" + APIV1Version
	ADMINAPIV1URL = ADMINAPIRoot + "/" + APIV1Version
	APPAPIV2URL   = APPAPIRoot + "/" + APIV2Version
	APIAPIV2URL   = APIAPIRoot + "/" + APIV2Version
	//time

	TIME_FORMAT_Y_M_D_H_I_S_MS = "2006-01-02 15:04:05.000"
	TIME_FORMAT_Y_M_D_H_I_S    = "2006-01-02 15:04:05"
	TIME_FORMAT_Y_M_D_H_I_S_2  = "2006/01/02 15:04:05"
	TIME_FORMAT_Y_M_D          = "2006.01.02"
	TIME_FORMAT_Y_M_D_         = "2006-01-02"
	TIME_FORMAT_Y_MS_D_        = "2006-January-02"
	TIME_FORMAT_M_D_H_I        = "01.02 15:04"
	TIME_FORMAT_H_I            = "15:04"

	TIME_ONE_SECOND = 1
	TIME_TWO_SECOND = 2
	TIME_ONE_MINUTE = 60
	TIME_TEN_MINUTE = 600
	TIME_ONE_HOUR   = 3600
	TIME_ONE_DAY    = 86400
	TIME_THREE_DAY  = 259200
	TIME_ONE_WEEK   = 604800
	TIME_ONE_MONTH  = 2592000
	TIME_ONE_YEAR   = 31536000

	//common

	COMMON_PAGE_LIMIT_NUM_10  = 10
	COMMON_PAGE_LIMIT_NUM_20  = 20
	COMMON_PAGE_LIMIT_NUM_MAX = 20

	COMMON_ERROR_RETRY_LIMIT_NUM = 5
)
