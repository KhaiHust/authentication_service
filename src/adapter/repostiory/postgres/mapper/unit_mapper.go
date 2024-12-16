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
