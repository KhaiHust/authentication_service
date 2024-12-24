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
	GetMealPlanByDate(ctx context.Context, userID int64, date int64) ([]*entity.MealPlanEntity, error)
	GetDetailMealPlan(ctx context.Context, userID, mealID int64) (*entity.MealPlanEntity, error)
}
type MealPlanService struct {
	createMealPlanUsecase usecase.ICreateMealPlanUsecase
	updateMealPlanUsecase usecase.IUpdateMealPlanUsecase
	deleteMealPlanUsecase usecase.IDeleteMealPlanUsecase
	getMealPlanUsecase    usecase.IGetMealPlanUsecase
}

func (m MealPlanService) GetDetailMealPlan(ctx context.Context, userID, mealID int64) (*entity.MealPlanEntity, error) {
	return m.getMealPlanUsecase.GetDetailsMealPlan(ctx, userID, mealID)
}

func (m MealPlanService) GetMealPlanByDate(ctx context.Context, userID int64, date int64) ([]*entity.MealPlanEntity, error) {
	return m.getMealPlanUsecase.GetMealPlanByDate(ctx, userID, date)
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

func NewMealPlanService(createMealPlanUsecase usecase.ICreateMealPlanUsecase, updateMealPlanUsecase usecase.IUpdateMealPlanUsecase, deleteMealPlanUsecase usecase.IDeleteMealPlanUsecase, getMealPlanUsecase usecase.IGetMealPlanUsecase) IMealPlanService {
	return &MealPlanService{
		createMealPlanUsecase: createMealPlanUsecase,
		updateMealPlanUsecase: updateMealPlanUsecase,
		deleteMealPlanUsecase: deleteMealPlanUsecase,
		getMealPlanUsecase:    getMealPlanUsecase,
	}
}
