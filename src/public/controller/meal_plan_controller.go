package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"strconv"
)

type MealPlanController struct {
	*BaseController
	mealPlanService service.IMealPlanService
}

func (m *MealPlanController) CreateNewMealPlan(c *gin.Context) {
	var req request.CreateNewMealPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := m.Validator.Struct(&req); err != nil {
		log.Error(c, "Validate request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	mpEntity := request.FromCreateRequest(&req)
	mpEntity, err = m.mealPlanService.CreateNewMealPlan(c, userID, mpEntity)
	if err != nil {
		log.Error(c, "Create new meal plan failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToMealPlanResourceResponse(mpEntity))
}
func (m *MealPlanController) DeleteMealPlan(c *gin.Context) {
	mealPlanID, err := strconv.ParseInt(c.Param("mealPlanId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse meal plan ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	if err := m.mealPlanService.DeleteMealPlan(c, userID, mealPlanID); err != nil {
		log.Error(c, "Delete meal plan failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func (m *MealPlanController) UpdateMealPlan(c *gin.Context) {
	mealPlanID, err := strconv.ParseInt(c.Param("mealPlanId"), 10, 64)
	var req request.UpdateMealPlanRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := m.Validator.Struct(&req); err != nil {
		log.Error(c, "Validate request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}

	mpEntity, err := m.mealPlanService.UpdateMealPlan(c, userID, mealPlanID, request.FromReqToUpdateMealPlanDto(&req))
	if err != nil {
		log.Error(c, "Update meal plan failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToMealPlanResourceResponse(mpEntity))
}
func (m *MealPlanController) GetPlanByDate(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	date, err := strconv.ParseInt(apihelper.GetRequestParams(c, constant.MealPlanDateParams), 10, 64)
	if err != nil {
		log.Error(c, "Parse date failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	mpEntities, err := m.mealPlanService.GetMealPlanByDate(c, userID, date)
	if err != nil {
		log.Error(c, "Get meal plan by date failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntitiesToMealPlanResourceResponses(mpEntities))
}
func (m *MealPlanController) GetDetailMealPlan(c *gin.Context) {
	mealPlanID, err := strconv.ParseInt(c.Param("mealId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse meal plan ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user ID failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	mpEntity, err := m.mealPlanService.GetDetailMealPlan(c, userID, mealPlanID)
	if err != nil {
		log.Error(c, "Get detail meal plan failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToMealPlanResourceResponse(mpEntity))
}
func NewMealPlanController(BaseController *BaseController, mealPlanService service.IMealPlanService) *MealPlanController {
	return &MealPlanController{
		BaseController:  BaseController,
		mealPlanService: mealPlanService}
}
