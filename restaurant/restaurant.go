package restaurant

import (
	abstract "github.com/Dparty/common/abstract"
	"github.com/Dparty/common/snowflake"
	"github.com/Dparty/dao"
	"github.com/Dparty/dao/auth"
	"github.com/Dparty/dao/common"
	"gorm.io/gorm"
)

var restaurantIdGenerator = snowflake.NewIdGenertor(1)

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
	Offset      int64
	Categories  common.StringList
}

func (a *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = restaurantIdGenerator.Uint()
	return err
}

// Own implements interfaces.Owner.
func (r Restaurant) Own(asset abstract.Asset) bool {
	return r.ID() == asset.Owner().ID()
}

func (r Restaurant) ID() uint {
	return r.Model.ID
}

func (r Restaurant) Owner() abstract.Owner {
	return restaurantRepository.accountRepoitory.GetById(r.AccountId)
}

func (r *Restaurant) SetOwner(owner abstract.Owner) *Restaurant {
	r.AccountId = owner.ID()
	return r
}

func (r Restaurant) AddTable(table *Table) *Table {
	table.SetOwner(r)
	restaurantRepository.tableRepository.Save(table)
	return table
}

func (r Restaurant) AddItem(item *Item) *Item {
	item.SetOwner(r)
	restaurantRepository.itemRepository.Save(item)
	return item
}

func (r Restaurant) Items() []Item {
	return restaurantRepository.itemRepository.List("restaurant_id = ?", r.ID())
}

func (r Restaurant) AddPrinter(printer *Printer) *Printer {
	printer.SetOwner(r)
	restaurantRepository.printerRepository.Save(printer)
	return printer
}

func (r Restaurant) PickUpCode() int64 {
	var bill Bill
	billRepository.db.Order("pick_up_code DESC").Find(&bill, "restaurant_id = ?", r.ID())
	return (bill.PickUpCode + 1)
}

var restaurantRepository *RestaurantRepository

// GetRestaurantRepository returns the restaurant repository by Lazy bones
func GetRestaurantRepository() *RestaurantRepository {
	if restaurantRepository == nil {
		restaurantRepository = NewRestaurantRepository()
	}
	return restaurantRepository
}

type RestaurantRepository struct {
	db                *gorm.DB
	accountRepoitory  *auth.AccountRepository
	tableRepository   *TableRepository
	itemRepository    *ItemRepository
	printerRepository *PrinterRepository
}

func NewRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		db:                dao.GetDBInstance(),
		accountRepoitory:  auth.GetAccountRepository(),
		tableRepository:   GetTableRepository(),
		itemRepository:    GetItemRepository(),
		printerRepository: GetPrinterRepository(),
	}
}

func (r RestaurantRepository) Get(conds ...any) *Restaurant {
	var restaurant Restaurant
	ctx := r.db.Find(&restaurant, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &restaurant
}

func (r RestaurantRepository) GetById(id uint) *Restaurant {
	return r.Get(id)
}

func (r RestaurantRepository) List(conds ...any) []Restaurant {
	var restaurants []Restaurant
	r.db.Find(&restaurants, conds...)
	return restaurants
}

func (r RestaurantRepository) ListBy(accountId *uint) []Restaurant {
	ctx := r.db.Model(&Restaurant{})
	if accountId != nil {
		ctx.Where("account_id = ?", accountId)
	}
	var restaurants []Restaurant
	ctx.Find(&restaurants)
	return restaurants
}

func (r RestaurantRepository) Create(owner abstract.Owner, name, description string) Restaurant {
	restaurant := Restaurant{
		Name:        name,
		Description: description,
	}
	restaurant.SetOwner(owner)
	r.db.Save(&restaurant)
	return restaurant
}

func (r RestaurantRepository) Save(restaurant *Restaurant) {
	r.db.Save(restaurant)
}

func (r RestaurantRepository) Delete(restaurant *Restaurant) {
	r.db.Delete(restaurant)
}
