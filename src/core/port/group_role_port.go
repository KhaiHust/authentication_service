package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
)

type IGroupRolePort interface {
	GetRoleByCode(ctx context.Context, code string) (*entity.GroupRoleEntity, error)
}