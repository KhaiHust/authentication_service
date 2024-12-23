package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IMealPlanPort interface {
	SaveNewMealPlan(ctx context.Context, tx *gorm.DB, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
}
