package auth

import "gorm.io/gorm"

var db *gorm.DB

var accountRepository AccountRepository

func Init(inject *gorm.DB) {
	db = inject
	db.AutoMigrate(&Account{})
	accountRepository = NewAccountRepository(db)
}
