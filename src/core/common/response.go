package common

import "net/http"

const (
	// GeneralServiceUnavailable General
	GeneralServiceUnavailable = 500

	// GeneralBadRequest General
	GeneralBadRequest = 400

	//ErrBadRequest error
	InvalidEmailErrCode    = 400026
	InvalidPasswordErrCode = 400027
	InvalidNameErrCode     = 400028
	ExistedEmailErrCode    = 400032
)

const (
	InvalidEmailMessage    = "Vui lòng cung cấp email hợp lệ!"
	InvalidPasswordMessage = "Vui lòng cung cấp mật khẩu dài hơn 6 ký tự và ít hơn 20 ký tự!"
	InvalidNameMessage     = "Vui lòng cung cấp tên dài hơn 3 ký tự và ít hơn 30 ký tự!"
	ExistedEmailMessage    = "Một tài khoản đã tồn tại với email này!"
)

const ()

// ErrorResponse error response struct
type ErrorResponse struct {
	HTTPCode    int
	ServiceCode int
	Message     string
}

var errorResponseMap = map[int]ErrorResponse{
	GeneralServiceUnavailable: {
		HTTPCode:    http.StatusInternalServerError,
		ServiceCode: GeneralServiceUnavailable,
		Message:     "Service unavailable",
	},
	GeneralBadRequest: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: GeneralBadRequest,
		Message:     "Bad request",
	},
	InvalidEmailErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: InvalidEmailErrCode,
		Message:     InvalidEmailMessage,
	},
	InvalidPasswordErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: InvalidPasswordErrCode,
		Message:     InvalidPasswordMessage,
	},
	InvalidNameErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: InvalidNameErrCode,
		Message:     InvalidNameMessage,
	},
	ExistedEmailErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: ExistedEmailErrCode,
		Message:     ExistedEmailMessage,
	},
}

// GetErrorResponse get error response from code
func GetErrorResponse(code int) ErrorResponse {
	if val, ok := errorResponseMap[code]; ok {
		return val
	}

	// default response
	return ErrorResponse{
		HTTPCode:    http.StatusInternalServerError,
		ServiceCode: code,
		Message:     http.StatusText(http.StatusInternalServerError),
	}
}
