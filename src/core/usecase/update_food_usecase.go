package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/exception"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IUpdateFoodUseCase interface {
	UpdateFood(ctx context.Context, userID, foodID int64, updateFoodDto *request.UpdateFoodDto) (*entity.FoodEntity, error)
}
type UpdateFoodUseCase struct {
	foodPort                   port.IFoodPort
	getFoodUseCase             IGetFoodUseCase
	databaseTransactionUsecase IDatabaseTransactionUsecase
}

func (u UpdateFoodUseCase) UpdateFood(ctx context.Context, userID, foodID int64, updateFoodDto *request.UpdateFoodDto) (*entity.FoodEntity, error) {
	food, err := u.getFoodUseCase.GetFoodByUserIDAndID(ctx, userID, foodID)
	if err != nil {
		log.Error(ctx, "Get food failed: ", err)
		return nil, err
	}
	if updateFoodDto.Name != nil {
		food.Name = *updateFoodDto.Name
	}
	if updateFoodDto.Type != nil {
		food.Type = *updateFoodDto.Type
	}
	if updateFoodDto.ImgUrl != nil {
		food.ImgUrl = *updateFoodDto.ImgUrl
	}
	if updateFoodDto.CategoryID != nil {
		food.CategoryID = *updateFoodDto.CategoryID
	}
	if updateFoodDto.UnitID != nil {
		food.UnitID = *updateFoodDto.UnitID
	}
	tx := u.databaseTransactionUsecase.StartTransaction()
	defer func() {
		if r := recover(); r != nil {
			err = exception.InternalServerErrorException
		}
		if errRollback := u.databaseTransactionUsecase.Rollback(tx); errRollback != nil {
			log.Error(ctx, "Rollback failed: ", errRollback)
		} else {
			log.Info(ctx, "Rollback successfully")
		}
	}()
	food, err = u.foodPort.UpdateFood(ctx, tx, food)
	if err != nil {
		log.Error(ctx, "Update food failed: ", err)
		return nil, err
	}
	if errCommit := u.databaseTransactionUsecase.Commit(tx); errCommit != nil {
		log.Error(ctx, "Commit failed: ", errCommit)
		return nil, errCommit
	}
	return food, nil
}

func NewUpdateFoodUseCase(foodPort port.IFoodPort, getFoodUseCase IGetFoodUseCase, databaseTransactionUsecase IDatabaseTransactionUsecase) IUpdateFoodUseCase {
	return &UpdateFoodUseCase{
		foodPort:                   foodPort,
		getFoodUseCase:             getFoodUseCase,
		databaseTransactionUsecase: databaseTransactionUsecase,
	}
}
