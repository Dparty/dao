package verification

import (
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

var verificationCodeRepository *VerificationCodeRepository

func GetVerificationCodeRepository() *VerificationCodeRepository {
	if verificationCodeRepository == nil {
		verificationCodeRepository = NewVerificationCodeRepository()
	}
	return verificationCodeRepository
}

func NewVerificationCodeRepository() *VerificationCodeRepository {
	if verificationCodeRepository == nil {
		verificationCodeRepository = &VerificationCodeRepository{dao.GetDBInstance()}
	}
	return verificationCodeRepository
}

// Repository of VerificationCode
type VerificationCodeRepository struct {
	db *gorm.DB
}

func (r VerificationCodeRepository) FindOne(conds ...any) *VerificationCode {
	var verificationCode VerificationCode
	ctx := r.db.Find(&verificationCode, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &verificationCode
}

func (r VerificationCodeRepository) FindMany(conds ...any) []VerificationCode {
	var verificationCodes []VerificationCode
	r.db.Find(&verificationCodes, conds...)
	return verificationCodes
}

func (r VerificationCodeRepository) Create(verificationCode *VerificationCode) error {
	return r.db.Save(verificationCode).Error
}

func (r VerificationCodeRepository) Update(verificationCode *VerificationCode) error {
	return r.db.Save(verificationCode).Error
}

func (r VerificationCodeRepository) Delete(verificationCode *VerificationCode) error {
	return r.db.Delete(verificationCode).Error
}
