package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IMealPlanService interface {
	CreateNewMealPlan(ctx context.Context, userID int64, ml *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
	UpdateMealPlan(ctx context.Context, userID, mealID int64, req *request.UpdateMealPlanDTO) (*entity.MealPlanEntity, error)
	DeleteMealPlan(ctx context.Context, userID, mealID int64) error
}
type MealPlanService struct {
	createMealPlanUsecase usecase.ICreateMealPlanUsecase
	updateMealPlanUsecase usecase.IUpdateMealPlanUsecase
	deleteMealPlanUsecase usecase.IDeleteMealPlanUsecase
}

func (m MealPlanService) DeleteMealPlan(ctx context.Context, userID, mealID int64) error {
	return m.deleteMealPlanUsecase.DeleteMealPlan(ctx, userID, mealID)
}

func (m MealPlanService) UpdateMealPlan(ctx context.Context, userID, mealID int64, req *request.UpdateMealPlanDTO) (*entity.MealPlanEntity, error) {
	return m.updateMealPlanUsecase.UpdateMealPlan(ctx, userID, mealID, req)
}

func (m MealPlanService) CreateNewMealPlan(ctx context.Context, userID int64, ml *entity.MealPlanEntity) (*entity.MealPlanEntity, error) {
	ml.UserID = userID
	return m.createMealPlanUsecase.CreateNewMealPlan(ctx, ml)
}

func NewMealPlanService(createMealPlanUsecase usecase.ICreateMealPlanUsecase, updateMealPlanUsecase usecase.IUpdateMealPlanUsecase) IMealPlanService {
	return &MealPlanService{
		createMealPlanUsecase: createMealPlanUsecase,
		updateMealPlanUsecase: updateMealPlanUsecase,
	}
}
