package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IMealPlanService interface {
	CreateNewMealPlan(ctx context.Context, userID int64, ml *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
}
type MealPlanService struct {
	createMealPlanUsecase usecase.ICreateMealPlanUsecase
}

func (m MealPlanService) CreateNewMealPlan(ctx context.Context, userID int64, ml *entity.MealPlanEntity) (*entity.MealPlanEntity, error) {
	ml.UserID = userID
	return m.createMealPlanUsecase.CreateNewMealPlan(ctx, ml)
}

func NewMealPlanService(createMealPlanUsecase usecase.ICreateMealPlanUsecase) IMealPlanService {
	return &MealPlanService{createMealPlanUsecase: createMealPlanUsecase}
}
