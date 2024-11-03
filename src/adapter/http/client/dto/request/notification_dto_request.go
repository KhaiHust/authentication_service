package request

import (
	"github.com/KhaiHust/authen_service/core/entity/dto"
)

type SendNotificationRequest struct {
	Message SendNotificationMessage `json:"message"`
}
type SendNotificationMessage struct {
	To       interface{} `json:"to"`
	Template string      `json:"template"`
	Data     interface{} `json:"data"`
}
type SendOTPSignUpTo struct {
	Email string `json:"email"`
}
type SendOTPForRegistrationRequest struct {
	RecipientName string `json:"recipientName"`
	OTPSignUp     string `json:"otpSignUp"`
}

func ToSendOTPForRegistrationRequest(request *dto.OTPSignupNotificationDto, template string) *SendNotificationRequest {
	if request == nil {
		return nil
	}
	var sendNotiMessage SendNotificationMessage

	sendNotiMessage.To = SendOTPSignUpTo{Email: request.Email}
	sendNotiMessage.Template = template
	sendNotiMessage.Data = SendOTPForRegistrationRequest{
		RecipientName: request.RecipientName,
		OTPSignUp:     request.OTP,
	}

	return &SendNotificationRequest{Message: sendNotiMessage}
}
