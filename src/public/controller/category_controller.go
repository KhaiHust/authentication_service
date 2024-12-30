package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/resource/request"
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
func (c CategoryController) CreateNewCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Error(ctx, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(ctx, common.GeneralBadRequest)
		return
	}
	if err := c.Validator.Struct(&req); err != nil {
		log.Error(ctx, "Validate request failed: ", err)
		apihelper.AbortErrorHandle(ctx, common.GeneralBadRequest)
		return
	}
	categoryEntity := &entity.CategoryEntity{Name: req.Name}
	category, err := c.categoryService.CreateNewCategory(ctx, categoryEntity)
	if err != nil {
		log.Error(ctx, "Create new category failed: ", err)
		apihelper.AbortErrorHandle(ctx, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(ctx, category)
}
func NewCategoryController(base *BaseController, categoryService service.ICategoryService) *CategoryController {
	return &CategoryController{
		BaseController:  *base,
		categoryService: categoryService,
	}
}
