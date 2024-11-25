package controller

import (
	"github.com/KhaiHust/authen_service/public/apihelper"
)

type BaseController struct {
	Validator *apihelper.CustomValidate
}

func NewBaseController(validator *apihelper.CustomValidate) *BaseController {
	return &BaseController{
		Validator: validator,
	}
}
