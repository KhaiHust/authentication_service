package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IDeleteShoppingListUseCase interface {
	DeleteShoppingListByID(ctx context.Context, userID int64, shoppingListID int64) error
}
type DeleteShoppingListUseCase struct {
	databaseTransactionUsecase IDatabaseTransactionUsecase
	getGroupUseCase            IGetGroupUseCase
	getGroupMemberUseCase      IGetGroupMemberUseCase
	shoppingListPort           port.IShoppingListPort
}

func (d DeleteShoppingListUseCase) DeleteShoppingListByID(ctx context.Context, userID int64, shoppingListID int64) error {
	shoppingList, err := d.shoppingListPort.GetShoppingListByID(ctx, shoppingListID)
	if err != nil {
		log.Error(ctx, "Get shopping list by id error: %v", err)
		return err
	}
	if shoppingList.GroupID != nil {
		group, err := d.getGroupUseCase.GetGroupById(ctx, *shoppingList.GroupID)
		if err != nil {
			log.Error(ctx, "Get group by id error: %v", err)
			return err
		}
		groupMembers, err := d.getGroupMemberUseCase.GetListMemberByGroupID(ctx, userID, group.ID)
		if err != nil {
			log.Error(ctx, "Get list member by group id error: %v", err)
			return err
		}
		hasPermission := false
		for _, member := range groupMembers {
			if member.UserID == userID {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			log.Error(ctx, "User %d has no permission to delete shopping list for group %d", userID, *shoppingList.GroupID)
			return errors.New(constant.ErrForbiddenDeleteShoppingList)
		}
	}
	tx := d.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := d.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback transaction error: %v", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction success")
		}
	}()
	if err := d.shoppingListPort.DeleteShoppingListByID(ctx, tx, shoppingListID); err != nil {
		log.Error(ctx, "Delete shopping list by id error: %v", err)
		return err
	}
	errCommit := d.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "Commit transaction error: %v", errCommit)
		return errCommit
	}
	return nil
}

func NewDeleteShoppingListUseCase(databaseTransactionUsecase IDatabaseTransactionUsecase, getGroupUseCase IGetGroupUseCase, getGroupMemberUseCase IGetGroupMemberUseCase, shoppingListPort port.IShoppingListPort) IDeleteShoppingListUseCase {
	return &DeleteShoppingListUseCase{
		databaseTransactionUsecase: databaseTransactionUsecase,
		getGroupUseCase:            getGroupUseCase,
		getGroupMemberUseCase:      getGroupMemberUseCase,
		shoppingListPort:           shoppingListPort,
	}
}
