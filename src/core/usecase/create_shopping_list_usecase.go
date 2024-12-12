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

type ICreateShoppingListUseCase interface {
	CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error)
}
type CreateShoppingListUseCase struct {
	shoppingListPort           port.IShoppingListPort
	shoppingListGroupPort      port.IShoppingListGroupPort
	getGroupUseCase            IGetGroupUseCase
	getGroupMemberUseCase      IGetGroupMemberUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (c CreateShoppingListUseCase) CreateNewShoppingList(ctx context.Context, req *request.CreateShoppingListDto) (*entity.ShoppingListEntity, error) {
	validAssignTo := int64(0)
	if req.GroupID > 0 {
		//create shopping list for group
		group, err := c.getGroupUseCase.GetGroupById(ctx, req.GroupID)
		if err != nil {
			log.Error(ctx, "CreateNewShoppingList: GetGroupById error", err)
			return nil, err
		}
		groupMembers, err := c.getGroupMemberUseCase.GetListMemberByGroupID(ctx, req.CreatedBy, group.ID)
		if err != nil {
			log.Error(ctx, "CreateNewShoppingList: GetListMemberByGroupID error", err)
			return nil, err
		}
		//check permission
		hasPermission := false
		for _, member := range groupMembers {
			if member.UserID == req.CreatedBy {
				hasPermission = true
				break
			}
		}
		if !hasPermission {
			log.Error(ctx, "CreateNewShoppingList: User %d has no permission to create shopping list for group %d", req.CreatedBy, req.GroupID)
			return nil, errors.New(constant.ErrForbiddenCreateShoppingList)
		}
		if req.AssignedTo > 0 {
			hasPermission = false
			for _, member := range groupMembers {
				if member.UserID == req.AssignedTo {
					hasPermission = true
					break
				}
			}
			if !hasPermission {
				log.Error(ctx, "CreateNewShoppingList: User %d has no permission to assign shopping list for group %d", req.AssignedTo, req.GroupID)
				return nil, errors.New(constant.ErrForbiddenCreateShoppingList)
			}
			validAssignTo = req.AssignedTo
		}
	}
	if validAssignTo == 0 {
		validAssignTo = req.CreatedBy
	}
	//create shopping list
	shoppingList := &entity.ShoppingListEntity{
		Name:        req.Name,
		Description: req.Description,
		CreatedBy:   req.CreatedBy,
		AssignedTo:  validAssignTo,
		DueDate:     req.DueDate,
	}
	var err error
	tx := c.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := c.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "CreateNewShoppingList: Rollback error", errRollback)
		} else {
			log.Info(ctx, "CreateNewShoppingList: Rollback successfully")
		}
	}()
	shoppingList, err = c.shoppingListPort.CreateNewShoppingList(ctx, tx, shoppingList)
	if err != nil {
		log.Error(ctx, "CreateNewShoppingList: CreateNewShoppingList error", err)
		return nil, err
	}
	if req.GroupID > 0 {
		_, err = c.shoppingListGroupPort.CreateNewShoppingListGroup(ctx, tx, &entity.ShoppingListGroupEntity{
			ShoppingListID: shoppingList.ID,
			GroupID:        req.GroupID,
		})
		if err != nil {
			log.Error(ctx, "CreateNewShoppingList: CreateNewShoppingListGroup error", err)
			return nil, err
		}
	}
	errCommit := c.databaseTransactionUsecase.Commit(tx)
	if errCommit != nil {
		log.Error(ctx, "CreateNewShoppingList: Commit error", errCommit)
		return nil, errCommit
	}
	return shoppingList, nil
}

func NewCreateShoppingListUseCase(shoppingListPort port.IShoppingListPort, shoppingListGroupPort port.IShoppingListGroupPort, getGroupUseCase IGetGroupUseCase, getGroupMemberUseCase IGetGroupMemberUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateShoppingListUseCase {
	return &CreateShoppingListUseCase{
		shoppingListPort:           shoppingListPort,
		shoppingListGroupPort:      shoppingListGroupPort,
		getGroupUseCase:            getGroupUseCase,
		getGroupMemberUseCase:      getGroupMemberUseCase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
