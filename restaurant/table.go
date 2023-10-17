package restaurant

import (
	abstract "github.com/Dparty/dao/abstract"
	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	RestaurantId uint
	Label        string `json:"label"`
	X            int64  `json:"x"`
	Y            int64  `json:"y"`
}

// Owner implements interfaces.Asset.
func (t *Table) Owner() abstract.Owner {
	return restaurantRepository.GetById(t.RestaurantId)
}

func (t Table) ID() uint {
	return t.Model.ID
}

func (t *Table) SetOwner(owner abstract.Owner) *Table {
	t.Model.ID = owner.ID()
	return t
}

type TableRepository struct {
	db *gorm.DB
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return TableRepository{
		db: db,
	}
}

func (t TableRepository) Find(conds ...any) *Table {
	var table Table
	ctx := t.db.Find(&table, conds)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &table
}

func (t TableRepository) List(conds ...any) []Table {
	var tables []Table
	t.db.Find(&tables, conds...)
	return tables
}

func (t TableRepository) Save(table *Table) *gorm.DB {
	return t.db.Save(table)
}

func (t TableRepository) Delete(table *Table) *gorm.DB {
	return t.db.Delete(&table)
}
