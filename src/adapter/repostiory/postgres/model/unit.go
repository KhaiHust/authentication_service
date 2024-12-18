package model

type UnitModel struct {
	BaseModel
	Name string `gorm:"column:name"`
}

func (UnitModel) TableName() string {
	return "units"
}
