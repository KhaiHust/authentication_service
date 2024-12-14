package usecase

import (
	"context"
	"fmt"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetTaskUsecase interface {
	GetTaskByShoppingListID(ctx context.Context, userID, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error)
}
type GetTaskUsecase struct {
	shoppingTaskPort      port.IShoppingTaskPort
	getGroupUseCase       IGetGroupUseCase
	getGroupMemberUseCase IGetGroupMemberUseCase
	shoppingListPort      port.IShoppingListPort
}

func (g GetTaskUsecase) GetTaskByShoppingListID(ctx context.Context, userID, shoppingListID int64) ([]*entity.ShoppingTaskEntity, error) {
	shoppingList, err := g.shoppingListPort.GetShoppingListByID(ctx, shoppingListID)
	if err != nil {
		log.Error(ctx, "Get shopping list by id error: ", err)
		return nil, err
	}
	if shoppingList.GroupID != nil {
		group, err := g.getGroupUseCase.GetGroupById(ctx, *shoppingList.GroupID)
		if err != nil {
			log.Error(ctx, "Get group by id error: ", err)
			return nil, err
		}
		isMember, err := g.getGroupMemberUseCase.IsMemberOfGroup(ctx, userID, group.ID)
		if err != nil {
			log.Error(ctx, "Check user is member of group error: ", err)
			return nil, err
		}
		if !isMember {
			log.Error(ctx, "User %d is not member of group %d", userID, group.ID)
			return nil, fmt.Errorf(constant.ErrForbiddenGetShoppingTask)
		}
	} else {
		if shoppingList.CreatedBy != userID {
			log.Error(ctx, "User %d is not owner of shopping list %d", userID, shoppingList.ID)
			return nil, fmt.Errorf(constant.ErrForbiddenGetShoppingTask)
		}
	}
	shoppingTasks, err := g.shoppingTaskPort.GetShoppingTasksByShoppingListID(ctx, shoppingListID)
	if err != nil {
		log.Error(ctx, "Get shopping tasks by shopping list id error: ", err)
		return nil, err
	}
	return shoppingTasks, nil
}

func NewGetTaskUsecase(shoppingTaskPort port.IShoppingTaskPort, getGroupUseCase IGetGroupUseCase, getGroupMemberUseCase IGetGroupMemberUseCase, shoppingListPort port.IShoppingListPort) IGetTaskUsecase {
	return &GetTaskUsecase{shoppingTaskPort, getGroupUseCase, getGroupMemberUseCase, shoppingListPort}
}
