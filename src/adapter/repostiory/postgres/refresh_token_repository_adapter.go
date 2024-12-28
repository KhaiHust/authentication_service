package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type RefreshTokenRepositoryAdapter struct {
	base
}

func (r *RefreshTokenRepositoryAdapter) DeleteTokenByUserID(ctx context.Context, userID int64) error {
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).Delete(&model.RefreshTokenModel{}).Error; err != nil {
		return err
	}
	return nil
}

func (r *RefreshTokenRepositoryAdapter) GetRefreshTokenByToken(ctx context.Context, refreshToken string) (*entity.RefreshTokenEntity, error) {
	refreshTokenModel := &model.RefreshTokenModel{}
	if err := r.db.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(refreshTokenModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToRefreshTokenEntity(refreshTokenModel), nil
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
