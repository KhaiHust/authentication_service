package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetFridgeItemUsecase interface {
	GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error)
}
type GetFridgeItemUsecase struct {
	fridgeItemPort port.IFridgeItemPort
}

func (g GetFridgeItemUsecase) GetItemByIDAndCreatedBy(ctx context.Context, id, createdBy int64) (*entity.FridgeItemEntity, error) {
	item, err := g.fridgeItemPort.GetItemByIDAndCreatedBy(ctx, id, createdBy)
	if err != nil {
		log.Error(ctx, "Get fridge item failed: ", err)
		return nil, err
	}
	return item, nil
}

func NewGetFridgeItemUsecase(fridgeItemPort port.IFridgeItemPort) IGetFridgeItemUsecase {
	return &GetFridgeItemUsecase{fridgeItemPort: fridgeItemPort}
}
