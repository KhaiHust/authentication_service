package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IUpdateFridgeItemUsecase interface {
	UpdateItem(ctx context.Context, userID, itemID int64, req *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error)
}
type UpdateFridgeItemUsecase struct {
	fridgeItemPort             port.IFridgeItemPort
	getFridgeItemUsecase       IGetFridgeItemUsecase
	getFoodUseCase             IGetFoodUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (u UpdateFridgeItemUsecase) UpdateItem(ctx context.Context, userID, itemID int64, req *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error) {
	item, err := u.getFridgeItemUsecase.GetItemByIDAndCreatedBy(ctx, itemID, userID)
	if err != nil {
		log.Error(ctx, "Get fridge item failed: ", err)
		return nil, err
	}
	if req.FoodID != nil {
		food, err := u.getFoodUseCase.GetFoodByUserIDAndID(ctx, item.CreatedBy, *req.FoodID)
		if err != nil {
			log.Error(ctx, "Get food failed: ", err)
			return nil, err
		}
		item.FoodID = food.ID
	}
	if req.Quantity != nil {
		item.Quantity = *req.Quantity
	}
	tx := u.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback successfully")
		}
	}()
	item, err = u.fridgeItemPort.UpdateItem(ctx, tx, item)
	if err != nil {
		log.Error(ctx, "Update fridge item failed: ", err)
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit failed: ", err)
		return nil, err
	}
	return item, nil
}

func NewUpdateFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort, getFridgeItemUsecase IGetFridgeItemUsecase, getFoodUseCase IGetFoodUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) IUpdateFridgeItemUsecase {
	return &UpdateFridgeItemUsecase{fridgeItemPort, getFridgeItemUsecase, getFoodUseCase, databaseTransactionUsecase}
}
