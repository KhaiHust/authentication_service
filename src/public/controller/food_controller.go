package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"strconv"
)

type FoodController struct {
	BaseController
	foodService service.IFoodService
}

func (f *FoodController) CreateFood(c *gin.Context) {
	var req request.CreateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := f.Validator.Struct(req); err != nil {
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err != nil {
			log.Error(c, "Parse error code failed: ", err)
			apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
			return
		}
		apihelper.AbortErrorHandle(c, int(errCode))
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralForbidden)
		return
	}
	foodEntity := request.FromReqToFoodEntity(&req)
	foodEntity.CreatedBy = userID
	foodEntity, err = f.foodService.CreateNewFood(c, foodEntity)
	if err != nil {
		log.Error(c, "Create food failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToFoodResponse(foodEntity))
}
func (f *FoodController) UpdatedFood(c *gin.Context) {
	foodID, err := strconv.ParseInt(c.Param("foodId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse food id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	var req request.UpdateFoodRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralForbidden)
		return
	}
	reqDto := request.FromUpdateFoodRequestDto(&req)
	foodEntity, err := f.foodService.UpdateFood(c, userID, foodID, reqDto)
	if err != nil {
		log.Error(c, "Update food failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToFoodResponse(foodEntity))
}
func (f *FoodController) DeleteFood(c *gin.Context) {
	foodID, err := strconv.ParseInt(c.Param("foodId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse food id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralForbidden)
		return
	}
	err = f.foodService.DeleteFood(c, userID, foodID)
	if err != nil {
		log.Error(c, "Delete food failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func (f *FoodController) GetAllFood(c *gin.Context) {
	pageSize, page := middleware.GetPagingParams(c)
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralForbidden)
		return
	}
	params := &dto.FoodParams{}
	params.Page = &page
	params.PageSize = &pageSize
	if name := c.Query("name"); name != "" {
		params.Name = &name
	}
	if foodType := c.Query("type"); foodType != "" {
		params.Type = &foodType
	}
	foods, err := f.foodService.GetAllFood(c, userID, params)
	if err != nil {
		log.Error(c, "Get all food failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, foods)
}
func NewFoodController(base *BaseController, foodService service.IFoodService) *FoodController {
	return &FoodController{BaseController: *base, foodService: foodService}
}
