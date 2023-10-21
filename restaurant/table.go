package restaurant

import (
	"github.com/Dparty/common/utils"
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

func (a *Table) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
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

func (t Table) Bills() []Bill {
	return billRepository.List("table_id = ?", t.ID())
}

type TableRepository struct {
	db             *gorm.DB
	billRepository BillRepository
}

func NewTableRepository(db *gorm.DB) TableRepository {
	return TableRepository{
		db:             db,
		billRepository: NewBillRepository(db),
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
