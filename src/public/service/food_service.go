package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IFoodService interface {
	CreateNewFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
	UpdateFood(ctx context.Context, userID, foodID int64, updateFoodDto *request.UpdateFoodDto) (*entity.FoodEntity, error)
	DeleteFood(ctx context.Context, userID, foodID int64) error
}

type FoodService struct {
	createFoodUsecase usecase.ICreateFoodUsecase
	updateFoodUseCase usecase.IUpdateFoodUseCase
	deleteFoodUseCase usecase.IDeleteFoodUseCase
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

func NewFoodService(createFoodUsecase usecase.ICreateFoodUsecase, updateFoodUseCase usecase.IUpdateFoodUseCase, deleteFoodUseCase usecase.IDeleteFoodUseCase) IFoodService {
	return &FoodService{
		createFoodUsecase: createFoodUsecase,
		updateFoodUseCase: updateFoodUseCase,
		deleteFoodUseCase: deleteFoodUseCase,
	}
}
