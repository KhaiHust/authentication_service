package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateGroupUsecase interface {
	CreateGroup(ctx context.Context, userID int64, req *request.CreateGroupDTO) (*entity.GroupEntity, error)
}

type CreateGroupUsecase struct {
	groupPort                  port.IGroupPort
	getGroupRoleUsecase        IGetGroupRoleUsecase
	databaseTransactionUsecase IDatabaseTransactionUsecase
	groupMemberPort            port.IGroupMemberPort
}

func (c *CreateGroupUsecase) CreateGroup(ctx context.Context, userID int64, req *request.CreateGroupDTO) (*entity.GroupEntity, error) {
	// get OWNER role
	role, err := c.getGroupRoleUsecase.GetRoleByCode(ctx, common.GROUP_OWNER_CODE)
	if err != nil {
		return nil, err
	}
	//start transaction
	tx := c.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			log.Error(ctx, "Panic error: %v", r)
			err = exception.InternalServerErrorException
		}
		if errRollback := c.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback create group error: %v", errRollback)
		} else {
			log.Info(ctx, "Rollback create group success")
		}
	}()
	// create group
	groupEntity := &entity.GroupEntity{
		Name:        req.Name,
		Description: req.Description,
	}
	groupEntity, err = c.groupPort.CreateGroup(ctx, tx, groupEntity)
	if err != nil {
		log.Error(ctx, "CreateGroup error: %v", err)
		return nil, err
	}
	groupMemberEntity := &entity.GroupMemberEntity{
		GroupID: groupEntity.ID,
		UserID:  userID,
		RoleID:  role.ID,
	}
	_, err = c.groupMemberPort.CreateGroupMember(ctx, tx, groupMemberEntity)
	if err != nil {
		log.Error(ctx, "CreateGroupMember error: %v", err)
		return nil, err
	}
	errCommitTxn := c.databaseTransactionUsecase.Commit(tx)
	if errCommitTxn != nil {
		log.Error(ctx, "Commit error: %v", errCommitTxn)
		return nil, errCommitTxn
	}
	return groupEntity, nil

}

func NewCreateGroupUsecase(groupPort port.IGroupPort,
	getGroupRoleUsecase IGetGroupRoleUsecase,
	databaseTransactionUsecase IDatabaseTransactionUsecase,
	groupMemberPort port.IGroupMemberPort) ICreateGroupUsecase {
	return &CreateGroupUsecase{
		groupPort:                  groupPort,
		getGroupRoleUsecase:        getGroupRoleUsecase,
		databaseTransactionUsecase: databaseTransactionUsecase,
		groupMemberPort:            groupMemberPort,
	}
}
