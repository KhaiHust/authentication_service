package request

type OtpRegisterRequest struct {
	Email string `json:"email" validate:"email" message:"Email is invalid" errorCode:"400026"`
}
