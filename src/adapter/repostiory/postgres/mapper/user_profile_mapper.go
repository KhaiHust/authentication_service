package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"time"
)

func ToUserProfileModel(userProfile *entity.UserProfileEntity) *model.UserProfileModel {
	return &model.UserProfileModel{
		BaseModel: model.BaseModel{
			ID: userProfile.ID,
		},
		UserID:         userProfile.UserID,
		Email:          userProfile.Email,
		Name:           userProfile.Name,
		DoB:            time.Unix(userProfile.DoB, 0),
		AvatarImageUrl: userProfile.AvatarImageUrl,
	}
}

func ToUserProfileEntity(userProfileModel *model.UserProfileModel) *entity.UserProfileEntity {
	return &entity.UserProfileEntity{
		BaseEntity: entity.BaseEntity{
			ID:        userProfileModel.ID,
			CreatedAt: userProfileModel.CreatedAt.Unix(),
			UpdatedAt: userProfileModel.UpdatedAt.Unix(),
		},
		UserID:         userProfileModel.UserID,
		Email:          userProfileModel.Email,
		Name:           userProfileModel.Name,
		DoB:            userProfileModel.DoB.Unix(),
		AvatarImageUrl: userProfileModel.AvatarImageUrl,
	}
}

func ToListUserProfileEntity(userProfileModels []*model.UserProfileModel) []*entity.UserProfileEntity {
	userProfileEntities := make([]*entity.UserProfileEntity, 0)
	for _, userProfileModel := range userProfileModels {
		userProfileEntities = append(userProfileEntities, ToUserProfileEntity(userProfileModel))
	}
	return userProfileEntities
}
