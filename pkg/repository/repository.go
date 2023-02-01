package repository

import (
	"fuel-price/pkg/ds"
)

type Repository struct {
	DS       *ds.DataSource
	Division *divisionRepository
	Station  *stationRepository
	FuelLog  *FuelLogRepository
}

type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(rConfig *RepoConfig) *Repository {

	return &Repository{
		DS:       rConfig.DS,
		Division: newdivisionRepository(rConfig),
		Station:  newStationRepository(rConfig),
		FuelLog:  newFuelLogRepository(rConfig),
	}
}
