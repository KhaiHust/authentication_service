package response

// LoginUserResponse struct
type LoginUserResponse struct {
	IsVerified   bool   `json:"is_verified"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}

func ToLoginUserResponse(isVerified bool, accessToken, refreshToken string) *LoginUserResponse {
	return &LoginUserResponse{
		IsVerified:   isVerified,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
