package usecase

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
)

type ISendOtpUseCase interface {
	SendOtpForRegistration(ctx *context.Context, email string) error
}

type SendOtpUseCase struct {
	notificationPort port.INotificationPort
	getUserUseCase   IGetUserUsecase
	cachePort        port.ICachePort
}

func (s *SendOtpUseCase) SendOtpForRegistration(ctx *context.Context, email string) error {
	existedUser, err := s.getUserUseCase.GetUserByEmail(ctx, email)
	if err != nil {
		log.Error(ctx, "SendOtpForRegistration: GetUserByEmail error", err)
		return err
	}
	if existedUser != nil && !existedUser.IsVerified {
		otp, err := common.GenerateRandomOTPRegister()
		if err != nil {
			log.Error(ctx, "SendOtpForRegistration: GenerateRandomOTPRegister error", err)
			return err
		}
		reqSendOtp := &dto.OTPSignupNotificationDto{
			RecipientName: existedUser.Name,
			Email:         email,
			OTP:           otp,
		}
		err = s.notificationPort.SendOTPForRegistration(ctx, reqSendOtp)
		if err != nil {
			log.Error(ctx, "SendOtpForRegistration: SendOTPForRegistration error", err)
			return err
		}
		// Save OTP to cache
		key := common.BuildCacheKeyOTPRegister(email)
		err = s.cachePort.SetToCache(*ctx, key, otp, constant.DefaultCacheTTL)
		//todo: rate limit for resend otp
		if err != nil {
			log.Error(ctx, "SendOtpForRegistration: SetToCache error", err)
			return err
		}
		return nil
	}
	return errors.New(constant.ErrUserVerified)
}

func NewSendOtpUseCase(notificationPort port.INotificationPort, getUserUseCase IGetUserUsecase, cachePort port.ICachePort) ISendOtpUseCase {
	return &SendOtpUseCase{
		notificationPort: notificationPort,
		getUserUseCase:   getUserUseCase,
		cachePort:        cachePort,
	}
}
