package mapi

import (
	"github.com/ZTDBP/ZAManager/app/base/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
)

type ClientList struct {
	List     []mmysql.Client    `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
