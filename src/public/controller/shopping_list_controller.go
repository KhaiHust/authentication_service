package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	request2 "github.com/KhaiHust/authen_service/core/entity/dto/request"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/middleware"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/resource/response"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"strconv"
)

type ShoppingListController struct {
	BaseController
	shoppingListService service.IShoppingListService
}

func (s *ShoppingListController) CreateNewShoppingList(c *gin.Context) {
	var req request.CreateShoppingListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind Json error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := s.Validator.Struct(req); err != nil {
		log.Error(c, "Validate error: %v", err)
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err != nil {
			log.Error(c, "Parse error: %v", err)
			apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
			return
		}
		apihelper.AbortErrorHandle(c, int(errCode))
		return
	}
	userId, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	result, err := s.shoppingListService.CreateNewShoppingList(c, &request2.CreateShoppingListDto{
		Name:        req.Name,
		Description: req.Description,
		GroupID:     req.GroupID,
		AssignedTo:  req.AssignedTo,
		CreatedBy:   userId,
		DueDate:     req.DueDate,
	})
	if err != nil {
		log.Error(c, "CreateShoppingList error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToCreateShoppingListResponse(result))
}
func (s *ShoppingListController) UpdateShoppingList(c *gin.Context) {
	shoppingListId, err := strconv.ParseInt(c.Param("shoppingListId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	var req request.UpdateShoppingListRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind Json error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := s.Validator.Struct(req); err != nil {
		log.Error(c, "Validate error: %v", err)
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err != nil {
			log.Error(c, "Parse error: %v", err)
			apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
			return
		}
		apihelper.AbortErrorHandle(c, int(errCode))
		return
	}

	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	reqDto := request.ToUpdateShoppingListDto(&req)
	reqDto.ID = shoppingListId
	result, err := s.shoppingListService.UpdateShoppingList(c, userID, reqDto)
	if err != nil {
		log.Error(c, "UpdateShoppingList error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToCreateShoppingListResponse(result))
}
func (s *ShoppingListController) DeleteShoppingList(c *gin.Context) {
	shoppingListId, err := strconv.ParseInt(c.Param("shoppingListId"), 10, 64)
	if err != nil {
		log.Error(c, "Parse error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "error when getting user id", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	err = s.shoppingListService.DeleteShoppingListByID(c, userID, shoppingListId)
	if err != nil {
		log.Error(c, "DeleteShoppingList error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func NewShoppingListController(baseController *BaseController, shoppingListService service.IShoppingListService) *ShoppingListController {
	return &ShoppingListController{
		BaseController:      *baseController,
		shoppingListService: shoppingListService,
	}
}
