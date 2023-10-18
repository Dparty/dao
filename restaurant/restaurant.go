package restaurant

import (
	abstract "github.com/Dparty/dao/abstract"
	"github.com/Dparty/dao/auth"
	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	AccountId   uint
	Name        string
	Description string
}

// Own implements interfaces.Owner.
func (r Restaurant) Own(asset abstract.Asset) bool {
	return r.ID() == asset.Owner().ID()
}

func (r Restaurant) ID() uint {
	return r.Model.ID
}

func (r Restaurant) Owner() *auth.Account {
	return accountRepoitory.GetById(r.AccountId)
}

func (r *Restaurant) SetOwner(owner abstract.Owner) *Restaurant {
	r.AccountId = owner.ID()
	return r
}

func (r Restaurant) AddTable(table *Table) *Table {
	table.SetOwner(r)
	tableRepository.Save(table)
	return table
}

func (r Restaurant) Tables() []Table {
	return tableRepository.List("restaurant_id = ?", r.ID())
}

func (r Restaurant) AddItem(item *Item) *Item {
	item.SetOwner(r)
	itemRepository.Save(item)
	return item
}

func (r Restaurant) Items() []Item {
	return itemRepository.List("restaurant_id = ?", r.ID())
}

func (r Restaurant) AddPrinter(printer *Printer) *Printer {
	printer.SetOwner(r)
	printerRepository.Save(printer)
	return printer
}

func (r Restaurant) Printers() []Printer {
	return printerRepository.List("restaurant_id = ?", r.ID())
}
