package properties

import "github.com/golibs-starter/golib/config"

type NotificationServiceProperties struct {
	BaseUrl                    string
	Token                      string
	TemplateOTPForRegistration string
}

func (n NotificationServiceProperties) Prefix() string {
	return "app.services.notification-service"
}

func NewNotificationServiceProperties(loader config.Loader) (*NotificationServiceProperties, error) {
	props := &NotificationServiceProperties{}
	err := loader.Bind(props)
	return props, err
}
