package response

import "github.com/KhaiHust/authen_service/core/entity"

type ListUnitResponse struct {
	Units        []*UnitResponse `json:"units"`
	TotalItems   int64           `json:"total_items"`
	TotalPages   int64           `json:"total_pages"`
	CurrentPage  int64           `json:"current_page"`
	PageSize     int64           `json:"page_size"`
	PreviousPage *int64          `json:"previous_page"`
	NextPage     *int64          `json:"next_page"`
}
type UnitResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToUnitResponse(unit *entity.UnitEntity) *UnitResponse {
	if unit == nil {
		return nil
	}
	return &UnitResponse{
		ID:   unit.ID,
		Name: unit.Name,
	}
}
func ToUnitsResponse(units []*entity.UnitEntity) []*UnitResponse {
	var res []*UnitResponse
	for _, unit := range units {
		res = append(res, &UnitResponse{
			ID:   unit.ID,
			Name: unit.Name,
		})
	}
	return res
}
func ToListUnitResponse(units []*entity.UnitEntity, page, pageSize, totalPage, total int64, prePage, nextPage *int64) *ListUnitResponse {
	return &ListUnitResponse{
		Units:        ToUnitsResponse(units),
		CurrentPage:  page,
		PageSize:     pageSize,
		PreviousPage: prePage,
		TotalItems:   total,
		TotalPages:   totalPage,
		NextPage:     nextPage,
	}
}
