package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"time"
)

func ToFridgeItemModel(entity *entity.FridgeItemEntity) *model.FridgeItemModel {
	if entity == nil {
		return nil
	}
	return &model.FridgeItemModel{
		BaseModel: model.BaseModel{
			ID: entity.ID,
		},
		ExpiredDate: time.Unix(entity.ExpiredDate, 0),
		Quantity:    entity.Quantity,
		FoodID:      entity.FoodID,
		CreatedBy:   entity.CreatedBy,
		Note:        entity.Note,
	}
}
func ToFridgeItemEntity(model *model.FridgeItemModel) *entity.FridgeItemEntity {
	if model == nil {
		return nil
	}
	return &entity.FridgeItemEntity{
		BaseEntity: entity.BaseEntity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		ExpiredDate: model.ExpiredDate.Unix(),
		Quantity:    model.Quantity,
		FoodID:      model.FoodID,
		CreatedBy:   model.CreatedBy,
		Note:        model.Note,
	}
}
func ToListFridgeItemEntity(models []*model.FridgeItemModel) []*entity.FridgeItemEntity {
	var entities []*entity.FridgeItemEntity
	for _, model := range models {
		entities = append(entities, ToFridgeItemEntity(model))
	}
	return entities
}
