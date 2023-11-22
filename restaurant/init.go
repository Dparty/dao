package restaurant

import (
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = dao.GetDBInstance()
	db.AutoMigrate(&Restaurant{}, &Table{}, &Printer{}, &Item{}, &Bill{})
}
