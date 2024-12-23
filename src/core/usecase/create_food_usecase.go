package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateFoodUsecase interface {
	CreateFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error)
}
type CreateFoodUsecase struct {
	foodPort                   port.IFoodPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (c CreateFoodUsecase) CreateFood(ctx context.Context, foodEntity *entity.FoodEntity) (*entity.FoodEntity, error) {
	tx := c.databaseTransactionUsecase.StartTransaction()
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error("Rollback transaction failed: ", errRollback)
		} else {
			log.Info("Rollback transaction success")
		}
	}()
	foodEntity, err = c.foodPort.SaveFood(ctx, tx, foodEntity)
	if err != nil {
		log.Error(ctx, "Save food failed: ", err)
		return nil, err
	}
	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Error(ctx, "Commit transaction failed: ", errCommit)
		return nil, errCommit
	}
	return foodEntity, nil
}

func NewCreateFoodUsecase(foodPort port.IFoodPort, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateFoodUsecase {
	return &CreateFoodUsecase{
		foodPort:                   foodPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
