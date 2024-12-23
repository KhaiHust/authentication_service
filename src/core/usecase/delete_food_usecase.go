package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IDeleteFoodUseCase interface {
	DeleteFood(ctx context.Context, userID, foodID int64) error
}
type DeleteFoodUseCase struct {
	foodPort                   port.IFoodPort
	getFoodUseCase             IGetFoodUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (d DeleteFoodUseCase) DeleteFood(ctx context.Context, userID, foodID int64) error {
	food, err := d.getFoodUseCase.GetFoodByUserIDAndID(ctx, userID, foodID)
	if err != nil {
		return err
	}
	tx := d.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := d.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback successfully")
		}
	}()
	err = d.foodPort.DeleteFood(ctx, tx, food.ID)
	if err != nil {
		log.Error(ctx, "Delete food failed: ", err)
		return err
	}
	if errCommit := d.databaseTransactionUsecase.Commit(tx); errCommit != nil {
		log.Error(ctx, "Commit failed: ", errCommit)
		return errCommit
	}
	return nil
}

func NewDeleteFoodUseCase(foodPort port.IFoodPort, getFoodUseCase IGetFoodUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) IDeleteFoodUseCase {
	return &DeleteFoodUseCase{
		foodPort:                   foodPort,
		getFoodUseCase:             getFoodUseCase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
