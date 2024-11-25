package common

import "net/http"

const (
	// GeneralServiceUnavailable General
	GeneralServiceUnavailable = 500

	// GeneralBadRequest General
	GeneralBadRequest = 400

	//GeneralUnauthorized error
	GeneralUnauthorized = 401

	//ErrBadRequest error
	InvalidEmailErrCode    = 400026
	InvalidPasswordErrCode = 400027
	InvalidNameErrCode     = 400028
	ExistedEmailErrCode    = 400032
	UserNotExistErrCode    = 400036

	OtpNotProvidedErrCode = 400053
	OtpInvalidErrCode     = 400054
	OTPExpiredErrCode     = 400055

	WrongPasswordErrCode = 400056
)

const (
	InvalidEmailMessage    = "Vui lòng cung cấp email hợp lệ!"
	InvalidPasswordMessage = "Vui lòng cung cấp mật khẩu dài hơn 6 ký tự và ít hơn 20 ký tự!"
	InvalidNameMessage     = "Vui lòng cung cấp tên dài hơn 3 ký tự và ít hơn 30 ký tự!"
	ExistedEmailMessage    = "Một tài khoản đã tồn tại với email này!"
	UserNotExistMessage    = "Không tìm thấy tài khoản với địa chỉ email này."
	OtpNotProvidedMessage  = "Vui lòng cung cấp mã OTP!"
	OtpInvalidMessage      = "Mã OTP không hợp lệ!"
	OTPExpiredMessage      = "Mã OTP đã hết hạn!"
	WrongPasswordMessage   = "Email / Mật khẩu không đúng!"
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
	GeneralUnauthorized: {
		HTTPCode:    http.StatusUnauthorized,
		ServiceCode: GeneralUnauthorized,
		Message:     "Unauthorized",
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
	UserNotExistErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: UserNotExistErrCode,
		Message:     UserNotExistMessage,
	},
	OtpNotProvidedErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: OtpNotProvidedErrCode,
		Message:     OtpNotProvidedMessage,
	},
	OtpInvalidErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: OtpInvalidErrCode,
		Message:     OtpInvalidMessage,
	},
	OTPExpiredErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: OTPExpiredErrCode,
		Message:     OTPExpiredMessage,
	},
	WrongPasswordErrCode: {
		HTTPCode:    http.StatusBadRequest,
		ServiceCode: WrongPasswordErrCode,
		Message:     WrongPasswordMessage,
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
