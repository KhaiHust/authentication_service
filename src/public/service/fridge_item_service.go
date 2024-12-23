package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IFridgeItemService interface {
	SaveFridgeItem(ctx context.Context, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error)
	UpdateFridgeItem(ctx context.Context, userID, itemID int64, reqDto *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error)
	DeleteItem(ctx context.Context, userID, itemID int64) error
	GetFridgeItemDetail(ctx context.Context, userID, itemID int64) (*entity.FridgeItemEntity, error)
	GetAllItems(ctx context.Context, userID int64) ([]*entity.FridgeItemEntity, error)
}
type FridgeItemService struct {
	createFridgeItemUsecase usecase.ICreateFridgeItemUsecase
	updateFridgeItemUsecase usecase.IUpdateFridgeItemUsecase
	deleteFridgeItemUsecase usecase.IDeleteFridgeItemUsecase
	getFridgeItemUsecase    usecase.IGetFridgeItemUsecase
}

func (f FridgeItemService) GetAllItems(ctx context.Context, userID int64) ([]*entity.FridgeItemEntity, error) {
	return f.getFridgeItemUsecase.GetAllItems(ctx, userID)
}

func (f FridgeItemService) GetFridgeItemDetail(ctx context.Context, userID, itemID int64) (*entity.FridgeItemEntity, error) {
	return f.getFridgeItemUsecase.GetItemDetailByID(ctx, userID, itemID)
}

func (f FridgeItemService) DeleteItem(ctx context.Context, userID, itemID int64) error {
	return f.deleteFridgeItemUsecase.DeleteItem(ctx, userID, itemID)
}

func (f FridgeItemService) UpdateFridgeItem(ctx context.Context, userID, itemID int64, reqDto *request.UpdateFridgeItemDto) (*entity.FridgeItemEntity, error) {
	return f.updateFridgeItemUsecase.UpdateItem(ctx, userID, itemID, reqDto)
}

func (f FridgeItemService) SaveFridgeItem(ctx context.Context, fridgeItem *entity.FridgeItemEntity) (*entity.FridgeItemEntity, error) {
	return f.createFridgeItemUsecase.CreateNewFridgeItem(ctx, fridgeItem)
}

func NewFridgeItemService(createFridgeItemUsecase usecase.ICreateFridgeItemUsecase, updateFridgeItemUsecase usecase.IUpdateFridgeItemUsecase, deleteFridgeItemUsecase usecase.IDeleteFridgeItemUsecase, getFridgeItemUsecase usecase.IGetFridgeItemUsecase) IFridgeItemService {
	return &FridgeItemService{
		createFridgeItemUsecase: createFridgeItemUsecase,
		updateFridgeItemUsecase: updateFridgeItemUsecase,
		deleteFridgeItemUsecase: deleteFridgeItemUsecase,
		getFridgeItemUsecase:    getFridgeItemUsecase,
	}
}
