package restaurant

import (
	abstract "github.com/Dparty/common/abstract"
	"github.com/Dparty/common/utils"
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
	Offset      int64
}

func (a *Restaurant) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
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

func (r Restaurant) Tables() []Table {
	return restaurantRepository.tableRepository.List("restaurant_id = ?", r.ID())
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

func (r Restaurant) Printers() []Printer {
	return restaurantRepository.printerRepository.List("restaurant_id = ?", r.ID())
}

func (r Restaurant) PickUpCode() int64 {
	var bill Bill
	billRepository.db.Order("pick_up_code DESC").Find(&bill, "restaurant_id = ?", r.ID())
	return (bill.PickUpCode + 1) % 1000
}
