package auth

import (
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = dao.GetDBInstance()
	db.AutoMigrate(&Account{})
}
