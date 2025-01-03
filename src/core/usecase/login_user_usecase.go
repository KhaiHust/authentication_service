package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/response"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/KhaiHust/authen_service/core/properties"
	"github.com/golibs-starter/golib/log"
	context2 "github.com/golibs-starter/golib/web/context"
)

type ILoginUserUseCase interface {
	LoginUser(ctx *context.Context, email string, password string) (*response.LoginResponseDto, error)
}
type LoginUserUseCase struct {
	jwtProps                   *properties.TokenProperties
	refreshTokenPort           port.IRefreshTokenPort
	getUserUsecase             IGetUserUsecase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (l *LoginUserUseCase) LoginUser(ctx *context.Context, email string, password string) (*response.LoginResponseDto, error) {
	existedUser, err := l.getUserUsecase.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "Get user by email error: %v", err)
		return nil, err
	}
	ok := common.ComparePassword(existedUser.Password, password)
	if !ok {
		log.Error(ctx, "Password is invalid, email: %s", email)
		return nil, errors.New(constant.ErrWrongPassword)
	}
	var logRsp response.LoginResponseDto
	token, err := common.GenerateToken(existedUser, l.jwtProps)
	if err != nil {
		log.Error(ctx, "Generate token error: %v", err)
		return nil, err
	}
	logRsp.AccessToken = token
	refreshToken, err := common.GenerateRefreshToken(existedUser, l.jwtProps)
	if err != nil {
		log.Error(ctx, "Generate refresh token error: %v", err)
		return nil, err
	}
	logRsp.RefreshToken = refreshToken
	refreshTokenEntity := &entity.RefreshTokenEntity{
		UserId:       existedUser.ID,
		RefreshToken: refreshToken,
		IpAddress:    (*ctx).Value(constant.RequestAttributesKey).(*context2.RequestAttributes).ClientIpAddress,
		UserAgent:    (*ctx).Value(constant.RequestAttributesKey).(*context2.RequestAttributes).UserAgent,
	}
	tx := l.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if err != nil {
			errRollback := l.databaseTransactionUsecase.Rollback(tx)
			if errRollback != nil {
				log.Error(ctx, "Rollback transaction error: %v", errRollback)
			} else {
				log.Error(ctx, "Rollback transaction success")
			}
		}
	}()
	_, err = l.refreshTokenPort.SaveRefreshToken(ctx, refreshTokenEntity, tx)
	if err != nil {
		log.Error(ctx, "Save refresh token error: %v", err)
		return nil, err
	}
	errCommit := l.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "Commit transaction error: %v", errCommit)
		return nil, errCommit
	}
	return &logRsp, nil
}

func NewLoginUserUseCase(jwtProps *properties.TokenProperties, refreshTokenPort port.IRefreshTokenPort,
	getUserUsecase IGetUserUsecase,
	databaseTransactionUsecase IDatabaseTransactionUsecase) ILoginUserUseCase {
	return &LoginUserUseCase{jwtProps: jwtProps, refreshTokenPort: refreshTokenPort,
		getUserUsecase:             getUserUsecase,
		databaseTransactionUsecase: databaseTransactionUsecase}
}
