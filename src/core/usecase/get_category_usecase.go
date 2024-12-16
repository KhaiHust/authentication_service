package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetCategoryUsecase interface {
	GetAllCategory(ctx context.Context, spec *dto.CategorySpec) ([]*entity.CategoryEntity, error)
	CountAllCategory(ctx context.Context, spec *dto.CategorySpec) (int64, error)
}

type GetCategoryUsecase struct {
	categoryPort port.ICategoryPort
}

func (g GetCategoryUsecase) GetAllCategory(ctx context.Context, spec *dto.CategorySpec) ([]*entity.CategoryEntity, error) {
	result, err := g.categoryPort.GetAllCategory(ctx, spec)
	if err != nil {
		log.Error(ctx, "GetAllCategory", err)
		return nil, err
	}
	return result, nil
}

func (g GetCategoryUsecase) CountAllCategory(ctx context.Context, spec *dto.CategorySpec) (int64, error) {
	result, err := g.categoryPort.CountAllCategory(ctx, spec)
	if err != nil {
		log.Error(ctx, "CountAllCategory", err)
		return 0, err
	}
	return result, nil
}

func NewGetCategoryUsecase(categoryPort port.ICategoryPort) IGetCategoryUsecase {
	return &GetCategoryUsecase{
		categoryPort: categoryPort,
	}
}
