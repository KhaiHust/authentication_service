package entity

type FoodEntity struct {
	BaseEntity
	Name       string
	Type       string
	CategoryID int64
	UnitID     int64
	Category   *CategoryEntity
	Unit       *UnitEntity
	ImgUrl     string
	CreatedBy  int64
}
