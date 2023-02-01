package repository

import (
	"fuel-price/pkg/ds"
	"fuel-price/pkg/model"
	"log"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type fuelRepository struct {
	DB *gorm.DB
}

func newFuelRepository(rConfig *RepoConfig) *fuelRepository {
	return &fuelRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *fuelRepository) GetFuelPrices() {
	tb := ds.DB.Debug().Model(&model.FuelLog{})
	filterToQuery(tb)
	prices := make([]*model.Division, 0)
	if err := tb.Preload(clause.Associations).Find(&prices).Error; err != nil {
		log.Println(err)
		return
	}

	log.Println(prices)

}

// utilities start
func filterToQuery(tb *gorm.DB) {
	// tb.Table("divisions AS d")
	// tb.Joins("LEFT JOIN stations s ON s.division_id = d.id")
	// tb.Joins("LEFT JOIN fuel_logs fl ON fl.station_id = s.id")
	// tb.Order("s.created_at")
	// tb.Group("fl.station_id")

	tb.Table("fuel_logs AS fl")
	tb.Joins("LEFT JOIN stations s ON fl.station_id = s.id")
	tb.Joins("LEFT JOIN divisions d ON s.division_id = d.id")
	tb.Order("s.created_at")
	tb.Group("fl.station_id")

}

// utilities end
