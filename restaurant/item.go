package restaurant

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"github.com/Dparty/common/fault"
	"github.com/Dparty/common/utils"
	abstract "github.com/Dparty/dao/abstract"
	"github.com/Dparty/dao/common"
	"gorm.io/gorm"
)

type Item struct {
	gorm.Model
	RestaurantId uint              `json:"restaurantId" gorm:"index:restaurant_index"`
	Name         string            `json:"name"`
	Pricing      int64             `json:"pricing"`
	Attributes   Attributes        `json:"attributes"`
	Images       common.StringList `json:"images" gorm:"type:JSON"`
	Tags         common.StringList `json:"tags"`
	Printers     common.IDList     `json:"printers"`
	Status       string            `json:"status"`
}

func (a *Item) BeforeCreate(tx *gorm.DB) (err error) {
	a.Model.ID = utils.GenerteId()
	return err
}

func (i Item) ID() uint {
	return i.Model.ID
}

func (i *Item) SetOwner(owner abstract.Owner) *Item {
	i.RestaurantId = owner.ID()
	return i
}

func (i Item) Owner() *Restaurant {
	return restaurantRepository.GetById(i.RestaurantId)
}

func (i Item) CreateOrder(specification []Pair) (Order, error) {
	// TODO: specification verification
	return Order{
		Item:          i,
		Specification: specification,
	}, nil
}

type Attributes []Attribute

func (as Attributes) GetOption(left, right string) (Pair, error) {
	for _, a := range as {
		if left == a.Label {
			for _, option := range a.Options {
				if right == option.Label {
					return Pair{Left: left, Right: right}, nil
				}
			}
		}
	}
	return Pair{}, errors.New("NotFound")
}

func (Attributes) GormDataType() string {
	return "JSON"
}

func (s *Attributes) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Attributes) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Attribute struct {
	Label   string   `json:"label"`
	Options []Option `json:"options"`
}

type Option struct {
	Label string `json:"label"`
	Extra int64  `json:"extra"`
}

type Options []Option

func (Options) GormDataType() string {
	return "JSON"
}

func (s *Options) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Options) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

func (Attribute) GormDataType() string {
	return "JSON"
}

func (s *Attribute) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s Attribute) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type Pair struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return ItemRepository{
		db: db,
	}
}

func (i ItemRepository) Get(conds ...any) *Item {
	var item Item
	ctx := i.db.Find(&item, conds...)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &item
}

func (i ItemRepository) GetById(id uint) *Item {
	return i.Get(id)
}

func (i ItemRepository) Save(item *Item) (*Item, error) {
	var attributesMap map[string]bool = make(map[string]bool)
	for _, attribute := range item.Attributes {
		_, ok := attributesMap[attribute.Label]
		if ok {
			return nil, fault.ErrItemAttributesConflict
		}
		attributesMap[attribute.Label] = true
		var optionMap map[string]bool = make(map[string]bool)
		for _, option := range attribute.Options {
			_, ok := optionMap[option.Label]
			if ok {
				return nil, fault.ErrItemAttributesConflict
			}
			optionMap[option.Label] = true
		}
	}
	i.db.Save(item)
	return item, nil
}

func (i ItemRepository) List(conds ...any) []Item {
	var items []Item
	i.db.Find(&items, conds...)
	return items
}

func (i ItemRepository) Delete(item *Item) *gorm.DB {
	return i.db.Delete(item)
}
