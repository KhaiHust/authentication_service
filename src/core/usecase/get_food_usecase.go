package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetFoodUseCase interface {
	GetFoodByUserIDAndID(ctx context.Context, userID, foodID int64) (*entity.FoodEntity, error)
	GetAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) ([]*entity.FoodEntity, error)
	CountAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) (int64, error)
	GetFoodByIDs(ctx context.Context, foodIDs []int64) ([]*entity.FoodEntity, error)
}
type GetFoodUseCase struct {
	foodPort port.IFoodPort
}

func (g GetFoodUseCase) GetFoodByIDs(ctx context.Context, foodIDs []int64) ([]*entity.FoodEntity, error) {
	foods, err := g.foodPort.GetFoodByIDs(ctx, foodIDs)
	if err != nil {
		log.Error(ctx, "Get food by IDs failed: ", err)
		return nil, err
	}
	return foods, nil
}

func (g GetFoodUseCase) GetAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) ([]*entity.FoodEntity, error) {
	foodParams.UserID = &userID
	foods, err := g.foodPort.GetAllFood(ctx, foodParams)
	if err != nil {
		log.Error(ctx, "Get all food failed: ", err)
		return nil, err
	}
	return foods, nil
}

func (g GetFoodUseCase) CountAllFood(ctx context.Context, userID int64, foodParams *dto.FoodParams) (int64, error) {
	foodParams.UserID = &userID
	count, err := g.foodPort.CountAllFood(ctx, foodParams)
	if err != nil {
		log.Error(ctx, "Count all food failed: ", err)
		return 0, err
	}
	return count, nil
}

func (g GetFoodUseCase) GetFoodByUserIDAndID(ctx context.Context, userID, foodID int64) (*entity.FoodEntity, error) {
	food, err := g.foodPort.GetFoodByUserIDAndID(ctx, userID, foodID)
	if err != nil {
		log.Error(ctx, "Get food failed: ", err)
		return nil, err
	}
	return food, nil
}

func NewGetFoodUseCase(foodPort port.IFoodPort) IGetFoodUseCase {
	return &GetFoodUseCase{
		foodPort: foodPort,
	}
}
