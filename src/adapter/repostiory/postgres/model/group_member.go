package model

type GroupMemberModel struct {
	BaseModel
	GroupID int64 `gorm:"column:group_id;type:int;not null"`
	UserID  int64 `gorm:"column:user_id;type:int;not null"`
	RoleID  int64 `gorm:"column:role_id;type:int;not null"`
}

func (GroupMemberModel) TableName() string {
	return "group_members"
}
