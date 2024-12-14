package usecase

import (
	"context"
	"fmt"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateShoppingTaskUsecase interface {
	CreateNewShoppingTask(ctx context.Context, userId, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error)
}
type CreateShoppingTaskUsecase struct {
	databaseTransactionUsecase IDatabaseTransactionUsecase
	shoppingListPort           port.IShoppingListPort
	getGroupUseCase            IGetGroupUseCase
	getGroupMemberUseCase      IGetGroupMemberUseCase
	shoppingTaskPort           port.IShoppingTaskPort
}

func (c CreateShoppingTaskUsecase) CreateNewShoppingTask(ctx context.Context, userId, shoppingListId int64, shoppingTasks []*entity.ShoppingTaskEntity) ([]*entity.ShoppingTaskEntity, error) {
	shoppingList, err := c.shoppingListPort.GetShoppingListByID(ctx, shoppingListId)
	if err != nil {
		log.Error(ctx, "Get shopping list by id error: ", err)
		return nil, err
	}
	if shoppingList.GroupID != nil {
		group, err := c.getGroupUseCase.GetGroupById(ctx, *shoppingList.GroupID)
		if err != nil {
			log.Error(ctx, "Get group by id error: ", err)
			return nil, err
		}
		isMember, err := c.getGroupMemberUseCase.IsMemberOfGroup(ctx, userId, group.ID)
		if err != nil {
			log.Error(ctx, "Check user is member of group error: ", err)
			return nil, err
		}
		if !isMember {
			log.Error(ctx, "User %d is not member of group %d", userId, group.ID)
			return nil, fmt.Errorf(constant.ErrForbiddenCreateShoppingTask)
		}
	} else {
		if shoppingList.CreatedBy != userId {
			log.Error(ctx, "User %d is not owner of shopping list %d", userId, shoppingList.ID)
			return nil, fmt.Errorf(constant.ErrForbiddenCreateShoppingTask)
		}
	}
	tx := c.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction error: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction success")
		}
	}()
	for idx, _ := range shoppingTasks {
		shoppingTasks[idx].ShoppingListID = shoppingListId
		shoppingTasks[idx].Status = constant.TODO
	}
	shoppingTasks, err = c.shoppingTaskPort.CreateNewShoppingTasks(ctx, tx, shoppingTasks)
	if err != nil {
		log.Error(ctx, "Create new shopping task error: ", err)
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit transaction error: ", err)
		return nil, err
	}
	return shoppingTasks, nil

}

func NewCreateShoppingTaskUsecase(databaseTransactionUsecase IDatabaseTransactionUsecase, shoppingListPort port.IShoppingListPort, getGroupUseCase IGetGroupUseCase, getGroupMemberUseCase IGetGroupMemberUseCase, shoppingTaskPort port.IShoppingTaskPort) ICreateShoppingTaskUsecase {
	return &CreateShoppingTaskUsecase{databaseTransactionUsecase, shoppingListPort, getGroupUseCase, getGroupMemberUseCase, shoppingTaskPort}
}
