package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreatGroupMemberUseCase interface {
	CreateGroupMember(ctx context.Context, groupID, userID int64, newUserEmail string) error
}
type CreateGroupMemberUseCase struct {
	groupMemberPort            port.IGroupMemberPort
	getGroupRoleUsecase        IGetGroupRoleUsecase
	getGroupUseCase            IGetGroupUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
	getUserUsecase             IGetUserUsecase
}

func (c CreateGroupMemberUseCase) CreateGroupMember(ctx context.Context, groupID, userID int64, newUserEmail string) error {
	group, err := c.getGroupUseCase.GetGroupById(ctx, groupID)
	if err != nil {
		log.Error("Get group by id error: ", err)
		return err
	}
	//check group and user existed
	_, err = c.groupMemberPort.GetGroupMemberByGroupIdAndUserId(ctx, group.ID, userID)
	if err != nil {
		log.Error("Get group member by group id and user id error: ", err)
		return err
	}
	//check existed user
	existedUser, err := c.getUserUsecase.GetUserByEmail(&ctx, newUserEmail)
	if err != nil {
		log.Error("Get user by email error: ", err)
		return err
	}
	role, err := c.getGroupRoleUsecase.GetRoleByCode(ctx, common.GROUP_MEMBER_CODE)
	groupMemberEntity := &entity.GroupMemberEntity{
		GroupID: group.ID,
		UserID:  existedUser.ID,
		RoleID:  role.ID,
	}
	tx := c.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}

		if errRollback := c.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error("Rollback error: ", errRollback)
		} else {
			log.Info("Rollback successfully")
		}
	}()
	_, err = c.groupMemberPort.CreateGroupMember(ctx, tx, groupMemberEntity)
	if err != nil {
		log.Error("Create group member error: ", err)
		return err
	}
	errCommit := c.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error("Commit error: ", errCommit)
		return errCommit
	}
	return nil
}

func NewCreateGroupMemberUseCase(groupMemberPort port.IGroupMemberPort, getGroupRoleUsecase IGetGroupRoleUsecase, getGroupUseCase IGetGroupUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreatGroupMemberUseCase {
	return &CreateGroupMemberUseCase{
		groupMemberPort:            groupMemberPort,
		getGroupRoleUsecase:        getGroupRoleUsecase,
		getGroupUseCase:            getGroupUseCase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
