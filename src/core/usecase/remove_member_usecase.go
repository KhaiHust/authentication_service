package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IRemoveMemberUsecase interface {
	RemoveMemberByUserID(ctx context.Context, userID, groupID, removeUserID int64) error
}

type RemoveMemberUsecase struct {
	getGroupUseCase            IGetGroupUseCase
	getGroupRoleUsecase        IGetGroupRoleUsecase
	groupMemberPort            port.IGroupMemberPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (r RemoveMemberUsecase) RemoveMemberByUserID(ctx context.Context, userID, groupID, removeUserID int64) error {
	group, err := r.getGroupUseCase.GetGroupById(ctx, groupID)
	if err != nil {
		log.Error(ctx, "Get group by id error: %v", err)
		return err
	}
	role, err := r.getGroupRoleUsecase.GetRoleByCode(ctx, common.GROUP_OWNER_CODE)
	if err != nil {
		log.Error(ctx, "Get role by code error: %v", err)
		return err
	}
	userIDs := []int64{userID, removeUserID}
	groupMembers, err := r.groupMemberPort.GetMembersByGroupAndUserIDs(ctx, group.ID, userIDs)
	if err != nil {
		log.Error(ctx, "Get members by group and user ids error: %v", err)
		return err
	}
	mapMembers := make(map[int64]*entity.GroupMemberEntity)
	for _, member := range groupMembers {
		mapMembers[member.UserID] = member
	}
	if mapMembers[userID].RoleID != role.ID {
		log.Error(ctx, "User %d has no permission to remove member %d", userID, removeUserID)
		return errors.New(constant.ErrForbiddenRemoveMember)
	}
	tx := r.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := r.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback transaction error: %v", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction successfully")
		}
	}()
	if err = r.groupMemberPort.DeleteMemberByGroupIDAndUserID(ctx, tx, group.ID, removeUserID); err != nil {
		log.Error(ctx, "Delete member by group id and user id error: %v", err)
		return err
	}
	if errCommit := r.databaseTransactionUsecase.Commit(tx); errCommit != nil {
		log.Error(ctx, "Commit transaction error: %v", errCommit)
		return errCommit
	}
	return nil
}

func NewRemoveMemberUsecase(getGroupUseCase IGetGroupUseCase, getGroupRoleUsecase IGetGroupRoleUsecase, groupMemberPort port.IGroupMemberPort, databaseTransactionUsecase IDatabaseTransactionUsecase) IRemoveMemberUsecase {
	return &RemoveMemberUsecase{getGroupUseCase: getGroupUseCase, getGroupRoleUsecase: getGroupRoleUsecase, groupMemberPort: groupMemberPort, databaseTransactionUsecase: databaseTransactionUsecase}
}
