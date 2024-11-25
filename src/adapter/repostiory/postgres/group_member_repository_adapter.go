package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type GroupMemberRepositoryAdapter struct {
	base
}

func (g GroupMemberRepositoryAdapter) CreateGroupMember(ctx context.Context, tx *gorm.DB, groupMember *entity.GroupMemberEntity) (*entity.GroupMemberEntity, error) {
	groupMemberModel := mapper.ToGroupMemberModel(groupMember)
	if err := tx.WithContext(ctx).Create(groupMemberModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToGroupMemberEntity(groupMemberModel), nil
}

func NewGroupMemberRepositoryAdapter(db *gorm.DB) port.IGroupMemberPort {
	return &GroupMemberRepositoryAdapter{
		base: base{db: db},
	}
}
