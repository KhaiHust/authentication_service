package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"gorm.io/gorm"
)

type IUnitPort interface {
	GetAllUnit(ctx context.Context, spec *dto.UnitParamDto) ([]*entity.UnitEntity, error)
	CountAllUnit(ctx context.Context, spec *dto.UnitParamDto) (int64, error)
	SaveUnit(ctx context.Context, tx *gorm.DB, unit *entity.UnitEntity) (*entity.UnitEntity, error)
}
