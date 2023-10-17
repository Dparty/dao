package restaurant

import (
	"github.com/Dparty/dao/auth"
	"gorm.io/gorm"
)

var db *gorm.DB

var accountRepoitory auth.AccountRepository
var restaurantRepository RestaurantRepository
var tableRepository TableRepository
var itemRepository ItemRepository
var printerRepository PrinterRepository

func Init(inject *gorm.DB) {
	db = inject
	restaurantRepository = NewRestaurantRepository(db)
	printerRepository = NewPrinterRepository(db)
	accountRepoitory = auth.NewAccountRepository(db)
	tableRepository = NewTableRepository(db)
	itemRepository = NewItemRepository(db)
	db.AutoMigrate(&Restaurant{}, &Table{}, &Printer{}, &Item{})
}
