package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
)

type IUserProfilePort interface {
	GetUserProfilesByUserIDs(ctx context.Context, userIDs []int64) ([]*entity.UserProfileEntity, error)
}
