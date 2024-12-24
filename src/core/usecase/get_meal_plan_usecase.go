package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
	"time"
)

type IGetMealPlanUsecase interface {
	GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error)
	GetMealPlanByDate(ctx context.Context, userID int64, date int64) ([]*entity.MealPlanEntity, error)
	GetDetailsMealPlan(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error)
}
type GetMealPlanUsecase struct {
	mealPlanPort     port.IMealPlanPort
	mealPlanFoodPort port.IMealPlanFoodPort
	getFoodUseCase   IGetFoodUseCase
}

func (g GetMealPlanUsecase) GetDetailsMealPlan(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error) {
	mealPlan, err := g.GetMealPlanByUserIDAndID(ctx, userID, mealPlanID)
	if err != nil {
		log.Error(ctx, "Get meal plan failed: ", err)
		return nil, err
	}
	mPFEntities, err := g.mealPlanFoodPort.GetMealPlanFoodByMealPlanID(ctx, mealPlanID)
	if err != nil {
		log.Error(ctx, "Get meal plan food failed: ", err)
		return nil, err
	}
	foodIDs := make([]int64, 0)
	for _, mpf := range mPFEntities {
		foodIDs = append(foodIDs, mpf.FoodID)
	}
	foods, err := g.getFoodUseCase.GetFoodByIDs(ctx, foodIDs)
	if err != nil {
		log.Error(ctx, "Get food failed: ", err)
		return nil, err
	}
	mealPlan.Foods = foods
	return mealPlan, nil
}

func (g GetMealPlanUsecase) GetMealPlanByDate(ctx context.Context, userID int64, date int64) ([]*entity.MealPlanEntity, error) {
	dateTime := time.Unix(date, 0)
	startOfDay := time.Date(dateTime.Year(), dateTime.Month(), dateTime.Day(), 0, 0, 0, 0, time.UTC)
	endOfDay := startOfDay.Add(24*time.Hour - time.Nanosecond)
	start := startOfDay.Unix()
	end := endOfDay.Unix()
	params := &dto.MealPlanParams{
		UserID:      &userID,
		StartedDate: &start,
		EndedDate:   &end,
	}
	return g.mealPlanPort.GetMealPlan(ctx, params)
}

func (g GetMealPlanUsecase) GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error) {
	return g.mealPlanPort.GetMealPlanByUserIDAndID(ctx, userID, mealPlanID)
}

func NewGetMealPlanUsecase(mealPlanPort port.IMealPlanPort) IGetMealPlanUsecase {
	return &GetMealPlanUsecase{
		mealPlanPort: mealPlanPort,
	}
}
