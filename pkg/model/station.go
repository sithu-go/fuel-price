package model

import (
	"time"

	"gorm.io/gorm"
)

type Station struct {
	ID         uint64 `gorm:"column:id;primaryKey" json:"id"`
	Name       string `gorm:"column:name;type:varchar(100);not null" json:"name"`
	DivisionId uint64 `gorm:"column:division_id" json:"dividion_id"`
	Division   Division
	CreatedAt  time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt  time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt  gorm.DeletedAt `json:"-"`
}
