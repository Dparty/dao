package common

import (
	"database/sql/driver"
	"encoding/json"
)

type StringList []string

func StringBackward(s StringList) []string {
	var slist []string = make([]string, 0)
	for _, v := range s {
		slist = append(slist, v)
	}
	return slist
}

func StringForward(s []string) StringList {
	var slist StringList = make(StringList, 0)
	for _, v := range s {
		slist = append(slist, v)
	}
	return slist
}

func (StringList) GormDataType() string {
	return "JSON"
}

func (s *StringList) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s StringList) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}

type IDList []uint

func IDListBackward(s StringList) []string {
	var slist []string = make([]string, 0)
	for _, v := range s {
		slist = append(slist, v)
	}
	return slist
}

func IDListForward(s []string) StringList {
	var slist StringList = make(StringList, 0)
	for _, v := range s {
		slist = append(slist, v)
	}
	return slist
}

func (IDList) GormDataType() string {
	return "JSON"
}

func (s *IDList) Scan(value any) error {
	return json.Unmarshal(value.([]byte), s)
}

func (s IDList) Value() (driver.Value, error) {
	b, err := json.Marshal(s)
	return b, err
}
