package mapi

import (
	"github.com/ZTDBP/ZAManager/app/base/mapi"
	"github.com/ZTDBP/ZAManager/app/v1/access/model/mmysql"
)

type ResourceList struct {
	List     []mmysql.Resource  `json:"list"`
	Paginate mapi.AdminPaginate `json:"paginate"`
}
