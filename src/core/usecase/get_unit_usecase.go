package usecase

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type IGetUnitUsecase interface {
	GetAllUnit(ctx context.Context, params *dto.UnitParamDto) ([]*entity.UnitEntity, error)
	CountAllUnit(ctx context.Context, params *dto.UnitParamDto) (int64, error)
}
type GetUnitUsecase struct {
	unitPort port.IUnitPort
}

func (g GetUnitUsecase) GetAllUnit(ctx context.Context, params *dto.UnitParamDto) ([]*entity.UnitEntity, error) {
	result, err := g.unitPort.GetAllUnit(ctx, params)
	if err != nil {
		log.Error(ctx, "Get all unit failed", err)
		return nil, err
	}
	return result, nil
}

func (g GetUnitUsecase) CountAllUnit(ctx context.Context, params *dto.UnitParamDto) (int64, error) {
	result, err := g.unitPort.CountAllUnit(ctx, params)
	if err != nil {
		log.Error(ctx, "Count all unit failed", err)
		return 0, err
	}
	return result, nil
}

func NewGetUnitUsecase(unitPort port.IUnitPort) IGetUnitUsecase {
	return &GetUnitUsecase{unitPort: unitPort}
}
