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

type FridgeController struct {
	*BaseController
	fridgeItemService service.IFridgeItemService
}

func (f *FridgeController) CreateFridgeItem(c *gin.Context) {
	var req request.CreateNewFridgeItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := f.Validator.Struct(req); err != nil {
		log.Error(c, "Validate request failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	fridgeItem := request.FromRequestToFridgeItemEntity(&req)
	if fridgeItem == nil {
		log.Error(c, "Convert request to entity failed")
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	fridgeItem.CreatedBy = userID
	item, err := f.fridgeItemService.SaveFridgeItem(c, fridgeItem)
	if err != nil {
		log.Error(c, "Save fridge item failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToFridgeItemResponse(item))
}
func NewFridgeController(base *BaseController, fridgeItemService service.IFridgeItemService) *FridgeController {
	return &FridgeController{
		BaseController:    base,
		fridgeItemService: fridgeItemService}
}
