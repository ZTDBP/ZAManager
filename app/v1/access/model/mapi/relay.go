package mapi

import (
	"github.com/ZTDBP/ZAManager/app/base/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
)

type RelayList struct {
	List     []mmysql.Relay     `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
