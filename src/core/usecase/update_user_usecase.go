package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type IUpdateUserUseCase interface {
	UpdateVerifiedUser(ctx *context.Context, userEntity *entity.UserEntity, tx *gorm.DB) (*entity.UserEntity, error)
}
type UpdateUserUseCase struct {
	userPort port.IUserPort
}

func (u *UpdateUserUseCase) UpdateVerifiedUser(ctx *context.Context, userEntity *entity.UserEntity, tx *gorm.DB) (*entity.UserEntity, error) {
	return u.userPort.UpdateUser(ctx, userEntity, tx)
}

func NewUpdateUserUseCase(userPort port.IUserPort) IUpdateUserUseCase {
	return &UpdateUserUseCase{userPort: userPort}
}
