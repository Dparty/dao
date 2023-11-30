package auth

import (
	"github.com/Dparty/dao"
)

var db = dao.GetDBInstance()

func init() {
	db.AutoMigrate(&Account{})
}
