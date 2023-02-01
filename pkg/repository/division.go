package repository

import (
	"context"
	"fmt"
	"fuel-price/pkg/model"

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
	admin := model.Division{}
	err := db.First(&admin, fmt.Sprintf("BINARY %s = ?", field), value).Error
	return &admin, err
}

func (r *divisionRepository) FindOrByField(ctx context.Context, field1, field2, value string) (*model.Division, error) {
	db := r.DB.WithContext(ctx).Model(&model.Division{})
	admin := model.Division{}
	err := db.First(&admin, fmt.Sprintf("%s = ? OR %s = ?", field1, field2), value, value).Error
	return &admin, err
}

// func (r *divisionRepository) List(ctx context.Context, req *dto.AdminSS) ([]*dto.ResponseAdmin, int64, error) {
// 	var list []*dto.ResponseAdmin
// 	db := r.DB.Debug().Table("admins")

// 	db.Select("admins.*, roles.name as role_name")
// 	// db.Joins("JOIN app_clients ON app_clients.client_id = clients.id")
// 	db.Joins("LEFT JOIN roles ON admins.role_id = roles.id")
// 	db.Where("admins.deleted_at IS NULL")
// 	if req.ID != 0 {
// 		db.Where("admins.id", req.ID)
// 	}
// 	if req.Username != "" {
// 		db.Where("admins.username LIKE ?", "%"+req.Username+"%")
// 	}
// 	if req.Email != "" {
// 		db.Where("admins.email LIKE ?", "%"+req.Email+"%")
// 	}
// 	var total int64
// 	db.Count(&total)
// 	if err := db.Scopes(utils.Paginate(req.Page, req.PageSize)).Find(&list).Error; err != nil {
// 		return nil, 0, err
// 	}
// 	return list, total, nil
// }

// func (r *divisionRepository) Create(ctx context.Context, admin *model.Division) error {
// 	tb := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
// 	var count int64
// 	var oldAdmin model.Admin
// 	tb.Unscoped().Where("username = ? AND deleted_at IS NOT NULL", admin.Username).Count(&count).First(&oldAdmin)
// 	if count > 0 {
// 		admin.ID = oldAdmin.ID
// 		admin.CreatedAt = time.Now()
// 		admin.UpdatedAt = time.Now()
// 		admin.DeletedAt = gorm.DeletedAt{
// 			Time:  time.Time{},
// 			Valid: false,
// 		}
// 		return tb.Save(&admin).Error
// 	}
// 	db := r.DB.WithContext(ctx).Debug().Model(&model.Division{})
// 	return db.Create(&admin).Error
// }
