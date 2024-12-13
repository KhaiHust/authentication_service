package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"time"
)

func ToShoppingListEntity(model *model.ShoppingListModel) *entity.ShoppingListEntity {
	return &entity.ShoppingListEntity{
		BaseEntity: entity.BaseEntity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		Name:        model.Name,
		Description: model.Description,
		CreatedBy:   model.CreatedBy,
		AssignedTo:  model.AssignedTo,
		DueDate:     model.DueDate.Unix(),
		GroupID:     model.GroupID,
	}
}
func ToShoppingListModel(entityShoppingListEntity *entity.ShoppingListEntity) *model.ShoppingListModel {
	return &model.ShoppingListModel{
		BaseModel: model.BaseModel{
			ID: entityShoppingListEntity.ID,
		},
		Name:        entityShoppingListEntity.Name,
		Description: entityShoppingListEntity.Description,
		CreatedBy:   entityShoppingListEntity.CreatedBy,
		AssignedTo:  entityShoppingListEntity.AssignedTo,
		DueDate:     time.Unix(entityShoppingListEntity.DueDate, 0),
		GroupID:     entityShoppingListEntity.GroupID,
	}
}
