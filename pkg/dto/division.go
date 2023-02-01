package dto

import (
	"time"

	"gorm.io/gorm"
)

type CreateDivision struct {
	Name string `form:"name"  json:"name" binding:"required"`
}
type UpdateDivision struct {
	ID   uint64 `form:"id" json:"id" binding:"required"`
	Name string `form:"name"  json:"name"`
}

type SearchDivision struct {
	ID       uint64    `form:"id" json:"id"`
	Name     string    `form:"name" json:"name" `
	SDate    time.Time `form:"sdate" json:"sdate"`
	EDate    time.Time `form:"edate" json:"edate"`
	Page     int       `json:"page" form:"page"`
	PageSize int       `json:"page_size" form:"page_size"`
}

type ResponseDivision struct {
	ID        uint64         `json:"id"`
	Name      string         `json:"name"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
