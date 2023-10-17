package common

import (
	"database/sql/driver"
	"encoding/json"

	"gorm.io/gorm"
)

type ObjectStorageProvider string

const (
	COS ObjectStorageProvider = "COS"
)

type ObjectStorage struct {
	gorm.Model
	Provider ObjectStorageProvider `json:"provider" gorm:"type:VARCHAR(128)"`
	Url      string                `json:"url"`
}

func (ObjectStorage) GormDataType() string {
	return "JSON"
}

func (s *ObjectStorage) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s ObjectStorage) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
