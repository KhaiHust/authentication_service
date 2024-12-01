package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IGroupMemberPort interface {
	CreateGroupMember(ctx context.Context, tx *gorm.DB, groupMember *entity.GroupMemberEntity) (*entity.GroupMemberEntity, error)
	GetGroupMemberByGroupIDAndUserID(ctx context.Context, groupID int64, userID int64) (*entity.GroupMemberEntity, error)
	GetListMemberByGroupID(ctx context.Context, groupID int64) ([]*entity.GroupMemberEntity, error)
}
