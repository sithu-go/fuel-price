package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateStation struct {
	Name       string `form:"name"  json:"name" binding:"required"`
	DivisionId uint64 `form:"station_id"  json:"station_id" binding:"required"`
}
type UpdateStation struct {
	ID         uint64 `form:"id" json:"id" binding:"required"`
	Name       string `form:"name"  json:"name"`
	DivisionId uint64 `form:"station_id"  json:"station_id"`
}

type SearchStation struct {
	ID         uint64    `form:"id" json:"id"`
	Name       string    `form:"name" json:"name" `
	DivisionId uint64    `form:"station_id" json:"station_id" `
	SDate      time.Time `form:"sdate" json:"sdate"`
	EDate      time.Time `form:"edate" json:"edate"`
	Page       int       `json:"page" form:"page"`
	PageSize   int       `json:"page_size" form:"page_size"`
}

type ResponseStation struct {
	ID         uint64         `json:"id"`
	Name       string         `json:"name"`
	DivisionId uint64         `json:"station_id"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
