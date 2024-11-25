package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func ToGroupMemberModel(groupMember *entity.GroupMemberEntity) *model.GroupMemberModel {
	return &model.GroupMemberModel{
		BaseModel: model.BaseModel{
			ID: groupMember.ID,
		},
		GroupID: groupMember.GroupID,
		UserID:  groupMember.UserID,
		RoleID:  groupMember.RoleID,
	}
}
func ToGroupMemberEntity(groupMemberModel *model.GroupMemberModel) *entity.GroupMemberEntity {
	return &entity.GroupMemberEntity{
		BaseEntity: entity.BaseEntity{
			ID:        groupMemberModel.ID,
			CreatedAt: groupMemberModel.CreatedAt.Unix(),
			UpdatedAt: groupMemberModel.UpdatedAt.Unix(),
		},
		GroupID: groupMemberModel.GroupID,
		UserID:  groupMemberModel.UserID,
		RoleID:  groupMemberModel.RoleID,
	}
}
