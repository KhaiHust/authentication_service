package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"gorm.io/gorm"
)

type IFoodPort interface {
	SaveFood(ctx context.Context, tx *gorm.DB, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
	GetFoodByUserIDAndID(ctx context.Context, userID, foodID int64) (*entity.FoodEntity, error)
	UpdateFood(ctx context.Context, tx *gorm.DB, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
	DeleteFood(ctx context.Context, tx *gorm.DB, foodID int64) error
	GetAllFood(ctx context.Context, foodParams *dto.FoodParams) ([]*entity.FoodEntity, error)
	CountAllFood(ctx context.Context, foodParams *dto.FoodParams) (int64, error)
	GetFoodByIDs(ctx context.Context, foodIDs []int64) ([]*entity.FoodEntity, error)
}
