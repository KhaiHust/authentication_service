package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IFridgeItemService interface {
	SaveFridgeItem(ctx context.Context, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error)
}
type FridgeItemService struct {
	createFridgeItemUsecase usecase.ICreateFridgeItemUsecase
}

func (f FridgeItemService) SaveFridgeItem(ctx context.Context, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error) {
	return f.createFridgeItemUsecase.CreateNewFridgeItem(ctx, fridgeItem)
}

func NewFridgeItemService(createFridgeItemUsecase usecase.ICreateFridgeItemUsecase) IFridgeItemService {
	return &FridgeItemService{createFridgeItemUsecase}
}
