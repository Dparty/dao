package restaurant

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

var restaurantRepository RestaurantRepository
var billRepository BillRepository

func Init(inject *gorm.DB) {
	db = inject
	restaurantRepository = NewRestaurantRepository(db)
	billRepository = NewBillRepository(db)
	db.AutoMigrate(&Restaurant{}, &Table{}, &Printer{}, &Item{})
	db.AutoMigrate(&Bill{})
	fmt.Println(", &Bill{}")
}
