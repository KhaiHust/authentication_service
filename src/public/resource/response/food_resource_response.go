package response

import "github.com/KhaiHust/authen_service/core/entity"

type FoodResponse struct {
	ID         int64             `json:"id"`
	Name       string            `json:"name"`
	Type       string            `json:"type"`
	ImgUrl     string            `json:"img_url,omitempty"`
	CategoryID int64             `json:"category_id,omitempty"`
	UnitID     int64             `json:"unit_id"`
	Category   *CategoryResponse `json:"category,omitempty"`
	Unit       *UnitResponse     `json:"unit,omitempty"`
}

func FromEntityToFoodResponse(entity *entity.FoodEntity) *FoodResponse {
	if entity == nil {
		return nil
	}
	return &FoodResponse{
		ID:         entity.ID,
		Name:       entity.Name,
		Type:       entity.Type,
		ImgUrl:     entity.ImgUrl,
		CategoryID: entity.CategoryID,
		UnitID:     entity.UnitID,
		Category:   ToCategoryResponse(entity.Category),
		Unit:       ToUnitResponse(entity.Unit),
	}
}
