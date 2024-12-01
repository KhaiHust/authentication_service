package request

type RemoveMemberRequest struct {
	GroupID int64 `json:"group_id" validate:"required"`
	UserID  int64 `json:"user_id" validate:"required"`
}
