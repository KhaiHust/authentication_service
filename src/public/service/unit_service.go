package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/helper"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/response"
)

type IUnitService interface {
	GetAllUnits(ctx context.Context, params *dto.UnitParamDto) (*response.ListUnitResponse, error)
}
type UnitService struct {
	getUnitUsecase usecase.IGetUnitUsecase
}

func (u UnitService) GetAllUnits(ctx context.Context, params *dto.UnitParamDto) (*response.ListUnitResponse, error) {
	units, err := u.getUnitUsecase.GetAllUnit(ctx, params)
	if err != nil {
		return nil, err
	}
	total, err := u.getUnitUsecase.CountAllUnit(ctx, params)
	if err != nil {
		return nil, err
	}
	page := int64(*params.Page)
	pageSize := int64(*params.PageSize)
	nextPage, prePage, totalPage := helper.CalculateParameterForGetRequest(page, pageSize, total)

	return response.ToListUnitResponse(units, page, pageSize, totalPage, total, prePage, nextPage), nil
}

func NewUnitService(getUnitUsecase usecase.IGetUnitUsecase) IUnitService {
	return &UnitService{
		getUnitUsecase: getUnitUsecase,
	}
}
