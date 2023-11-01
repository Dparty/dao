package auth

import (
	"gorm.io/gorm"
)

type VerificationCode struct {
	gorm.Model
	Identity string `json:"identity" gorm:"index:verification_identity_index"`
	Code     string `json:"code" gorm:"type:VARCHAR(12)"`
	Tries    int64
}

type VerificationCodeRepository struct {
	db *gorm.DB
}

func NewVerificationCodeRepository(db *gorm.DB) VerificationCodeRepository {
	return VerificationCodeRepository{
		db: db,
	}
}
