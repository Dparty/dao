package user

import (
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

var db = dao.GetDBInstance()

func init() {
	db.AutoMigrate(&User{})
}

type User struct {
	gorm.Model
	AreaCode    string `json:"areaCode" gorm:"type:CHAR(8);index:area_code_index"`
	PhoneNumber string `json:"phoneNumber" gorm:"type:CHAR(11);index:phone_number_index"`
	Password    string `json:"password" gorm:"type:CHAR(128)"`
	Salt        []byte `json:"salt"`
}
