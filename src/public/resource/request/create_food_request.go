package request

import "github.com/KhaiHust/authen_service/core/entity"

type CreateFoodRequest struct {
	Name       string `json:"name" validate:"required" message:"Name is required"`
	Type       string `json:"type" validate:"required" message:"Type is required"`
	CategoryID int64  `json:"category_id"`
	UnitID     int64  `json:"unit_id" validate:"required" message:"Unit is required"`
	ImgUrl     string `json:"img_url"`
}

func FromReqToFoodEntity(req *CreateFoodRequest) *entity.FoodEntity {
	if req == nil {
		return nil
	}
	return &entity.FoodEntity{
		Name:       req.Name,
		Type:       req.Type,
		CategoryID: req.CategoryID,
		UnitID:     req.UnitID,
	}
}
