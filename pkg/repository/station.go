package repository

import (
	"context"
	"fmt"
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/utils"

	"gorm.io/gorm"
)

type stationRepository struct {
	DB *gorm.DB
}

func newStationRepository(rConfig *RepoConfig) *stationRepository {
	return &stationRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *stationRepository) FindByField(ctx context.Context, field, value any) (*model.Station, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Station{})
	station := model.Station{}
	err := db.First(&station, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &station, err
}

func (r *stationRepository) FindOrByField(ctx context.Context, field1, field2, value string) (*model.Station, error) {
	db := r.DB.WithContext(ctx).Model(&model.Station{})
	station := model.Station{}
	err := db.First(&station, fmt.Sprintf("%s = ? OR %s = ?", field1, field2), value, value).Error
	return &station, err
}

func (r *stationRepository) List(ctx context.Context, req *dto.SearchStation) ([]*dto.ResponseStation, int64, error) {
	var list []*dto.ResponseStation
	db := r.DB.Debug().Table("stations")
	db.Select("stations.*, divisions.name as division")
	// db.Joins("JOIN app_clients ON app_clients.client_id = clients.id")
	db.Joins("LEFT JOIN division ON divisions.id = stations.division_id")
	db.Where("stations.deleted_at IS NULL")
	if req.ID != 0 {
		db.Where("id", req.ID)
	}
	if req.Name != "" {
		db.Where("name LIKE ?", "%"+req.Name+"%")
	}
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *stationRepository) Create(ctx context.Context, station *model.Station) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Station{})
	return db.Create(&station).Error
}

func (r *stationRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Station{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *stationRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Station{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.Station{}).Error
}
