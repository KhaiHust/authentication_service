package request

import (
	"github.com/KhaiHust/authen_service/core/entity/dto/request"
)

type UpdateFoodRequest struct {
	Name       *string `json:"name"`
	Type       *string `json:"type"`
	ImgUrl     *string `json:"img_url"`
	CategoryID *int64  `json:"category_id"`
	UnitID     *int64  `json:"unit_id"`
}

func FromUpdateFoodRequestDto(req *UpdateFoodRequest) *request.UpdateFoodDto {
	if req == nil {
		return nil
	}
	return &request.UpdateFoodDto{
		Name:       req.Name,
		Type:       req.Type,
		ImgUrl:     req.ImgUrl,
		CategoryID: req.CategoryID,
		UnitID:     req.UnitID,
	}
}
