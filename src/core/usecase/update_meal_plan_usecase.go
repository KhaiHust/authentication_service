package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IUpdateMealPlanUsecase interface {
	UpdateMealPlan(ctx context.Context, userID, mealPlanID int64, req *request.UpdateMealPlanDTO) (*entity.MealPlanEntity, error)
}
type UpdateMealPlanUsecase struct {
	databaseTransactionUsecase IDatabaseTransactionUsecase
	mealPlanPort               port.IMealPlanPort
	mealPlanFoodPort           port.IMealPlanFoodPort
	getMealPlanUsecase         IGetMealPlanUsecase
}

func (u UpdateMealPlanUsecase) UpdateMealPlan(ctx context.Context, userID, mealPlanID int64, req *request.UpdateMealPlanDTO) (*entity.MealPlanEntity, error) {
	mealPlan, err := u.getMealPlanUsecase.GetMealPlanByUserIDAndID(ctx, userID, mealPlanID)
	if err != nil {
		log.Error(ctx, "Get meal plan failed: ", err)
		return nil, err
	}
	if req.Name != nil {
		mealPlan.Name = *req.Name
	}
	if req.Description != nil {
		mealPlan.Description = *req.Description
	}
	if req.Schedule != nil {
		mealPlan.Schedule = *req.Schedule
	}
	tx := u.databaseTransactionUsecase.StartTransaction()
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
	if len(req.FoodIDs) > 0 {
		err = u.mealPlanFoodPort.DeleteListMealPlanFood(ctx, tx, mealPlanID)
		if err != nil {
			log.Error(ctx, "Delete list meal plan food failed: ", err)
			return nil, err
		}
		foodIDs := make([]int64, 0)
		for _, foodID := range req.FoodIDs {
			foodIDs = append(foodIDs, foodID)
		}
		mpFEntities := make([]*entity.MealPlanFoodEntity, 0)
		for _, foodID := range mealPlan.FoodIDs {
			mpFEntities = append(mpFEntities, &entity.MealPlanFoodEntity{
				MealPlanID: mealPlan.ID,
				FoodID:     foodID,
			})
		}
		if len(mpFEntities) > 0 {
			mpFEntities, err = u.mealPlanFoodPort.SaveListMealPlanFood(ctx, tx, mpFEntities)
			if err != nil {
				log.Error(ctx, "Save list meal plan food failed: ", err)
				return nil, err
			}
		}
	}
	mealPlan, err = u.mealPlanPort.UpdateMealPlan(ctx, tx, mealPlan.ID, mealPlan)
	if err != nil {
		log.Error(ctx, "Update meal plan failed: ", err)
		return nil, err
	}
	if errCommit := tx.Commit().Error; errCommit != nil {
		log.Error(ctx, "Commit transaction failed: ", errCommit)
		return nil, errCommit
	}
	return mealPlan, nil
}

func NewUpdateMealPlanUsecase(databaseTransactionUsecase IDatabaseTransactionUsecase, mealPlanPort port.IMealPlanPort, mealPlanFoodPort port.IMealPlanFoodPort, getMealPlanUsecase IGetMealPlanUsecase) IUpdateMealPlanUsecase {
	return &UpdateMealPlanUsecase{
		databaseTransactionUsecase: databaseTransactionUsecase,
		mealPlanPort:               mealPlanPort,
		mealPlanFoodPort:           mealPlanFoodPort,
		getMealPlanUsecase:         getMealPlanUsecase,
	}
}
