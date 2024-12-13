package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetGroupMemberUseCase interface {
	GetListMemberByGroupID(ctx context.Context, userID, groupID int64) ([]*entity.GroupMemberEntity, error)
	IsMemberOfGroup(ctx context.Context, userID, groupID int64) (bool, error)
}
type GetGroupMemberUseCase struct {
	groupMemberPort     port.IGroupMemberPort
	getGroupUseCase     IGetGroupUseCase
	getGroupRoleUsecase IGetGroupRoleUsecase
	getUserProfilePort  port.IUserProfilePort
}

func (g GetGroupMemberUseCase) IsMemberOfGroup(ctx context.Context, userID, groupID int64) (bool, error) {
	groupMember, err := g.groupMemberPort.GetGroupMemberByGroupIDAndUserID(ctx, groupID, userID)
	if err != nil {
		log.Error(ctx, "Get group member by group id and user id error: %v", err)
		return false, err
	}
	if groupMember == nil {
		return false, nil
	}
	return true, nil
}

func (g GetGroupMemberUseCase) GetListMemberByGroupID(ctx context.Context, userID, groupID int64) ([]*entity.GroupMemberEntity, error) {
	group, err := g.getGroupUseCase.GetGroupById(ctx, groupID)
	if err != nil {
		log.Error(ctx, "Get group by id error: %v", err)
		return nil, err
	}
	groupMembers, err := g.groupMemberPort.GetListMemberByGroupID(ctx, group.ID)
	if err != nil {
		log.Error(ctx, "Get list member by group id error: %v", err)
		return nil, err
	}
	//check permission
	hasPermission := false
	for _, member := range groupMembers {
		if member.UserID == userID {
			hasPermission = true
			break
		}
	}
	if !hasPermission {
		log.Error(ctx, "User %d has no permission to get member of group %d", userID, groupID)
		return nil, errors.New(constant.ErrForbiddenGetMember)
	}
	// get user profiles
	userIDs := make([]int64, 0)
	for _, member := range groupMembers {
		userIDs = append(userIDs, member.UserID)
	}
	userProfiles, err := g.getUserProfilePort.GetUserProfilesByUserIDs(ctx, userIDs)
	if err != nil {
		log.Error(ctx, "Get user profiles by user ids error: %v", err)
		return nil, err
	}
	mapUserProfile := make(map[int64]*entity.UserProfileEntity)
	for _, userProfile := range userProfiles {
		mapUserProfile[userProfile.UserID] = userProfile
	}
	// get group roles
	roleIDs := make([]int64, 0)
	for _, member := range groupMembers {
		roleIDs = append(roleIDs, member.RoleID)
	}
	roles, err := g.getGroupRoleUsecase.GetRoleByIDs(ctx, roleIDs)
	if err != nil {
		log.Error(ctx, "Get roles by ids error: %v", err)
		return nil, err
	}
	mapRole := make(map[int64]*entity.GroupRoleEntity)
	for _, role := range roles {
		mapRole[role.ID] = role
	}
	for idx, member := range groupMembers {
		groupMembers[idx].UserProfile = mapUserProfile[member.UserID]
		groupMembers[idx].Role = mapRole[member.RoleID]
	}
	return groupMembers, nil
}

func NewGetGroupMemberUseCase(groupMemberPort port.IGroupMemberPort, getGroupUseCase IGetGroupUseCase, getGroupRoleUsecase IGetGroupRoleUsecase, getUserProfilePort port.IUserProfilePort) IGetGroupMemberUseCase {
	return &GetGroupMemberUseCase{
		groupMemberPort:     groupMemberPort,
		getGroupUseCase:     getGroupUseCase,
		getGroupRoleUsecase: getGroupRoleUsecase,
		getUserProfilePort:  getUserProfilePort,
	}
}
