package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/helper"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/golibs-starter/golib/log"
)

type ICategoryService interface {
	GetAllCategory(ctx context.Context, spec *dto.CategorySpec) (*response.ListCategoryResponse, error)
}
type CategoryService struct {
	getCategoryUsecase usecase.IGetCategoryUsecase
}

func (c CategoryService) GetAllCategory(ctx context.Context, spec *dto.CategorySpec) (*response.ListCategoryResponse, error) {
	categories, err := c.getCategoryUsecase.GetAllCategory(ctx, spec)
	if err != nil {
		log.Error(ctx, "Get all category failed", err)
		return nil, err
	}
	total, err := c.getCategoryUsecase.CountAllCategory(ctx, spec)
	if err != nil {
		log.Error(ctx, "Count all category failed", err)
		return nil, err
	}
	page := int64(*spec.Page)
	pageSize := int64(*spec.PageSize)
	nextPage, prePage, totalPage := helper.CalculateParameterForGetRequest(page, pageSize, total)

	return response.ToListCategoryResponse(categories, page, pageSize, total, *nextPage, prePage, &totalPage), nil

}

func NewCategoryService(getCategoryUsecase usecase.IGetCategoryUsecase) ICategoryService {
	return &CategoryService{
		getCategoryUsecase: getCategoryUsecase,
	}
}
