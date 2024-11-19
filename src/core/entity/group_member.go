package entity

type GroupMemberEntity struct {
	BaseEntity
	GroupID int64
	UserID  int64
	RoleID  int64
}
