package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
)

type IGetGroupUseCase interface {
	GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error)
}
type GetGroupUseCase struct {
	groupPort port.IGroupPort
}

func (g GetGroupUseCase) GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error) {
	return g.groupPort.GetGroupById(ctx, groupID)
}

func NewGetGroupUseCase(groupPort port.IGroupPort) IGetGroupUseCase {
	return &GetGroupUseCase{
		groupPort: groupPort,
	}
}
