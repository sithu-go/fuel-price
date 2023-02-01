package repository

import (
	"fuel-price/pkg/ds"
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/utils"
	"log"

	"gorm.io/gorm"
)

type fuelRepository struct {
	DB *gorm.DB
}

func newFuelRepository(rConfig *RepoConfig) *fuelRepository {
	return &fuelRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *fuelRepository) GetFuelPrices(req *dto.FuelPriceFilter) ([]*model.Station, int64, error) {
	tb := ds.DB.Debug().Model(&model.Station{})
	filterToQuery(tb, req)
	prices := make([]*model.Station, 0)
	var total int64
	tb.Count(&total)
	tb.Scopes(utils.Paginate(req.Page, req.PageSize))
	if err := tb.Find(&prices).Error; err != nil {
		log.Println(err)
		return nil, 0, err
	}
	return prices, total, nil
}

// utilities start
func filterToQuery(tb *gorm.DB, req *dto.FuelPriceFilter) {
	tb.Table("stations AS s")
	tb.Joins("LEFT JOIN divisions d ON s.division_id = d.id")
	tb.Joins("LEFT JOIN fuel_logs fl ON fl.station_id = s.id")
	tb.Order("fl.created_at DESC")

	if req.DivisionId != "" {
		tb.Where("d.id", req.DivisionId)
	}
	if req.DivisionName != "" {
		tb.Where("d.name LIKE ?", "%"+req.DivisionName+"%")
	}

	if req.StationId != "" {
		tb.Where("s.id", req.StationId)
	}
	if req.StationName != "" {
		tb.Where("d.name LIKE ?", "%"+req.StationName+"%")
	}

}

// utilities end
