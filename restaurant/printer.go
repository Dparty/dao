package restaurant

import (
	abstract "github.com/Dparty/dao/abstract"
	"gorm.io/gorm"
)

type PrinterType string

const (
	BILL    PrinterType = "BILL"
	KITCHEN PrinterType = "KITCHEN"
)

type Printer struct {
	gorm.Model
	RestaurantId uint
	Name         string      `json:"name"`
	Sn           string      `json:"sn"`
	Description  string      `json:"description"`
	Type         PrinterType `json:"type" gorm:"type:VARCHAR(128)"`
}

func (p Printer) ID() uint {
	return p.Model.ID
}

func (p Printer) Owner() abstract.Owner {
	return restaurantRepository.GetById(p.RestaurantId)
}

func (p *Printer) SetOwner(owner abstract.Owner) *Printer {
	p.Model.ID = owner.ID()
	return p
}

func NewPrinterRepository(db *gorm.DB) PrinterRepository {
	return PrinterRepository{
		db: db,
	}
}

type PrinterRepository struct {
	db *gorm.DB
}

func (p PrinterRepository) Save(printer *Printer) *Printer {
	p.db.Save(printer)
	return printer
}

func (p PrinterRepository) Find(conds ...any) *Printer {
	var printer Printer
	ctx := p.db.Find(&printer, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &printer
}

func (p PrinterRepository) GetById(id uint) *Printer {
	return p.Find(id)
}

func (p PrinterRepository) List(conds ...any) []Printer {
	var printers []Printer
	p.db.Find(&printers, conds...)
	return printers
}
