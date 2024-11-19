package request

// LoginRequest struct
type LoginRequest struct {
	Email    string `json:"email" validate:"email" message:"Email is invalid" errorCode:"400026"`
	Password string `json:"password" validate:"gt=6,lt=20" message:"Password is required and must be between 7 and 19 characters" errorCode:"400027"`
}
