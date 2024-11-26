package postgres

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type GroupRepositoryAdapter struct {
	base
}

func (g *GroupRepositoryAdapter) GetGroupById(ctx context.Context, groupID int64) (*entity.GroupEntity, error) {
	groupModel := &model.GroupModel{}
	if err := g.db.WithContext(ctx).Model(&model.GroupModel{}).Where("id = ?", groupID).First(groupModel).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}
	return mapper.ToGroupEntity(groupModel), nil
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
