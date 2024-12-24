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

type UserController struct {
	BaseController
	userService service.IUserService
}

func NewUserController(baseController *BaseController, userService service.IUserService) *UserController {
	return &UserController{
		BaseController: *baseController,
		userService:    userService,
	}
}

func (u *UserController) RegisterUser(c *gin.Context) {
	var req request.RegisterUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind Json error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}

	if err := u.Validator.Struct(req); err != nil {
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
	result, err := u.userService.CreateUser(c, &req)
	if err != nil {
		log.Error(c, "CreateUser error: %v", err)
		if err.Error() == constant.ErrExistedEmail {
			apihelper.AbortErrorHandle(c, common.ExistedEmailErrCode)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.FromEntityToRegisterUserResponse(result))

}
func (u *UserController) LoginUser(c *gin.Context) {
	var req request.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind Json error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := u.Validator.Struct(req); err != nil {
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
	result, err := u.userService.LoginUser(c, req.Email, req.Password)
	if err != nil {
		log.Error(c, "LoginUser error: %v", err)
		if err.Error() == constant.ErrUserNotFound {
			apihelper.AbortErrorHandle(c, common.UserNotExistErrCode)
			return
		}
		if err.Error() == constant.ErrWrongPassword {
			apihelper.AbortErrorHandle(c, common.WrongPasswordErrCode)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToLoginUserResponse(result.IsVerified, result.AccessToken, result.RefreshToken))
}
func (u *UserController) RefreshToken(c *gin.Context) {
	var req request.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Bind Json error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := u.Validator.Struct(req); err != nil {
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
	result, err := u.userService.GetRefreshTokenByToken(c, req.RefreshToken)
	if err != nil {
		log.Error(c, "GetRefreshTokenByToken error: %v", err)
		if err.Error() == constant.ErrInvalidRefreshToken {
			apihelper.AbortErrorHandle(c, common.InvalidRefreshTokenCode)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, response.ToLoginUserResponse(result.IsVerified, result.AccessToken, result.RefreshToken))
}
func (u *UserController) Logout(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		log.Error(c, "GetUserID error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	err = u.userService.Logout(c, userID)
	if err != nil {
		log.Error(c, "Logout error: %v", err)
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
