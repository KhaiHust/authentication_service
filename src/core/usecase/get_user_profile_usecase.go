package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
)

type IGetUserProfileUseCase interface {
	GetUserProfileByUserID(ctx context.Context, userID int64) (*entity.UserProfileEntity, error)
}
type GetUserProfileUseCase struct {
	userProfilePort port.IUserProfilePort
}

func (u *GetUserProfileUseCase) GetUserProfileByUserID(ctx context.Context, userID int64) (*entity.UserProfileEntity, error) {
	return u.userProfilePort.GetUserProfileByUserID(ctx, userID)
}
func NewGetUserProfileUseCase(userProfilePort port.IUserProfilePort) IGetUserProfileUseCase {
	return &GetUserProfileUseCase{
		userProfilePort: userProfilePort,
	}
}
