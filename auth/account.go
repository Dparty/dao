package auth

import (
	"github.com/Dparty/common/utils"
	interfaces "github.com/Dparty/dao/abstract"

	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string `json:"email" gorm:"index:email_index,unique"`
	Password string `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte `json:"salt"`
	Role     string `json:"role" gorm:"type:VARCHAR(128)"`
}

func (a Account) ID() uint {
	return a.Model.ID
}

func (a Account) Own(asset interfaces.Asset) bool {
	return a.ID() == asset.Owner().ID()
}

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}
