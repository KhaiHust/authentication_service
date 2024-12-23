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
	"strconv"
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
func (f *FridgeController) UpdateFridgeItem(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
	if err != nil {
		log.Error(c, "Get item id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	var req request.UpdateFridgeItemRequest
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

	fridgeItem, err := f.fridgeItemService.UpdateFridgeItem(c, userID, itemID, request.ToUpdateFridgeItemDto(&req))
	if err != nil {
		log.Error(c, "Update fridge item failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToFridgeItemResponse(fridgeItem))
}
func (f *FridgeController) DeleteFridgeItem(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
	if err != nil {
		log.Error(c, "Get item id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	err = f.fridgeItemService.DeleteItem(c, userID, itemID)
	if err != nil {
		log.Error(c, "Delete fridge item failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func (f *FridgeController) GetFridgeItem(c *gin.Context) {
	itemID, err := strconv.ParseInt(c.Param("itemId"), 10, 64)
	if err != nil {
		log.Error(c, "Get item id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	fridgeItem, err := f.fridgeItemService.GetFridgeItemDetail(c, userID, itemID)
	if err != nil {
		log.Error(c, "Get fridge item failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToFridgeItemResponse(fridgeItem))
}
func (f *FridgeController) GetAllFridgeItems(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "Get user id failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	items, err := f.fridgeItemService.GetAllItems(c, userID)
	if err != nil {
		log.Error(c, "Get all fridge items failed: ", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromListEntityToFridgeItemResponse(items))
}
func NewFridgeController(base *BaseController, fridgeItemService service.IFridgeItemService) *FridgeController {
	return &FridgeController{
		BaseController:    base,
		fridgeItemService: fridgeItemService}
}
