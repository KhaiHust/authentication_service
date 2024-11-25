package model

type GroupRoleModel struct {
	BaseModel
	Name string `gorm:"column:name;type:varchar(255);not null"`
	Code string `gorm:"column:code;type:varchar(255);not null"`
}

func (GroupRoleModel) TableName() string {
	return "group_roles"
}
