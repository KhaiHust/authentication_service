package model

type CategoryModel struct {
	BaseModel
	Name string `gorm:"column:name"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
