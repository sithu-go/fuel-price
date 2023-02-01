package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateFuelLog struct {
	StationId     uint64 `form:"station_id"  json:"station_id"`
	DiselPrice    uint64 `form:"disel_pirce"  json:"disel_pirce"`
	PreDiselPrice uint64 `form:"pre_disel_price"  json:"pre_disel_price"`
	O95Price      uint64 `form:"o95_price"  json:"o95_price"`
	O92Price      uint64 `form:"o92_price"  json:"o92_price"`
}
type UpdateFuelLog struct {
	ID            uint64 `form:"id" json:"id" binding:"required"`
	StationId     uint64 `form:"station_id"  json:"station_id" binding:"required"`
	DiselPrice    uint64 `form:"disel_pirce"  json:"disel_pirce"`
	PreDiselPrice uint64 `form:"pre_disel_price"  json:"pre_disel_price"`
	O95Price      uint64 `form:"o95_price"  json:"o95_price"`
	O92Price      uint64 `form:"o92_price"  json:"o92_price"`
}

type SearchFuelLog struct {
	ID            uint64    `form:"id" json:"id"`
	StationId     uint64    `form:"station_id"  json:"station_id"`
	DiselPrice    uint64    `form:"disel_pirce"  json:"disel_pirce"`
	PreDiselPrice uint64    `form:"pre_disel_price"  json:"pre_disel_price"`
	O95Price      uint64    `form:"o95_price"  json:"o95_price"`
	O92Price      uint64    `form:"o92_price"  json:"o92_price"`
	SDate         time.Time `form:"sdate" json:"sdate"`
	EDate         time.Time `form:"edate" json:"edate"`
	Page          int       `json:"page" form:"page"`
	PageSize      int       `json:"page_size" form:"page_size"`
}

type ResponseFuelLog struct {
	ID            uint64         `json:"id"`
	StationId     uint64         `json:"station_id"`
	DiselPrice    uint64         `json:"disel_pirce"`
	PreDiselPrice uint64         `json:"pre_disel_price"`
	O95Price      uint64         `json:"o95_price"`
	O92Price      uint64         `json:"o92_price"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"-"`
}

type FuelPriceFilter struct {
	Page         int    `json:"page" form:"page" binding:"required"`
	PageSize     int    `json:"page_size" form:"page_size" binding:"required"`
	DivisionId   string `json:"division_id" form:"division_id"`
	DivisionName string `json:"division_name" form:"division_name"`
	StationId    string `json:"station_id" form:"station_id"`
	StationName  string `json:"station_name" form:"station_name"`
}
