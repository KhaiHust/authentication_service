package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IMealPlanPort interface {
	SaveNewMealPlan(ctx context.Context, tx *gorm.DB, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
	GetMealPlanByUserIDAndID(ctx context.Context, userID, mealPlanID int64) (*entity.MealPlanEntity, error)
	UpdateMealPlan(ctx context.Context, tx *gorm.DB, mealPlanID int64, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
	DeleteMealPlanByID(ctx context.Context, tx *gorm.DB, mealPlanID int64) error
}
