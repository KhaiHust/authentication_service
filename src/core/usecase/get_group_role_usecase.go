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

type IGetGroupRoleUsecase interface {
	GetRoleByCode(ctx context.Context, code string) (*entity.GroupRoleEntity, error)
	GetRoleByIDs(ctx context.Context, ids []int64) ([]*entity.GroupRoleEntity, error)
}
type GetGroupRoleUsecase struct {
	groupRolePort port.IGroupRolePort
	cachePort     port.ICachePort
}

func (g GetGroupRoleUsecase) GetRoleByIDs(ctx context.Context, ids []int64) ([]*entity.GroupRoleEntity, error) {
	return g.groupRolePort.GetRoleByIDs(ctx, ids)
}

func (g GetGroupRoleUsecase) GetRoleByCode(ctx context.Context, code string) (*entity.GroupRoleEntity, error) {
	//get role from cache
	key := common.BuildCacheKeyGroupRoleCode(code)
	role, err := g.cachePort.GetFromCache(ctx, key)
	if err != nil && err.Error() != constant.ErrCacheKeyNil {

		log.Error(ctx, "GetFromCache error: %v", err)
		return nil, err
	}
	if role != nil {
		rsp := &entity.GroupRoleEntity{}
		err = json.Unmarshal([]byte(role.(string)), rsp)
		if err != nil {
			log.Error(ctx, "Unmarshal error: %v", err)
			return nil, err
		}
		return rsp, nil
	}
	roleEntity, err := g.groupRolePort.GetRoleByCode(ctx, code)
	if err != nil {
		log.Error(ctx, "GetRoleByCode error: %v", err)
		return nil, err
	}
	//set role to cache
	go func() {
		err = g.cachePort.SetToCache(ctx, key, roleEntity, constant.DefaultCacheTTL)
		if err != nil {
			log.Error(ctx, "SetToCache error: %v", err)
		}
	}()
	return roleEntity, nil
}

func NewGetGroupRoleUsecase(groupRolePort port.IGroupRolePort, cachePort port.ICachePort) IGetGroupRoleUsecase {
	return &GetGroupRoleUsecase{
		groupRolePort: groupRolePort,
		cachePort:     cachePort,
	}
}
