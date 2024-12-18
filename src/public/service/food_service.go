package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IFoodService interface {
	CreateNewFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
}

type FoodService struct {
	createFoodUsecase usecase.ICreateFoodUsecase
}

func (f FoodService) CreateNewFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error) {
	return f.createFoodUsecase.CreateFood(ctx, foodEntity)
}

func NewFoodService(createFoodUsecase usecase.ICreateFoodUsecase) IFoodService {
	return &FoodService{createFoodUsecase: createFoodUsecase}
}
