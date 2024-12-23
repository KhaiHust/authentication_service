package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
)

type ICategoryPort interface {
	GetAllCategory(ctx context.Context, spec *dto.CategorySpec) ([]*entity.CategoryEntity, error)
	CountAllCategory(ctx context.Context, spec *dto.CategorySpec) (int64, error)
}
