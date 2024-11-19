package port

import (
	"context"
	"github.com/KhaiHust/authen_service/core/entity"
	"gorm.io/gorm"
)

type IRefreshTokenPort interface {
	SaveRefreshToken(ctx *context.Context, refreshToken *entity.RefreshTokenEntity, tx *gorm.DB) (*entity.RefreshTokenEntity, error)
}
