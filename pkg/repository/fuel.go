package repository

import (
	"context"
	"fmt"
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/utils"

	"gorm.io/gorm"
)

type FuelLogRepository struct {
	DB *gorm.DB
}

func newFuelLogRepository(rConfig *RepoConfig) *FuelLogRepository {
	return &FuelLogRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *FuelLogRepository) FindByField(ctx context.Context, field, value any) (*model.FuelLog, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.FuelLog{})
	FuelLog := model.FuelLog{}
	err := db.First(&FuelLog, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &FuelLog, err
}

func (r *FuelLogRepository) FindOrByField(ctx context.Context, field1, field2, value string) (*model.FuelLog, error) {
	db := r.DB.WithContext(ctx).Model(&model.FuelLog{})
	FuelLog := model.FuelLog{}
	err := db.First(&FuelLog, fmt.Sprintf("%s = ? OR %s = ?", field1, field2), value, value).Error
	return &FuelLog, err
}

func (r *FuelLogRepository) List(ctx context.Context, req *dto.SearchFuelLog) ([]*dto.ResponseFuelLog, int64, error) {
	var list []*dto.ResponseFuelLog
	db := r.DB.Debug().Table("fuel_logs")
	db.Select("fuel_logs.*, stations.name as division")
	// db.Joins("JOIN app_clients ON app_clients.client_id = clients.id")
	db.Joins("LEFT JOIN division ON stations.id = fuel_logs.division_id")
	db.Where("fuel_logs.deleted_at IS NULL")
	if req.ID != 0 {
		db.Where("id", req.ID)
	}
	var total int64
	db.Count(&total)
	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *FuelLogRepository) Create(ctx context.Context, station *model.FuelLog) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.FuelLog{})
	return db.Create(&station).Error
}

func (r *FuelLogRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.FuelLog{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *FuelLogRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.FuelLog{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.FuelLog{}).Error
}
