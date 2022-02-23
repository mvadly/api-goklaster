package repositories

import (
	"api-goklaster/entities"
)

func GetBanners() ([]entities.ClusterLandingPage, error) {
	var banner []entities.ClusterLandingPage
	query := db.Find(&banner)
	if query.Error != nil {
		return nil, query.Error
	}
	return banner, nil
}
