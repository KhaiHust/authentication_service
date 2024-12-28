package response

type LoginResponseDto struct {
	IsVerified   bool   `json:"is_verified"`
	AccessToken  string `json:"access_token,omitempty"`
	RefreshToken string `json:"refresh_token,omitempty"`
}
