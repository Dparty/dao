package restaurant

import (
	"github.com/Dparty/dao/repository"
	"gorm.io/gorm"
)

type Offset struct {
	gorm.Model
	RestaurantId uint
	Label        string
	Offset       int64
}

type OffsetRepository repository.Repository[Offset]
