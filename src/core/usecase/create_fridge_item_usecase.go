package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateFridgeItemUsecase interface {
	CreateNewFridgeItem(ctx context.Context, item *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error)
}
type CreateFridgeItemUsecase struct {
	fridgeItemPort             port.IFridgeItemPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
	getFoodUseCase             IGetFoodUseCase
}

func (c CreateFridgeItemUsecase) CreateNewFridgeItem(ctx context.Context, item *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error) {
	food, err := c.getFoodUseCase.GetFoodByUserIDAndID(ctx, item.CreatedBy, item.FoodID)
	if err != nil {
		log.Error(ctx, "Get food failed: ", err)
		return nil, err
	}
	item.FoodID = food.ID
	tx := c.databaseTransactionUsecase.StartTransaction()
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
	item, err = c.fridgeItemPort.SaveFridgeItem(ctx, tx, item)
	if err != nil {
		log.Error(ctx, "Save fridge item failed: ", err)
		return nil, err
	}
	if err = tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit failed: ", err)
		return nil, err
	}
	return item, nil
}

func NewCreateFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort, databaseTransactionUsecase IDatabaseTransactionUsecase, getFoodUseCase IGetFoodUseCase) ICreateFridgeItemUsecase {
	return &CreateFridgeItemUsecase{fridgeItemPort, databaseTransactionUsecase, getFoodUseCase}
}
