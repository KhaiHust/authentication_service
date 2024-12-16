package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/response"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/request"
)

type IUserService interface {
	CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.UserEntity, error)
	LoginUser(ctx context.Context, email, password string) (*response.LoginResponseDto, error)
}
type UserService struct {
	createUserUsecase usecase.ICreateUserUsecase
	loginUserUseCase  usecase.ILoginUserUseCase
}

func (u *UserService) LoginUser(ctx context.Context, email, password string) (*response.LoginResponseDto, error) {
	return u.loginUserUseCase.LoginUser(&ctx, email, password)
}

func (u *UserService) CreateUser(ctx context.Context, req *request.RegisterUserRequest) (*entity.UserEntity, error) {
	userEntity := req.ToEntity()
	return u.createUserUsecase.CreateNewUser(&ctx, userEntity)
}

func NewUserService(createUserUsecase usecase.ICreateUserUsecase, loginUserUseCase usecase.ILoginUserUseCase) IUserService {
	return &UserService{
		createUserUsecase: createUserUsecase,
		loginUserUseCase:  loginUserUseCase,
	}
}
