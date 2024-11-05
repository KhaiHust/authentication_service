package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	exception2 "github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateUserUsecase interface {
	CreateNewUser(ctx *context.Context, userEntity *entity.UserEntity) (*entity.UserEntity, error)
}
type CreateUserUsecase struct {
	dbTxUsecase IDatabaseTransactionUsecase
	userPort    port.IUserPort
}

func (c *CreateUserUsecase) CreateNewUser(ctx *context.Context, userEntity *entity.UserEntity) (*entity.UserEntity, error) {
	existedUser, err := c.userPort.GetUserByEmail(ctx, userEntity.Email)
	if err != nil && err.Error() != constant.ErrUserNotFound {
		log.Error(ctx, "GetUserByEmail error: ", err)
		return nil, err
	}
	if existedUser != nil {
		log.Info(ctx, "User with email %s already existed", userEntity.Email)
		return nil, errors.New(constant.ErrExistedEmail)
	}
	hashPassword, err := common.HashPassword(userEntity.Password)
	if err != nil {
		log.Error(ctx, "HashPassword error: %v", err)
		return nil, err
	}
	userEntity.Password = hashPassword
	userEntity.IsActive = true
	tx := c.dbTxUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			log.Error(ctx, "Panic error: %v", r)
			err = exception2.InternalServerErrorException
		}
		if errRollback := c.dbTxUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback save user error: %v", errRollback)
		} else {
			log.Info(ctx, "Rollback save user success")
		}
	}()
	userEntity, err = c.userPort.SaveUser(ctx, userEntity, tx)
	if err != nil {
		log.Error(ctx, "SaveUser error: %v", err)
		return nil, err
	}
	errCommitTxn := c.dbTxUsecase.Commit(tx)
	if errCommitTxn != nil {
		log.Error(ctx, "Commit error: %v", errCommitTxn)
		return nil, errCommitTxn
	}
	userEntity.Password = ""
	return userEntity, nil
}

func NewCreateUserUsecase(dbTxUsecase IDatabaseTransactionUsecase, userPort port.IUserPort) ICreateUserUsecase {
	return &CreateUserUsecase{
		dbTxUsecase: dbTxUsecase,
		userPort:    userPort,
	}
}
