package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type GroupRepositoryAdapter struct {
	base
}

func (g *GroupRepositoryAdapter) CreateGroup(ctx context.Context, tx *gorm.DB, group *entity.GroupEntity) (*entity.GroupEntity, error) {
	groupModel := mapper.ToGroupModel(group)
	if err := tx.WithContext(ctx).Create(groupModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToGroupEntity(groupModel), nil
}

func NewGroupRepositoryAdapter(db *gorm.DB) port.IGroupPort {
	return &GroupRepositoryAdapter{
		base: base{db: db},
	}
}
