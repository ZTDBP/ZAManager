package mapi

import (
	"github.com/ZTDBP/ZAManager/app/base/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
)

type ServerList struct {
	List     []mmysql.Server    `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
