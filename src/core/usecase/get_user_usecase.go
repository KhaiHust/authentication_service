package usecase

import (
	"context"
	"encoding/json"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetUserUsecase interface {
	GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error)
}

type GetUserUsecase struct {
	userPort  port.IUserPort
	cachePort port.ICachePort
}

func (g *GetUserUsecase) GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error) {
	// Get user from cache
	key := common.BuildCacheKeyGetUserInfoByEmail(email)
	user, err := g.cachePort.GetFromCache(*ctx, key)
	if err != nil && err.Error() != constant.ErrCacheKeyNil {
		log.Error(ctx, "GetUserByEmail: GetFromCache error", err)
		return nil, err
	}
	if user != nil {
		rsp := &entity.UserEntity{}
		err = json.Unmarshal([]byte(user.(string)), rsp)
		if err != nil {
			log.Error(ctx, "GetUserByEmail: Unmarshal error", err)
			return nil, err
		}
		return rsp, nil
	}
	userEntity, err := g.userPort.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "GetUserByEmail: GetUserByEmail error", err)
		return nil, err
	}
	// Set user to cache
	go func() {
		err = g.cachePort.SetToCache(*ctx, key, userEntity, constant.DefaultCacheTTL)
		if err != nil {
			log.Error(ctx, "GetUserByEmail: SetToCache error", err)
		}
	}()
	return userEntity, nil
}

func NewGetUserUsecase(userPort port.IUserPort, cache port.ICachePort) IGetUserUsecase {
	return &GetUserUsecase{
		userPort:  userPort,
		cachePort: cache,
	}
}
