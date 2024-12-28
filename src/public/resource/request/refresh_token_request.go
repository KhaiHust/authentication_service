package request

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required" message:"Refresh token is required" errorCode:"400028"`
}
