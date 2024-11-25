package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	request2 "github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IGroupService interface {
	CreateGroup(ctx context.Context, userID int64, dto *request2.CreateGroupDTO) (*entity.GroupEntity, error)
}
type GroupService struct {
	createGroupUsecase usecase.ICreateGroupUsecase
}

func (g GroupService) CreateGroup(ctx context.Context, userID int64, dto *request2.CreateGroupDTO) (*entity.GroupEntity, error) {
	return g.createGroupUsecase.CreateGroup(ctx, userID, dto)
}

func NewGroupService(createGroupUsecase usecase.ICreateGroupUsecase) IGroupService {
	return &GroupService{createGroupUsecase: createGroupUsecase}
}
