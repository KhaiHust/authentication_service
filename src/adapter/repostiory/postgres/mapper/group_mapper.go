package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToGroupModel(group *entity.GroupEntity) *model.GroupModel {
	return &model.GroupModel{
		BaseModel: model.BaseModel{
			ID: group.ID,
		},
		Name:        group.Name,
		Description: group.Description,
	}
}
func ToGroupEntity(groupModel *model.GroupModel) *entity.GroupEntity {
	return &entity.GroupEntity{
		BaseEntity: entity.BaseEntity{
			ID:        groupModel.ID,
			CreatedAt: groupModel.CreatedAt.Unix(),
			UpdatedAt: groupModel.UpdatedAt.Unix(),
		},
		Name:        groupModel.Name,
		Description: groupModel.Description,
	}
}
