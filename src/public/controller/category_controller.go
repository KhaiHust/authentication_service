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

type CategoryController struct {
	BaseController
	categoryService service.ICategoryService
}

func (c CategoryController) GetAllCategory(ctx *gin.Context) {
	pageSize, page := middleware.GetPagingParams(ctx)
	spec := &dto.CategorySpec{
		BaseSpec: dto.BaseSpec{
			PageSize: &pageSize,
			Page:     &page,
		},
	}
	name := ctx.Query("name")
	if len(name) > 0 {
		spec.Name = &name
	}
	result, err := c.categoryService.GetAllCategory(ctx, spec)
	if err != nil {
		log.Error(ctx, "Get all category failed", err)
		apihelper.AbortErrorHandle(ctx, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(ctx, result)

}
func NewCategoryController(base *BaseController, categoryService service.ICategoryService) *CategoryController {
	return &CategoryController{
		BaseController:  *base,
		categoryService: categoryService,
	}
}
