package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/specification"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type CategoryRepoAdapter struct {
	base
}

func (c CategoryRepoAdapter) CountAllCategory(ctx context.Context, spec *dto.CategorySpec) (int64, error) {
	rawQuery := specification.ToCountCategorySpecification(spec)
	var count int64
	if err := c.db.WithContext(ctx).Raw("SELECT COUNT(*) FROM categories " + rawQuery).
		Scan(&count).
		Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (c CategoryRepoAdapter) GetAllCategory(ctx context.Context, spec *dto.CategorySpec) ([]*entity.CategoryEntity, error) {
	rawQuery := specification.ToCategorySpecification(spec)
	var models []*model.CategoryModel
	if err := c.db.WithContext(ctx).
		Raw("SELECT * FROM categories " + rawQuery).
		Scan(&models).
		Limit(*spec.PageSize).
		Offset(*spec.PageSize * (*spec.Page)).Error; err != nil {
		return nil, err
	}
	return mapper.ToListCategoryEntity(models), nil
}

func NewCategoryRepoAdapter(db *gorm.DB) port.ICategoryPort {
	return &CategoryRepoAdapter{
		base: base{db: db},
	}
}
