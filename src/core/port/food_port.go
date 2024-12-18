package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IFoodPort interface {
	SaveFood(ctx context.Context, tx *gorm.DB, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
}
