package controller

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/KhaiHust/authen_service/public/resource/request"
	"github.com/KhaiHust/authen_service/public/service"
	"github.com/gin-gonic/gin"
	"github.com/golibs-starter/golib/log"
	"strconv"
)

type OtpController struct {
	otpService service.IOtpService
	BaseController
}

func NewOtpController(otpService service.IOtpService, base *BaseController) *OtpController {
	return &OtpController{otpService: otpService, BaseController: *base}
}
func (o *OtpController) SendOtpForRegistration(c *gin.Context) {
	var req request.OtpRegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error(c, "Error while binding request", err)
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	if err := o.Validator.Struct(req); err != nil {
		log.Error(c, "Error while validating request", err)
		errCode, err := strconv.ParseInt(err.Error(), 10, 64)
		if err == nil {
			apihelper.AbortErrorHandle(c, int(errCode))
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralBadRequest)
		return
	}
	err := o.otpService.SendOtpForRegistration(c, req.Email)
	if err != nil {
		log.Error(c, "Error while sending otp", err)
		if err.Error() == constant.ErrUserNotFound {
			apihelper.AbortErrorHandle(c, common.UserNotExistErrCode)
			return
		}
		apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
		return
	}
	apihelper.SuccessfulHandle(c, nil)
}
