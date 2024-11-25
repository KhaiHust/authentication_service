package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type RefreshTokenRepositoryAdapter struct {
	base
}

func (r *RefreshTokenRepositoryAdapter) SaveRefreshToken(ctx *context.Context, refreshToken *entity.RefreshTokenEntity, tx *gorm.DB) (*entity.RefreshTokenEntity, error) {
	refreshTokenModel := mapper.ToRefreshTokenModel(refreshToken)
	if err := tx.WithContext(*ctx).Create(refreshTokenModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToRefreshTokenEntity(refreshTokenModel), nil
}

func NewRefreshTokenRepositoryAdapter(db *gorm.DB) port.IRefreshTokenPort {
	return &RefreshTokenRepositoryAdapter{base: base{db: db}}
}
