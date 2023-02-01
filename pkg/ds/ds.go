package ds

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type DataSource struct {
	DB *gorm.DB
}

func NewDataSource() (*DataSource, error) {
	db, err := LoadDB()
	if err != nil {
		return nil, err
	}

	DB = db

	return &DataSource{
		DB: db,
	}, nil
}
