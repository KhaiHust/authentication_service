package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity/dto"
)

type INotificationPort interface {
	SendOTPForRegistration(ctx *context.Context, request *dto.OTPSignupNotificationDto) error
}
