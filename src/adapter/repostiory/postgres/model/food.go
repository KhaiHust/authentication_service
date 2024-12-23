package model

type FoodModel struct {
	BaseModel
	Name       string         `gorm:"column:name"`
	Type       string         `gorm:"column:type"`
	CategoryID *int64         `gorm:"column:category_id"`
	UnitID     int64          `gorm:"column:unit_id"`
	ImgUrl     string         `gorm:"column:img_url"`
	Category   *CategoryModel `gorm:"foreignKey:CategoryID"`
	Unit       *UnitModel     `gorm:"foreignKey:UnitID"`
	CreatedBy  int64          `gorm:"column:created_by"`
}

func (FoodModel) TableName() string {
	return "foods"
}
