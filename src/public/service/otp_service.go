package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IOtpService interface {
	SendOtpForRegistration(ctx context.Context, email string) error
	VerifiedOtpForRegistration(ctx context.Context, email string, otp string) error
}
type OtpService struct {
	sendOtpUsecase   usecase.ISendOtpUseCase
	verifyOtpUsecase usecase.IVerifyOtpUseCase
}

func (o *OtpService) VerifiedOtpForRegistration(ctx context.Context, email string, otp string) error {
	return o.verifyOtpUsecase.VerifyOtpRegister(&ctx, email, otp)
}

func (o *OtpService) SendOtpForRegistration(ctx context.Context, email string) error {
	return o.sendOtpUsecase.SendOtpForRegistration(&ctx, email)
}

func NewOtpService(sendOtpUsecase usecase.ISendOtpUseCase, verifyOtpUsecase usecase.IVerifyOtpUseCase) IOtpService {
	return &OtpService{
		sendOtpUsecase:   sendOtpUsecase,
		verifyOtpUsecase: verifyOtpUsecase,
	}
}
