package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/helper"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/golibs-starter/golib/log"
)

type ICategoryService interface {
	GetAllCategory(ctx context.Context, spec *dto.CategorySpec) (*response.ListCategoryResponse, error)
	CreateNewCategory(ctx context.Context, cate *entity.CategoryEntity) (*response.CategoryResponse, error)
}
type CategoryService struct {
	getCategoryUsecase    usecase.IGetCategoryUsecase
	createCategoryUsecase usecase.ICreateCategoryUsecase
}

func (c CategoryService) CreateNewCategory(ctx context.Context, cate *entity.CategoryEntity) (*response.CategoryResponse, error) {
	category, err := c.createCategoryUsecase.CreateCategory(ctx, cate)
	if err != nil {
		log.Error(ctx, "Create new category failed", err)
		return nil, err
	}
	return response.ToCategoryResponse(category), nil
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

	return response.ToListCategoryResponse(categories, page, pageSize, totalPage, total, prePage, nextPage), nil

}

func NewCategoryService(getCategoryUsecase usecase.IGetCategoryUsecase, createCategoryUsecase usecase.ICreateCategoryUsecase) ICategoryService {
	return &CategoryService{
		getCategoryUsecase:    getCategoryUsecase,
		createCategoryUsecase: createCategoryUsecase,
	}
}
