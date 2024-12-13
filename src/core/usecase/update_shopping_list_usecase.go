package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IUpdateShoppingListUseCase interface {
	UpdateShoppingListByID(ctx context.Context, userID int64, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error)
}

type UpdateShoppingListUseCase struct {
	shoppingListPort           port.IShoppingListPort
	shoppingListGroupPort      port.IShoppingListGroupPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
	getGroupUseCase            IGetGroupUseCase
	getGroupMemberUseCase      IGetGroupMemberUseCase
}

func (u UpdateShoppingListUseCase) UpdateShoppingListByID(ctx context.Context, userID int64, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error) {
	shoppingList, err := u.shoppingListPort.GetShoppingListByID(ctx, req.ID)
	if err != nil {
		log.Error(ctx, "Get shopping list by id error: %v", err)
		return nil, err
	}
	shoppingList.Name = req.Name
	shoppingList.Description = req.Description
	if shoppingList.GroupID != nil && req.AssignedTo > 0 && shoppingList.AssignedTo != req.AssignedTo {
		group, err := u.getGroupUseCase.GetGroupById(ctx, *shoppingList.GroupID)
		if err != nil {
			log.Error(ctx, "Get group by id error: %v", err)
			return nil, err
		}
		groupMembers, err := u.getGroupMemberUseCase.GetListMemberByGroupID(ctx, req.CreatedBy, group.ID)
		if err != nil {
			log.Error(ctx, "Get list member by group id error: %v", err)
			return nil, err
		}
		hasPermission := false

		for _, member := range groupMembers {
			if member.UserID == userID {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			log.Error(ctx, "User %d has no permission to update shopping list for group %d", userID, *shoppingList.GroupID)
			return nil, errors.New(constant.ErrForbiddenUpdateShoppingList)
		}
		for _, member := range groupMembers {
			if member.UserID == req.AssignedTo {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			log.Error(ctx, "User %d has no permission to update shopping list for group %d", req.CreatedBy, *shoppingList.GroupID)
			return nil, errors.New(constant.ErrForbiddenUpdateShoppingList)
		}

		shoppingList.AssignedTo = req.AssignedTo
	}
	tx := u.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction error: %v", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction successfully")
		}
	}()
	shoppingList, err = u.shoppingListPort.UpdateShoppingListByID(ctx, tx, shoppingList)
	if err != nil {
		log.Error(ctx, "Update shopping list by id error: %v", err)
		return nil, err
	}
	errCommit := u.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "Commit transaction error: %v", errCommit)
		return nil, errCommit
	}
	return shoppingList, nil
}

func NewUpdateShoppingListUseCase(shoppingListPort port.IShoppingListPort, shoppingListGroupPort port.IShoppingListGroupPort, databaseTransactionUsecase IDatabaseTransactionUsecase, getGroupUseCase IGetGroupUseCase, getGroupMemberUseCase IGetGroupMemberUseCase) IUpdateShoppingListUseCase {
	return UpdateShoppingListUseCase{
		shoppingListPort:           shoppingListPort,
		shoppingListGroupPort:      shoppingListGroupPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
		getGroupUseCase:            getGroupUseCase,
		getGroupMemberUseCase:      getGroupMemberUseCase,
	}
}
