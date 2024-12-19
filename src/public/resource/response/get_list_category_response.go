package response

import "github.com/KhaiHust/authen_service/core/entity"

type ListCategoryResponse struct {
	Categories   []*CategoryResponse `json:"categories"`
	TotalItems   int64               `json:"total_items"`
	TotalPages   int64               `json:"total_pages"`
	CurrentPage  int64               `json:"current_page"`
	PageSize     int64               `json:"page_size"`
	PreviousPage *int64              `json:"previous_page"`
	NextPage     *int64              `json:"next_page"`
}

type CategoryResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func ToCategoryResponse(cate *entity.CategoryEntity) *CategoryResponse {
	if cate == nil {
		return nil
	}
	return &CategoryResponse{
		ID:   cate.ID,
		Name: cate.Name,
	}
}
func ToCategoriesResponse(cates []*entity.CategoryEntity) []*CategoryResponse {
	var res []*CategoryResponse
	for _, cate := range cates {
		res = append(res, &CategoryResponse{
			ID:   cate.ID,
			Name: cate.Name,
		})
	}
	return res
}
func ToListCategoryResponse(cates []*entity.CategoryEntity, page, pageSize, totalPage, total int64, prePage, nextPage *int64) *ListCategoryResponse {
	return &ListCategoryResponse{
		Categories:   ToCategoriesResponse(cates),
		CurrentPage:  page,
		PageSize:     pageSize,
		PreviousPage: prePage,
		TotalItems:   total,
		TotalPages:   totalPage,
		NextPage:     nextPage,
	}
}
