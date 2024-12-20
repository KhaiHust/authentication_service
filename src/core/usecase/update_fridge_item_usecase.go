package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
)

type IUpdateFridgeItemUsecase interface {
	UpdateItem(ctx context.Context, itemID int64, req *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error)
}
type UpdateFridgeItemUsecase struct {
	fridgeItemPort             port.IFridgeItemPort
	getFridgeItemUsecase       IGetFridgeItemUsecase
	getFoodUseCase             IGetFoodUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (u UpdateFridgeItemUsecase) UpdateItem(ctx context.Context, itemID int64, req *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error) {
	item, err := u.getFridgeItemUsecase.GetItemByIDAndCreatedBy(ctx, itemID, 1)
	if err != nil {
		return nil, err
	}

	food, err := u.getFoodUseCase.GetFoodByUserIDAndID(ctx, item.CreatedBy, *req.FoodID)
	if err != nil {
		return nil, err
	}
	item.FoodID = food.ID
	item.Quantity = *req.Quantity
	tx := u.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			return
		}
	}()
	item, err = u.fridgeItemPort.UpdateItem(ctx, tx, item)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		return nil, err
	}
	return item, nil
}

func NewUpdateFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort, getFridgeItemUsecase IGetFridgeItemUsecase, getFoodUseCase IGetFoodUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) IUpdateFridgeItemUsecase {
	return &UpdateFridgeItemUsecase{fridgeItemPort, getFridgeItemUsecase, getFoodUseCase, databaseTransactionUsecase}
}
