package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ICreateMealPlanUsecase interface {
	CreateNewMealPlan(ctx context.Context, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error)
}
type CreateMealPlanUsecase struct {
	mealPlanPort               port.IMealPlanPort
	mealPlanFoodPort           port.IMealPlanFoodPort
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (c CreateMealPlanUsecase) CreateNewMealPlan(ctx context.Context, mpEntity *entity.MealPlanEntity) (*entity.MealPlanEntity, error) {
	tx := c.databaseTransactionUsecase.StartTransaction()
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := tx.Rollback(); errRollback != nil {
			log.Error(ctx, "Rollback transaction failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback transaction successfully")
		}
	}()
	foodIds := make([]int64, 0)
	for _, foodID := range mpEntity.FoodIDs {
		foodIds = append(foodIds, foodID)
	}
	mpEntity, err = c.mealPlanPort.SaveNewMealPlan(ctx, tx, mpEntity)
	if err != nil {
		log.Error(ctx, "Save new meal plan failed: ", err)
		return nil, err
	}
	mpFEntities := make([]*entity.MealPlanFoodEntity, 0)
	for _, foodID := range foodIds {
		mpFEntities = append(mpFEntities, &entity.MealPlanFoodEntity{
			MealPlanID: mpEntity.ID,
			FoodID:     foodID,
		})
	}
	if len(mpFEntities) > 0 {
		mpFEntities, err = c.mealPlanFoodPort.SaveListMealPlanFood(ctx, tx, mpFEntities)
		if err != nil {
			log.Error(ctx, "Save list meal plan food failed: ", err)
			return nil, err
		}
	}
	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Error(ctx, "Commit transaction failed: ", errCommit)
		return nil, errCommit
	}
	mpEntity.FoodIDs = foodIds
	return mpEntity, nil
}

func NewCreateMealPlanUsecase(mealPlanPort port.IMealPlanPort, mealPlanFoodPort port.IMealPlanFoodPort, databaseTransactionUsecase IDatabaseTransactionUsecase) ICreateMealPlanUsecase {
	return &CreateMealPlanUsecase{
		mealPlanPort:               mealPlanPort,
		mealPlanFoodPort:           mealPlanFoodPort,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
