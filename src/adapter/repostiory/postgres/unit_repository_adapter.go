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

type UnitRepositoryAdapter struct {
	base
}

func (u UnitRepositoryAdapter) SaveUnit(ctx context.Context, tx *gorm.DB, unit *entity.UnitEntity) (*entity.UnitEntity, error) {
	unitModel := mapper.ToUnitModel(unit)
	if err := tx.WithContext(ctx).Model(&model.UnitModel{}).Create(unitModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUnitEntity(unitModel), nil
}

func (u UnitRepositoryAdapter) GetAllUnit(ctx context.Context, spec *dto.UnitParamDto) ([]*entity.UnitEntity, error) {
	var unitModels []*model.UnitModel
	if err := u.db.WithContext(ctx).
		Raw("SELECT * FROM units " + specification.ToUnitSpecification(spec)).
		Scan(&unitModels).
		Limit(*spec.PageSize).
		Offset(*spec.PageSize * (*spec.Page)).Error; err != nil {
		return nil, err
	}
	return mapper.ToListUnitEntity(unitModels), nil
}

func (u UnitRepositoryAdapter) CountAllUnit(ctx context.Context, spec *dto.UnitParamDto) (int64, error) {
	var count int64
	if err := u.db.WithContext(ctx).
		Raw("SELECT COUNT(*) FROM units " + specification.ToCountUnitSpecification(spec)).
		Scan(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func NewUnitRepositoryAdapter(db *gorm.DB) port.IUnitPort {
	return &UnitRepositoryAdapter{base{db}}
}
