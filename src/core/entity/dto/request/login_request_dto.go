package request

type LoginRequestDto struct {
	Email     string
	Password  string
	IpAddress string
	UserAgent string
}
