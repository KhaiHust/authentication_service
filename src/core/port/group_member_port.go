package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IGroupMemberPort interface {
	CreateGroupMember(ctx context.Context, tx *gorm.DB, groupMember *entity.GroupMemberEntity) (*entity.GroupMemberEntity, error)
	GetGroupMemberByGroupIdAndUserId(ctx context.Context, groupId, userId int64) (*entity.GroupMemberEntity, error)
}
