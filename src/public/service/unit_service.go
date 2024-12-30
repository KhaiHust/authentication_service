package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/helper"
	"github.com/KhaiHust/authen_service/core/usecase"
	"github.com/KhaiHust/authen_service/public/resource/response"
)

type IUnitService interface {
	GetAllUnits(ctx context.Context, params *dto.UnitParamDto) (*response.ListUnitResponse, error)
	CreateUnit(ctx context.Context, unit *entity.UnitEntity) (*response.UnitResponse, error)
}
type UnitService struct {
	getUnitUsecase    usecase.IGetUnitUsecase
	createUnitUseCase usecase.ICreateUnitUseCase
}

func (u UnitService) CreateUnit(ctx context.Context, unit *entity.UnitEntity) (*response.UnitResponse, error) {
	unit, err := u.createUnitUseCase.CreateNewUnit(ctx, unit)
	if err != nil {
		return nil, err
	}
	return response.ToUnitResponse(unit), nil
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

func NewUnitService(getUnitUsecase usecase.IGetUnitUsecase, createUnitUseCase usecase.ICreateUnitUseCase) IUnitService {
	return &UnitService{
		getUnitUsecase:    getUnitUsecase,
		createUnitUseCase: createUnitUseCase,
	}
}
