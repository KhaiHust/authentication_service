package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IUserPort interface {
	GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error)
	SaveUser(ctx *context.Context, userEntity *entity.UserEntity, tx *gorm.DB) (*entity.UserEntity, error)
	UpdateUser(ctx *context.Context, userEntity *entity.UserEntity, tx *gorm.DB) (*entity.UserEntity, error)
}
