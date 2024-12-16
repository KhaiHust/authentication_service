package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToCategoryModel(entity *entity.CategoryEntity) *model.CategoryModel {
	return &model.CategoryModel{
		BaseModel: model.BaseModel{
			ID: entity.ID,
		},
		Name: entity.Name,
	}
}
func ToCategoryEntity(model *model.CategoryModel) *entity.CategoryEntity {
	return &entity.CategoryEntity{
		BaseEntity: entity.BaseEntity{
			ID:        model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		Name: model.Name,
	}
}
func ToListCategoryEntity(models []*model.CategoryModel) []*entity.CategoryEntity {
	if models == nil {
		return nil
	}
	var entities []*entity.CategoryEntity
	for _, model := range models {
		entities = append(entities, ToCategoryEntity(model))
	}
	return entities
}
