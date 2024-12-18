package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetFoodUseCase interface {
	GetFoodByUserIDAndID(ctx context.Context, userID, foodID int64) (*entity.FoodEntity, error)
}
type GetFoodUseCase struct {
	foodPort port.IFoodPort
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
