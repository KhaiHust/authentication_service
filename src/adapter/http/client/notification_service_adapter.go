package client

import (
	"context"
	"errors"
	request2 "github.com/KhaiHust/authen_service/adapter/http/client/dto/request"
	"github.com/KhaiHust/authen_service/adapter/http/client/dto/response"
	"github.com/KhaiHust/authen_service/adapter/properties"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
	"github.com/golibs-starter/golib/web/client"
)

const (
	SEND_NOTIFICATION_ENDPOINT = "notification/send"
)

type NotificationServiceAdapter struct {
	props      *properties.NotificationServiceProperties
	httpClient client.ContextualHttpClient
}

func (n *NotificationServiceAdapter) SendOTPForRegistration(ctx context.Context, request *dto.OTPSignupNotificationDto) error {
	sendOtpRequest := request2.ToSendOTPForRegistrationRequest(request, n.props.TemplateOTPForRegistration)
	if sendOtpRequest == nil {
		log.Error(ctx, "SendOTPForRegistrationRequest is nil")
		return errors.New("SendOTPForRegistrationRequest is nil")
	}
	var notiResponse *response.NotificationDtoResponse
	resp, err := n.httpClient.Post(ctx, n.props.BaseUrl+SEND_NOTIFICATION_ENDPOINT, sendOtpRequest, notiResponse)
	if err != nil {
		log.Error(ctx, "SendOTPForRegistrationRequest error: %v", err)
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		log.Error(ctx, "SendOTPForRegistrationRequest fail, http status code: %d", resp.StatusCode)
		return errors.New("SendOTPForRegistrationRequest fail, http status code: " + string(rune(resp.StatusCode)))
	}
	if notiResponse == nil || notiResponse.RequestId == "" {
		log.Error(ctx, "SendOTPForRegistrationRequest fail, response is nil")
		return errors.New("SendOTPForRegistrationRequest fail, response is nil")
	}
	return nil
}

func NewNotificationServiceAdapter(props *properties.NotificationServiceProperties) port.INotificationPort {
	return &NotificationServiceAdapter{props: props}
}
