package request

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required" message:"Name is required"`
}
