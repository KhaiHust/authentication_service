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

type IAddMemberGroupUsecase interface {
	AddNewMemberByEmail(ctx context.Context, userID int64, groupID int64, email string) (*entity.GroupMemberEntity, error)
}

type AddMemberGroupUsecase struct {
	getGroupUseCase            IGetGroupUseCase
	groupMemberPort            port.IGroupMemberPort
	getGroupRoleUsecase        IGetGroupRoleUsecase
	getUserUsecase             IGetUserUsecase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (a AddMemberGroupUsecase) AddNewMemberByEmail(ctx context.Context, userID int64, groupID int64, email string) (*entity.GroupMemberEntity, error) {
	//get group
	group, err := a.getGroupUseCase.GetGroupById(ctx, groupID)
	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: GetGroupById error", err)
		return nil, err
	}
	//get existed user
	existedUser, err := a.getUserUsecase.GetUserByEmail(&ctx, email)

	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: GetUserByEmail error", err)
		return nil, err
	}

	//get role
	role, err := a.getGroupRoleUsecase.GetRoleByCode(ctx, common.GROUP_OWNER_CODE)
	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: GetRoleByCode error", err)
		return nil, err
	}
	existedMember, err := a.groupMemberPort.GetGroupMemberByGroupIDAndUserID(ctx, group.ID, userID)
	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: GetGroupMemberByGroupIDAndUserID error", err)
		return nil, err
	}

	if existedMember.RoleID != role.ID {
		log.Error(ctx, "AddNewMemberByEmail: User is not owner")
		return nil, errors.New(constant.ErrForbiddenAddMember)
	}
	roleMember, err := a.getGroupRoleUsecase.GetRoleByCode(ctx, common.GROUP_MEMBER_CODE)
	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: GetRoleByCode error", err)
		return nil, err
	}
	//create group member
	groupMember := &entity.GroupMemberEntity{
		GroupID: group.ID,
		UserID:  existedUser.ID,
		RoleID:  roleMember.ID,
	}
	tx := a.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := a.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "AddNewMemberByEmail: RollbackTransaction error", errRollback)
		} else {
			log.Info(ctx, "AddNewMemberByEmail: RollbackTransaction success")
		}
	}()

	groupMember, err = a.groupMemberPort.CreateGroupMember(ctx, tx, groupMember)
	if err != nil {
		log.Error(ctx, "AddNewMemberByEmail: CreateGroupMember error", err)
		return nil, err
	}
	errCommit := a.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "AddNewMemberByEmail: CommitTransaction error", errCommit)
		return nil, errCommit
	}
	return groupMember, nil
}

func NewAddMemberGroupUsecase(getGroupUseCase IGetGroupUseCase, groupMemberPort port.IGroupMemberPort,
	getGroupRoleUsecase IGetGroupRoleUsecase, getUserUsecase IGetUserUsecase,
	databaseTransactionUsecase IDatabaseTransactionUsecase) IAddMemberGroupUsecase {
	return &AddMemberGroupUsecase{getGroupUseCase: getGroupUseCase, groupMemberPort: groupMemberPort,
		getGroupRoleUsecase: getGroupRoleUsecase, getUserUsecase: getUserUsecase,
		databaseTransactionUsecase: databaseTransactionUsecase}
}
