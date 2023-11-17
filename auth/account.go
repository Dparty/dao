package auth

import (
	"time"

	abstract "github.com/Dparty/common/abstract"
	"github.com/Dparty/common/snowflake"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	Email    string `json:"email" gorm:"index:email_index,unique"`
	Password string `json:"password" gorm:"type:CHAR(128)"`
	Salt     []byte `json:"salt"`
	Role     string `json:"role" gorm:"type:VARCHAR(128)"`
	Gender   string `json:"gender"`
	Birthday time.Time
}

func (a Account) ID() uint {
	return a.Model.ID
}

func (a Account) Own(asset abstract.Asset) bool {
	return abstract.Own(a, asset)
}

var accountIdGenerator = snowflake.NewIdGenertor(1)

func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = accountIdGenerator.Uint()
	return err
}

func (a Account) Owner() abstract.Owner {
	return nil
}
