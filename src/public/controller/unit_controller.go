package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
)

type UnitController struct {
	unitService service.IUnitService
	BaseController
}

func (u *UnitController) GetAllUnits(c *gin.Context) {
	pageSize, page := middleware.GetPagingParams(c)
	params := &dto.UnitParamDto{
		BaseSpec: dto.BaseSpec{
			PageSize: &pageSize,
			Page:     &page,
		},
	}
	name := c.Query("name")
	if len(name) > 0 {
		params.Name = &name
	}
	result, err := u.unitService.GetAllUnits(c, params)
	if err != nil {
		log.Error(c, "Get all unit failed", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, result)

}

func NewUnitController(base *BaseController, unitService service.IUnitService) *UnitController {
	return &UnitController{
		BaseController: *base,
		unitService:    unitService,
	}
}
