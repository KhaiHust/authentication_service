package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToShoppingTaskModel(entity *entity.ShoppingTaskEntity) *model.ShoppingTaskModel {
	return &model.ShoppingTaskModel{
		BaseModel:      model.BaseModel{ID: entity.ID},
		ShoppingListID: entity.ShoppingListID,
		FoodName:       entity.FoodName,
		Quantity:       entity.Quantity,
		Status:         entity.Status,
	}
}
func ToShoppingTaskEntity(model *model.ShoppingTaskModel) *entity.ShoppingTaskEntity {
	return &entity.ShoppingTaskEntity{
		BaseEntity: entity.BaseEntity{ID: model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix()},
		ShoppingListID: model.ShoppingListID,
		FoodName:       model.FoodName,
		Quantity:       model.Quantity,
		Status:         model.Status,
	}
}
func ToListShoppingTaskModel(entities []*entity.ShoppingTaskEntity) []*model.ShoppingTaskModel {
	var models []*model.ShoppingTaskModel
	for _, entity := range entities {
		models = append(models, ToShoppingTaskModel(entity))
	}
	return models
}
func ToListShoppingTaskEntity(models []*model.ShoppingTaskModel) []*entity.ShoppingTaskEntity {
	var entities []*entity.ShoppingTaskEntity
	for _, model := range models {
		entities = append(entities, ToShoppingTaskEntity(model))
	}
	return entities
}
