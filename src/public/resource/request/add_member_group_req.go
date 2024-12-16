package request

type AddMemberGroupReq struct {
	GroupID int64  `json:"group_id" validate:"required"`
	Email   string `json:"email" validate:"email" validate:"email" message:"Email is invalid" errorCode:"400026"`
}
