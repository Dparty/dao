package restaurant

import (
	"gorm.io/gorm"
)

var db *gorm.DB

var restaurantRepository RestaurantRepository

func Init(inject *gorm.DB) {
	db = inject
	restaurantRepository = NewRestaurantRepository(db)
	db.AutoMigrate(&Restaurant{}, &Table{}, &Printer{}, &Item{})
}
