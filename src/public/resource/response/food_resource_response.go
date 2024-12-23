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
func ToListFoodResponse(entities []*entity.FoodEntity) []*FoodResponse {
	responses := make([]*FoodResponse, 0)
	for _, foodEntity := range entities {
		responses = append(responses, FromEntityToFoodResponse(foodEntity))
	}
	return responses
}

type GetFoodResponse struct {
	Foods        []*FoodResponse `json:"foods"`
	TotalItems   int64           `json:"total_items"`
	TotalPages   int64           `json:"total_pages"`
	CurrentPage  int64           `json:"current_page"`
	PageSize     int64           `json:"page_size"`
	PreviousPage *int64          `json:"previous_page"`
	NextPage     *int64          `json:"next_page"`
}

func ToGetFoodResponse(foods []*entity.FoodEntity, page, pageSize, totalPage, total int64, prePage, nextPage *int64) *GetFoodResponse {
	return &GetFoodResponse{
		Foods:        ToListFoodResponse(foods),
		CurrentPage:  page,
		PageSize:     pageSize,
		PreviousPage: prePage,
		TotalItems:   total,
		TotalPages:   totalPage,
		NextPage:     nextPage,
	}
}
