package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type GroupRoleRepositoryAdapter struct {
	base
}

func (g *GroupRoleRepositoryAdapter) GetRoleByCode(ctx context.Context, code string) (*entity.GroupRoleEntity, error) {
	var roleModel model.GroupRoleModel
	if err := g.db.WithContext(ctx).Where("code = ?", code).First(&roleModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToGroupRoleEntity(&roleModel), nil
}

func NewGroupRoleRepositoryAdapter(db *gorm.DB) port.IGroupRolePort {
	return &GroupRoleRepositoryAdapter{
		base: base{db: db},
	}
}
