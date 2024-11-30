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

type IGetGroupUseCase interface {
	GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error)
}
type GetGroupUseCase struct {
	groupPort port.IGroupPort
	cachePort port.ICachePort
}

func (g GetGroupUseCase) GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error) {
	//get from cache
	group, err := g.cachePort.GetFromCache(ctx, common.BuildCacheKeyGroup(groupID))
	if err != nil && err.Error() != constant.ErrCacheKeyNil {
		log.Error(ctx, "GetGroupById: GetFromCache error", err)
		return nil, err
	}
	if group != nil {
		rsp := &entity.GroupEntity{}
		err = json.Unmarshal([]byte(group.(string)), rsp)
		if err != nil {
			log.Error(ctx, "GetGroupById: Unmarshal error", err)
			return nil, err
		}
	}
	groupEntity, err := g.groupPort.GetGroupById(ctx, groupID)
	if err != nil {
		log.Error(ctx, "GetGroupById: GetGroupById error", err)
		return nil, err
	}
	//set to cache
	go func() {
		err = g.cachePort.SetToCache(ctx, common.BuildCacheKeyGroup(groupID), groupEntity, constant.DefaultCacheTTL)
		if err != nil {
			log.Error(ctx, "GetGroupById: SetToCache error", err)
		}
	}()
	return groupEntity, nil
}

func NewGetGroupUseCase(groupPort port.IGroupPort, cachePort port.ICachePort) IGetGroupUseCase {
	return &GetGroupUseCase{groupPort: groupPort, cachePort: cachePort}
}
