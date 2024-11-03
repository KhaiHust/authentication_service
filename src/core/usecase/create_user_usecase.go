package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
)

type ICreateUserUsecase interface {
	CreateNewUser(ctx *context.Context, userEntity *entity.UserEntity) (*entity.UserEntity, error)
}
type CreateUserUsecase struct {
	dbTxUsecase IDatabaseTransactionUsecase
	userPort    port.IUserPort
}

func (c *CreateUserUsecase) CreateNewUser(ctx *context.Context, userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	return nil, nil
}

func NewCreateUserUsecase(dbTxUsecase IDatabaseTransactionUsecase, userPort port.IUserPort) ICreateUserUsecase {
	return &CreateUserUsecase{
		dbTxUsecase: dbTxUsecase,
		userPort:    userPort,
	}
}
