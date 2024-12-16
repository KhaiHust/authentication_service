package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IUserProfilePort interface {
	GetUserProfilesByUserIDs(ctx context.Context, userIDs []int64) ([]*entity.UserProfileEntity, error)
	CreateNewProfile(ctx context.Context, txn *gorm.DB, userProfile *entity.UserProfileEntity) (*entity.UserProfileEntity, error)
}
