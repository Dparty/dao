package auth

import (
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return AccountRepository{
		db: db,
	}
}

func (r AccountRepository) Get(conds ...any) *Account {
	var account Account
	ctx := r.db.Find(account, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}

func (r AccountRepository) GetById(id uint) *Account {
	return r.Get(id)
}

func (r AccountRepository) GetByEmail(email string) *Account {
	return r.Get("email = ?", email)
}

func (r AccountRepository) List(conds ...any) []Account {
	var accounts []Account
	r.db.Find(&accounts, conds...)
	return accounts
}
