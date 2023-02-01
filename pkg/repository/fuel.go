package repository

import (
	"context"
	"fmt"
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/utils"

	"gorm.io/gorm"
)

type fuelRepository struct {
	DB *gorm.DB
}

func newfuelRepository(rConfig *RepoConfig) *fuelRepository {
	return &fuelRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *fuelRepository) FindByField(ctx context.Context, field, value any) (*model.Fuel, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Fuel{})
	fuel := model.Fuel{}
	err := db.First(&fuel, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &fuel, err
}

func (r *fuelRepository) FindOrByField(ctx context.Context, field1, field2, value string) (*model.Fuel, error) {
	db := r.DB.WithContext(ctx).Model(&model.Fuel{})
	fuel := model.Fuel{}
	err := db.First(&fuel, fmt.Sprintf("%s = ? OR %s = ?", field1, field2), value, value).Error
	return &fuel, err
}

func (r *fuelRepository) List(ctx context.Context, req *dto.SearchFuel) ([]*dto.ResponseFuel, int64, error) {
	var list []*dto.ResponseFuel
	db := r.DB.Debug().Table("fuels")
	db.Select("fuels.*, stations.name as division")
	// db.Joins("JOIN app_clients ON app_clients.client_id = clients.id")
	db.Joins("LEFT JOIN division ON stations.id = fuels.division_id")
	db.Where("fuels.deleted_at IS NULL")
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

func (r *fuelRepository) Create(ctx context.Context, station *model.Fuel) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Fuel{})
	return db.Create(&station).Error
}

func (r *fuelRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Fuel{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *fuelRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Fuel{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.Fuel{}).Error
}
