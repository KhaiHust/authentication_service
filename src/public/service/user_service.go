package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/request"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.UserEntity, error)
}
type UserService struct {
	createUserUsecase usecase.ICreateUserUsecase
}

func (u *UserService) CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.UserEntity, error) {
	userEntity := req.ToEntity()
	return u.createUserUsecase.CreateNewUser(&ctx, userEntity)
}

func NewUserService(createUserUsecase usecase.ICreateUserUsecase) IUserService {
	return &UserService{
		createUserUsecase: createUserUsecase,
	}
}
