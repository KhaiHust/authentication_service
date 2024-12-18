package request

type UpdateFoodDto struct {
	Name       *string
	Type       *string
	ImgUrl     *string
	CategoryID *int64
	UnitID     *int64
}
