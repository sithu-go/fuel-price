package repository

import (
	"fuel-price/pkg/ds"
)

type Repository struct {
	DS       *ds.DataSource
	Division *divisionRepository
}

type RepoConfig struct {
	DS *ds.DataSource
}

func NewRepository(rConfig *RepoConfig) *Repository {

	return &Repository{
		DS:       rConfig.DS,
		Division: newdivisionRepository(rConfig),
	}
}
