package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToFoodModel(foodEntity *entity.FoodEntity) *model.FoodModel {
	if foodEntity == nil {
		return nil
	}

	foodModel := &model.FoodModel{
		BaseModel: model.BaseModel{ID: foodEntity.ID},
		Name:      foodEntity.Name,
		Type:      foodEntity.Type,
		UnitID:    foodEntity.UnitID,
		ImgUrl:    foodEntity.ImgUrl,
		CreatedBy: foodEntity.CreatedBy,
	}
	if foodEntity.CategoryID != 0 {
		foodModel.CategoryID = &foodEntity.CategoryID
	}
	return foodModel
}
func ToFoodEntity(foodModel *model.FoodModel) *entity.FoodEntity {
	if foodModel == nil {
		return nil
	}
	foodEntity := &entity.FoodEntity{
		BaseEntity: entity.BaseEntity{ID: foodModel.ID,
			CreatedAt: foodModel.CreatedAt.Unix(),
			UpdatedAt: foodModel.UpdatedAt.Unix()},
		Name:      foodModel.Name,
		Type:      foodModel.Type,
		UnitID:    foodModel.UnitID,
		ImgUrl:    foodModel.ImgUrl,
		Category:  ToCategoryEntity(foodModel.Category),
		Unit:      ToUnitEntity(foodModel.Unit),
		CreatedBy: foodModel.CreatedBy,
	}
	if foodModel.CategoryID != nil {
		foodEntity.CategoryID = *foodModel.CategoryID
	}
	return foodEntity
}
