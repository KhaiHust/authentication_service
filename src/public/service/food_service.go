package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/helper"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/response"
)

type IFoodService interface {
	CreateNewFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
	UpdateFood(ctx context.Context, userID, foodID int64, updateFoodDto *request.UpdateFoodDto) (*entity.FoodEntity, error)
	DeleteFood(ctx context.Context, userID, foodID int64) error
	GetAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) (*response.GetFoodResponse, error)
}

type FoodService struct {
	createFoodUsecase usecase.ICreateFoodUsecase
	updateFoodUseCase usecase.IUpdateFoodUseCase
	deleteFoodUseCase usecase.IDeleteFoodUseCase
	getFoodUseCase    usecase.IGetFoodUseCase
}

func (f FoodService) GetAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) (*response.GetFoodResponse, error) {
	foods, err := f.getFoodUseCase.GetAllFood(ctx, userID, foodParams)
	if err != nil {
		return nil, err
	}
	total, err := f.getFoodUseCase.CountAllFood(ctx, userID, foodParams)
	if err != nil {
		return nil, err
	}
	page := int64(*foodParams.Page)
	pageSize := int64(*foodParams.PageSize)
	nextPage, prePage, totalPage := helper.CalculateParameterForGetRequest(page, pageSize, total)
	return response.ToGetFoodResponse(foods, page, pageSize, totalPage, total, prePage, nextPage), nil
}

func (f FoodService) DeleteFood(ctx context.Context, userID, foodID int64) error {
	return f.deleteFoodUseCase.DeleteFood(ctx, userID, foodID)
}

func (f FoodService) UpdateFood(ctx context.Context, userID, foodID int64, updateFoodDto *request.UpdateFoodDto) (*entity.FoodEntity, error) {
	return f.updateFoodUseCase.UpdateFood(ctx, userID, foodID, updateFoodDto)
}

func (f FoodService) CreateNewFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error) {
	return f.createFoodUsecase.CreateFood(ctx, foodEntity)
}

func NewFoodService(createFoodUsecase usecase.ICreateFoodUsecase, updateFoodUseCase usecase.IUpdateFoodUseCase, deleteFoodUseCase usecase.IDeleteFoodUseCase, getFoodUseCase usecase.IGetFoodUseCase) IFoodService {
	return &FoodService{
		createFoodUsecase: createFoodUsecase,
		updateFoodUseCase: updateFoodUseCase,
		deleteFoodUseCase: deleteFoodUseCase,
		getFoodUseCase:    getFoodUseCase,
	}
}
