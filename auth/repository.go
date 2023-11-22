package auth

import (
	"errors"

	"github.com/Dparty/common/utils"
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

var accountRepository *AccountRepository

// GetAccountRepository returns the account repository by Lazy bones
func GetAccountRepository() *AccountRepository {
	if accountRepository == nil {
		accountRepository = NewAccountRepository(dao.GetDBInstance())
	}
	return accountRepository
}

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) *AccountRepository {
	return &AccountRepository{
		db: db,
	}
}

func (r AccountRepository) Find(conds ...any) *Account {
	var account Account
	ctx := r.db.Find(&account, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &account
}

func (r AccountRepository) GetById(id uint) *Account {
	return r.Find(id)
}

func (r AccountRepository) GetByEmail(email string) *Account {
	return r.Find("email = ?", email)
}

func (r AccountRepository) List(conds ...any) []Account {
	var accounts []Account
	r.db.Find(&accounts, conds...)
	return accounts
}

func (r AccountRepository) Save(account *Account) *gorm.DB {
	return r.db.Save(account)
}

var ErrEmailExists = errors.New("email exists")

func (r AccountRepository) Create(email, password, role string) (*Account, error) {
	if account := r.GetByEmail(email); account != nil {
		return nil, ErrEmailExists
	}
	hashed, salt := utils.HashWithSalt(password)
	account := Account{
		Email:    email,
		Password: hashed,
		Salt:     salt,
		Role:     role,
	}
	r.db.Save(&account)
	return &account, nil
}
