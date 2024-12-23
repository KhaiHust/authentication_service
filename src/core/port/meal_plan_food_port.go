package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IMealPlanFoodPort interface {
	SaveMealPlanFood(ctx context.Context, tx *gorm.DB, mpFEntity *entity.MealPlanFoodEntity) (*entity.MealPlanFoodEntity, error)
	SaveListMealPlanFood(ctx context.Context, tx *gorm.DB, mpFEntities []*entity.MealPlanFoodEntity) ([]*entity.MealPlanFoodEntity, error)
}
