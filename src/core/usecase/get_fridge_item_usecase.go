package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetFridgeItemUsecase interface {
	GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error)
	GetItemDetailByID(ctx context.Context, userID, itemID int64) (*entity.FridgeItemEntity, error)
	GetAllItems(ctx context.Context, userID int64) ([]*entity.FridgeItemEntity, error)
}
type GetFridgeItemUsecase struct {
	fridgeItemPort port.IFridgeItemPort
	getFoodUseCase IGetFoodUseCase
}

func (g GetFridgeItemUsecase) GetAllItems(ctx context.Context, userID int64) ([]*entity.FridgeItemEntity, error) {
	items, err := g.fridgeItemPort.GetAllItems(ctx, userID)
	if err != nil {
		log.Error(ctx, "Get all fridge item failed: ", err)
		return nil, err
	}
	foodIDs := make([]int64, 0)
	for _, item := range items {
		foodIDs = append(foodIDs, item.FoodID)
	}
	foods, err := g.getFoodUseCase.GetFoodByIDs(ctx, foodIDs)
	if err != nil {
		log.Error(ctx, "Get food by IDs failed: ", err)
		return nil, err
	}
	foodMap := make(map[int64]*entity.FoodEntity)
	for _, food := range foods {
		foodMap[food.ID] = food
	}
	for idx, _ := range items {
		items[idx].Food = foodMap[items[idx].FoodID]
	}
	return items, nil
}

func (g GetFridgeItemUsecase) GetItemDetailByID(ctx context.Context, userID, itemID int64) (*entity.FridgeItemEntity, error) {
	item, err := g.fridgeItemPort.GetItemByIDAndCreatedBy(ctx, itemID, userID)
	if err != nil {
		log.Error(ctx, "Get fridge item failed: ", err)
		return nil, err
	}
	food, err := g.getFoodUseCase.GetFoodByUserIDAndID(ctx, item.CreatedBy, item.FoodID)
	if err != nil {
		log.Error(ctx, "Get food failed: ", err)
		return nil, err
	}
	item.Food = food
	return item, nil
}

func (g GetFridgeItemUsecase) GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error) {
	item, err := g.fridgeItemPort.GetItemByIDAndCreatedBy(ctx, id, createdBy)
	if err != nil {
		log.Error(ctx, "Get fridge item failed: ", err)
		return nil, err
	}
	return item, nil
}

func NewGetFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort, getFoodUseCase IGetFoodUseCase) IGetFridgeItemUsecase {
	return &GetFridgeItemUsecase{
		fridgeItemPort: fridgeItemPort,
		getFoodUseCase: getFoodUseCase,
	}
}
