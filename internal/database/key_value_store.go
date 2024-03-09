package database

import (
	"github.com/minnowo/astoryofand/internal/database/models"
	"gorm.io/gorm/clause"
)

func GetValues(keys []string) (models.KeyValueStoreSlice, error) {

	var values models.KeyValueStoreSlice

	err := GetDB().Find(&values, keys).Error

	if err != nil {

		return nil, err
	}

	return values, nil
}

func SaveKeyValues(values map[string]string) {

	kvalues := make([]models.KeyValueStore, len(values))

	i := 0

	for key, value := range values {

		kvalues[i] = models.KeyValueStore{
			Key: key, Value: value,
		}

		i++
	}

	GetDB().Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(kvalues)
}
