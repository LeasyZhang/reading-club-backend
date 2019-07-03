package feature

import (
	"reading-club-backend/database/entity"

	db "reading-club-backend/database"
)

// SaveOrUpdate : save or update entity
func SaveOrUpdate(config entity.Feature) (entity.Feature, error) {
	db.Conn.Save(&config)

	return config, nil
}

// FindFeatureByName : find feature by name
func FindFeatureByName(name string) (feature entity.Feature, errs error) {
	errors := db.Conn.Where("name = ?", name).Find(&feature).GetErrors()

	for _, err := range errors {
		return feature, err
	}

	return feature, nil
}

// Delete : delete feature by name
func Delete(name string) (bool, error) {

	var feature entity.Feature

	errors := db.Conn.Where("name = ?", name).Find(&feature).GetErrors()

	for _, err := range errors {
		return false, err
	}

	db.Conn.Delete(&feature)

	return true, nil
}
