package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
)

type IGetUserUsecase interface {
	GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error)
}

type GetUserUsecase struct {
	userPort port.IUserPort
}

func (g *GetUserUsecase) GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error) {
	return g.userPort.GetUserByEmail(ctx, email)
}

func NewGetUserUsecase(userPort port.IUserPort) IGetUserUsecase {
	return &GetUserUsecase{
		userPort: userPort,
	}
}
