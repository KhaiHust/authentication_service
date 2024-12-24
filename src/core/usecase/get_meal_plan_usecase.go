package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
)

type IGetMealPlanUsecase interface {
	GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error)
}
type GetMealPlanUsecase struct {
	mealPlanPort port.IMealPlanPort
}

func (g GetMealPlanUsecase) GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error) {
	return g.mealPlanPort.GetMealPlanByUserIDAndID(ctx, userID, mealPlanID)
}

func NewGetMealPlanUsecase(mealPlanPort port.IMealPlanPort) IGetMealPlanUsecase {
	return &GetMealPlanUsecase{
		mealPlanPort: mealPlanPort,
	}
}
