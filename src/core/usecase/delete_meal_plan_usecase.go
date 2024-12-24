package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IDeleteMealPlanUsecase interface {
	DeleteMealPlan(ctx context.Context, userID, mealID int64) error
}
type DeleteMealPlanUsecase struct {
	mealPlanPort               port.IMealPlanPort
	mealPlanFoodPort           port.IMealPlanFoodPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (d DeleteMealPlanUsecase) DeleteMealPlan(ctx context.Context, userID, mealID int64) error {
	mealPlan, err := d.mealPlanPort.GetMealPlanByUserIDAndID(ctx, userID, mealID)
	if err != nil {
		log.Error(ctx, "Get meal plan failed: ", err)
		return err
	}
	tx := d.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction success")
		}
	}()
	if err := d.mealPlanFoodPort.DeleteListMealPlanFood(ctx, tx, mealPlan.ID); err != nil {
		log.Error(ctx, "Delete meal plan food failed: ", err)
		return err
	}
	if err := d.mealPlanPort.DeleteMealPlanByID(ctx, tx, mealPlan.ID); err != nil {
		log.Error(ctx, "Delete meal plan failed: ", err)
		return err
	}
	if err := tx.Commit().Error; err != nil {
		log.Error(ctx, "Commit transaction failed: ", err)
		return err
	}
	return nil
}

func NewDeleteMealPlanUsecase(mealPlanPort port.IMealPlanPort, mealPlanFoodPort port.IMealPlanFoodPort, databaseTransactionUsecase IDatabaseTransactionUsecase) IDeleteMealPlanUsecase {
	return &DeleteMealPlanUsecase{
		mealPlanPort:               mealPlanPort,
		mealPlanFoodPort:           mealPlanFoodPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
