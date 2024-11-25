package model

type GroupModel struct {
	BaseModel
	Name        string `gorm:"column:name;type:varchar(255);not null"`
	Description string `gorm:"column:description;type:text;not null"`
}

func (GroupModel) TableName() string {
	return "groups"
}
