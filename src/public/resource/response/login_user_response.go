package response

// LoginUserResponse struct
type LoginUserResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func ToLoginUserResponse(accessToken string, refreshToken string) *LoginUserResponse {
	return &LoginUserResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
}
