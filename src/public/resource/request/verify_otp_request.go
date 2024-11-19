package request

type OtpVerifyRequest struct {
	Email string `json:"email" validate:"email" message:"Email is invalid" errorCode:"400026"`
	OTP   string `json:"otp" validate:"required" message:"OTP is required" errorCode:"400053"`
}
