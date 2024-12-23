package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
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

func NewMealPlanController(BaseController *BaseController, mealPlanService service.IMealPlanService) *MealPlanController {
	return &MealPlanController{
		BaseController:  BaseController,
		mealPlanService: mealPlanService}
}
