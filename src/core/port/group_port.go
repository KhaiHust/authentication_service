package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IGroupPort interface {
	CreateGroup(ctx context.Context, tx *gorm.DB, group *entity.GroupEntity) (*entity.GroupEntity, error)
	GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error)
}
