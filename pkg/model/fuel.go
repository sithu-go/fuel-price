package model

import (
	"gorm.io/gorm"
)

type FuelLog struct {
	gorm.Model
	StationId     uint64 `gorm:"column:station_id" json:"station_id"`
	DiselPrice    uint64 `gorm:"column:disel_price;" json:"disel_price"`
	PreDiselPrice uint64 `gorm:"column:pre_disel_price;" json:"pre_disel_price"`
	O95Price      uint64 `gorm:"column:o95_price;" json:"o95_price"`
	O92Price      uint64 `gorm:"column:o92_price;" json:"o92_price"`
}
