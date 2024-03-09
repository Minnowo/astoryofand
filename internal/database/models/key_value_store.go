package models

import (
	"time"

	"gorm.io/gorm"
)

type KeyValueStore struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Key   string `gorm:"primarykey"`
	Value string
}

type KeyValueStoreSlice []KeyValueStore

func (kvSlice KeyValueStoreSlice) ToMap() map[string]string {

	result := make(map[string]string)

	for _, kv := range kvSlice {
		result[kv.Key] = kv.Value
	}

	return result
}
