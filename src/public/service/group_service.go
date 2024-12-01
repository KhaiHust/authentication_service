package service

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	request2 "github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/golibs-starter/golib/log"
)

type IGroupService interface {
	CreateGroup(ctx context.Context, userID int64, dto *request2.CreateGroupDTO) (*entity.GroupEntity, error)
	AddNewMember(ctx context.Context, userID, groupID int64, email string) (*entity.GroupMemberEntity, error)
	GetMembers(ctx context.Context, userID, groupID int64) ([]*entity.GroupMemberEntity, error)
	RemoveMember(ctx context.Context, userID, groupID, removeUserID int64) error
}
type GroupService struct {
	createGroupUsecase    usecase.ICreateGroupUsecase
	addMemberGroupUsecase usecase.IAddMemberGroupUsecase
	getGroupMemberUseCase usecase.IGetGroupMemberUseCase
	removeMemberUsecase   usecase.IRemoveMemberUsecase
}

func (g GroupService) RemoveMember(ctx context.Context, userID, groupID, removeUserID int64) error {
	if userID == removeUserID {
		log.Error(ctx, "User %d can not remove yourself", userID)
		return errors.New(constant.ErrForbiddenRemoveMember)
	}
	return g.removeMemberUsecase.RemoveMemberByUserID(ctx, userID, groupID, removeUserID)
}

func (g GroupService) GetMembers(ctx context.Context, userID, groupID int64) ([]*entity.GroupMemberEntity, error) {
	return g.getGroupMemberUseCase.GetListMemberByGroupID(ctx, userID, groupID)
}

func (g GroupService) AddNewMember(ctx context.Context, userID, groupID int64, email string) (*entity.GroupMemberEntity, error) {
	return g.addMemberGroupUsecase.AddNewMemberByEmail(ctx, userID, groupID, email)
}

func (g GroupService) CreateGroup(ctx context.Context, userID int64, dto *request2.CreateGroupDTO) (*entity.GroupEntity, error) {
	return g.createGroupUsecase.CreateGroup(ctx, userID, dto)
}

func NewGroupService(createGroupUsecase usecase.ICreateGroupUsecase, addMemberGroupUsecase usecase.IAddMemberGroupUsecase, getGroupMemberUseCase usecase.IGetGroupMemberUseCase, removeMemberUsecase usecase.IRemoveMemberUsecase) IGroupService {
	return &GroupService{createGroupUsecase: createGroupUsecase, addMemberGroupUsecase: addMemberGroupUsecase, getGroupMemberUseCase: getGroupMemberUseCase, removeMemberUsecase: removeMemberUsecase}
}
