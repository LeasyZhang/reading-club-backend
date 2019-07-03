package feature

import (
	featureDao "reading-club-backend/database/dao/feature"
	"reading-club-backend/database/entity"
)

// EnableFeature : enable target feature
func EnableFeature(name string) bool {
	feature, err := featureDao.FindFeatureByName(name)

	if err != nil {
		return false
	}

	feature.Enabled = true
	featureDao.SaveOrUpdate(feature)
	return true
}

// DisableFeature : disable target feature
func DisableFeature(name string) bool {

	feature, err := featureDao.FindFeatureByName(name)

	if err != nil {
		return false
	}

	feature.Enabled = false
	featureDao.SaveOrUpdate(feature)
	return true
}

// AddFeature : add a new feature
func AddFeature(name string) (entity.Feature, error) {

	var feature entity.Feature
	feature.Name = name
	feature.Enabled = false

	res, err := featureDao.SaveOrUpdate(feature)

	if err != nil {
		return res, err
	}

	return feature, nil
}

// DeleteFeature : delete specified feature
func DeleteFeature(name string) bool {
	rlt, _ := featureDao.Delete(name)
	return rlt
}

// FindFeatureState : find state : true|false
func FindFeatureState(name string) bool {
	feature, err := featureDao.FindFeatureByName(name)
	if err != nil {
		return false
	}
	return feature.Enabled
}
