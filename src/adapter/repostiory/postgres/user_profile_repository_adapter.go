package postgres

import (
	"context"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"gorm.io/gorm"
)

type UserProfileRepositoryAdapter struct {
	*base
}

func (u UserProfileRepositoryAdapter) CreateNewProfile(ctx context.Context, txn *gorm.DB, userProfile *entity.UserProfileEntity) (*entity.UserProfileEntity, error) {
	userProfileModel := mapper.ToUserProfileModel(userProfile)
	if err := txn.WithContext(ctx).Create(userProfileModel).Error; err != nil {
		return nil, err
	}
	return mapper.ToUserProfileEntity(userProfileModel), nil
}

func (u UserProfileRepositoryAdapter) GetUserProfilesByUserIDs(ctx context.Context, userIDs []int64) ([]*entity.UserProfileEntity, error) {
	var userProfileModels []*model.UserProfileModel

	if err := u.db.WithContext(ctx).Model(&model.UserProfileModel{}).Where("user_id IN ?", userIDs).Find(&userProfileModels).Error; err != nil {
		return nil, err
	}
	return mapper.ToListUserProfileEntity(userProfileModels), nil
}

func NewUserProfileRepositoryAdapter(db *gorm.DB) port.IUserProfilePort {
	return &UserProfileRepositoryAdapter{base: &base{db: db}}
}
