package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IFridgeItemPort interface {
	SaveFridgeItem(ctx context.Context, tx *gorm.DB, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error)
	GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error)
	UpdateItem(ctx context.Context, tx *gorm.DB, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error)
	DeleteItem(ctx context.Context, tx *gorm.DB, itemID int64) error
	GetAllItems(ctx context.Context, userID int64) ([]*entity.FridgeItemEntity, error)
}
