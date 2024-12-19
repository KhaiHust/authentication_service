package dto

type FoodParams struct {
	UserID *int64
	Name   *string
	Type   *string
	BaseSpec
}
