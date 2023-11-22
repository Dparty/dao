package dao

import (
	"fmt"

	"github.com/Dparty/common/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

// GetDBInstance returns the db instance by Lazy bones
func GetDBInstance() *gorm.DB {
	if db == nil {
		var err error
		db, err = NewConnection(
			config.GetString("database.user"),
			config.GetString("database.password"),
			config.GetString("database.host"),
			config.GetString("database.port"),
			config.GetString("database.database"),
		)
		if err != nil {
			panic(err)
		}
	}
	return db
}

func NewConnection(user, password, host, port, database string) (db *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, database,
	)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}
