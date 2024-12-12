package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToShoppingListGroupEntity(model *model.ShoppingListGroupModel) *entity.ShoppingListGroupEntity {
	return &entity.ShoppingListGroupEntity{
		BaseEntity: entity.BaseEntity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		ShoppingListID: model.ShoppingListID,
		GroupID:        model.GroupID,
	}
}
func ToShoppingListGroupModel(entityShoppingListGroupEntity *entity.ShoppingListGroupEntity) *model.ShoppingListGroupModel {
	return &model.ShoppingListGroupModel{
		BaseModel: model.BaseModel{
			ID: entityShoppingListGroupEntity.ID,
		},
		ShoppingListID: entityShoppingListGroupEntity.ShoppingListID,
		GroupID:        entityShoppingListGroupEntity.GroupID,
	}
}
