package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToGroupRoleEntity(groupRoleModel *model.GroupRoleModel) *entity.GroupRoleEntity {
	return &entity.GroupRoleEntity{
		BaseEntity: entity.BaseEntity{ID: groupRoleModel.ID,
			CreatedAt: groupRoleModel.CreatedAt.Unix(),
			UpdatedAt: groupRoleModel.UpdatedAt.Unix(),
		},
		Code: groupRoleModel.Code,
		Name: groupRoleModel.Name,
	}
}
func ToListGroupRoleEntity(groupRoleModels []*model.GroupRoleModel) []*entity.GroupRoleEntity {
	groupRoleEntities := make([]*entity.GroupRoleEntity, 0)
	for _, groupRoleModel := range groupRoleModels {
		groupRoleEntities = append(groupRoleEntities, ToGroupRoleEntity(groupRoleModel))
	}
	return groupRoleEntities
}
