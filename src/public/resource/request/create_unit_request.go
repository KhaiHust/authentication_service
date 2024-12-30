package request

type CreateUnitRequest struct {
	Name string `json:"name" validate:"required" message:"name is required"`
}
