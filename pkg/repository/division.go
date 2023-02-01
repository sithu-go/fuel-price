package repository

import (
	"context"
	"fmt"
	"fuel-price/pkg/dto"
	"fuel-price/pkg/model"
	"fuel-price/pkg/utils"

	"gorm.io/gorm"
)

type divisionRepository struct {
	DB *gorm.DB
}

func newdivisionRepository(rConfig *RepoConfig) *divisionRepository {
	return &divisionRepository{
		DB: rConfig.DS.DB,
	}
}

func (r *divisionRepository) FindByField(ctx context.Context, field, value any) (*model.Division, error) {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
	division := model.Division{}
	err := db.First(&division, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &division, err
}

func (r *divisionRepository) FindOrByField(ctx context.Context, field1, field2, value string) (*model.Division, error) {
	db := r.DB.WithContext(ctx).Model(&model.Division{})
	division := model.Division{}
	err := db.First(&division, fmt.Sprintf("%s = ? OR %s = ?", field1, field2), value, value).Error
	return &division, err
}

func (r *divisionRepository) List(ctx context.Context, req *dto.SearchDivision) ([]*dto.ResponseDivision, int64, error) {
	var list []*dto.ResponseDivision
	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
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

func (r *divisionRepository) Create(ctx context.Context, division *model.Division) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
	return db.Create(&division).Error
}

func (r *divisionRepository) Update(ctx context.Context, updateFields *model.UpdateFields) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
	db.Where(updateFields.Field, updateFields.Value)
	return db.Updates(&updateFields.Data).Error
}

func (r *divisionRepository) Delete(ctx context.Context, ids string) error {
	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
	db.Where(fmt.Sprintf("id in (%s)", ids))
	return db.Delete(&model.Division{}).Error
}
