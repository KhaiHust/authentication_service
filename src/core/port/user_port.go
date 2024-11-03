package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
)

type IUserPort interface {
	GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error)
}
