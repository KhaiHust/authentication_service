package service

import (
	"context"
	"github.com/KhaiHust/authen_service/core/usecase"
)

type IOtpService interface {
	SendOtpForRegistration(ctx context.Context, email string) error
}
type OtpService struct {
	sendOtpUsecase usecase.ISendOtpUseCase
}

func (o *OtpService) SendOtpForRegistration(ctx context.Context, email string) error {
	return o.sendOtpUsecase.SendOtpForRegistration(&ctx, email)
}

func NewOtpService(sendOtpUsecase usecase.ISendOtpUseCase) IOtpService {
	return &OtpService{sendOtpUsecase: sendOtpUsecase}
}
