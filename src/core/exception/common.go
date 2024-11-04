package exception

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/golibs-starter/golib/exception"
)

var (
	InternalServerErrorException = exception.New(common.GeneralServiceUnavailable, common.GetErrorResponse(common.GeneralServiceUnavailable).Message)
)
