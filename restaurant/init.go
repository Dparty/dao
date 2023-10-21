package restaurant

import (
	"gorm.io/gorm"
)

var db *gorm.DB

var restaurantRepository RestaurantRepository
var billRepository BillRepository

func Init(inject *gorm.DB) {
	db = inject
	db.AutoMigrate(&Restaurant{}, &Table{}, &Printer{}, &Item{}, &Bill{})
	restaurantRepository = NewRestaurantRepository(db)
	billRepository = NewBillRepository(db)
}
