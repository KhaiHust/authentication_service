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

type ShoppingTaskController struct {
	BaseController
	shoppingTaskService service.IShoppingTaskService
}

func (s *ShoppingTaskController) CreateNewShoppingTask(c *gin.Context) {
	shoppingListId, err := strconv.ParseInt(c.Param("shoppingListId"), 10, 64)
	if err != nil {
		log.Error(c, "CreateNewShoppingTask: ParseInt error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	var req request.CreateShoppingTaskRequest
	if err := c.ShouldBind(&req); err != nil {
		log.Error(c, "CreateNewShoppingTask: ShouldBindJSON error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := s.Validator.Struct(req); err != nil {
		log.Error(c, "CreateNewShoppingTask: Validator error", err)
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err != nil {
			log.Error(c, "CreateNewShoppingTask: ParseInt error", err)
			apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
			return
		}
		apihelper.AbortErrorHandle(c, int(errCode))
		return
	}
	userId, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "CreateNewShoppingTask: GetUserID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	result, err := s.shoppingTaskService.CreateNewShoppingTask(c, userId, shoppingListId, request.ToCreateTasksEntity(req))
	if err != nil {
		log.Error(c, "CreateNewShoppingTask: CreateNewShoppingTask error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToCreateTaskResponse(shoppingListId, result))
}
func (s *ShoppingTaskController) GetShoppingTasksByShoppingListID(c *gin.Context) {
	shoppingListId, err := strconv.ParseInt(c.Param("shoppingListId"), 10, 64)
	if err != nil {
		log.Error(c, "GetShoppingTasksByShoppingListID: ParseInt error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userId, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "GetShoppingTasksByShoppingListID: GetUserID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	result, err := s.shoppingTaskService.GetShoppingTasksByShoppingListID(c, userId, shoppingListId)
	if err != nil {
		log.Error(c, "GetShoppingTasksByShoppingListID: GetShoppingTasksByShoppingListID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToCreateTaskResponse(shoppingListId, result))
}
func (s *ShoppingTaskController) DeleteTaskByID(c *gin.Context) {
	shoppingListId, err := strconv.ParseInt(c.Param("shoppingListId"), 10, 64)
	if err != nil {
		log.Error(c, "DeleteTaskByID: ParseInt error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	taskId, err := strconv.ParseInt(c.Param("taskId"), 10, 64)
	if err != nil {
		log.Error(c, "DeleteTaskByID: ParseInt error", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	userId, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "DeleteTaskByID: GetUserID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
		return
	}
	err = s.shoppingTaskService.DeleteTaskByID(c, userId, shoppingListId, taskId)
	if err != nil {
		log.Error(c, "DeleteTaskByID: DeleteTaskByID error", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
func NewShoppingTaskController(base *BaseController, shoppingTaskService service.IShoppingTaskService) *ShoppingTaskController {
	return &ShoppingTaskController{BaseController: *base,
		shoppingTaskService: shoppingTaskService}
}
