package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToUnitModel(entity *entity.UnitEntity) *model.UnitModel {
	return &model.UnitModel{
		BaseModel: model.BaseModel{
			ID: entity.ID,
		},
		Name: entity.Name,
	}
}
func ToUnitEntity(model *model.UnitModel) *entity.UnitEntity {
	if model == nil {
		return nil
	}
	return &entity.UnitEntity{
		BaseEntity: entity.BaseEntity{ID: model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix()},
		Name: model.Name,
	}
}
func ToListUnitEntity(models []*model.UnitModel) []*entity.UnitEntity {
	var entities []*entity.UnitEntity
	for _, unitModel := range models {
		entities = append(entities, ToUnitEntity(unitModel))
	}
	return entities
}
