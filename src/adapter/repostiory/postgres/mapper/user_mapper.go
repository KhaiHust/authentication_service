package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
)

func EntityToUserModel(userEntity *entity.UserEntity) *model.UserModel {
	if userEntity == nil {
		return nil
	}
	return &model.UserModel{
		BaseModel: model.BaseModel{
			ID: userEntity.ID,
		},
		Name:       userEntity.Name,
		Email:      userEntity.Email,
		Password:   userEntity.Password,
		DeviceID:   userEntity.DeviceID,
		IsActive:   userEntity.IsActive,
		IsVerified: userEntity.IsVerified,
		Language:   userEntity.Language,
		Timezone:   userEntity.Timezone,
	}
}
func ModelToUserEntity(userModel *model.UserModel) *entity.UserEntity {
	if userModel == nil {
		return nil
	}
	return &entity.UserEntity{
		BaseEntity: entity.BaseEntity{
			ID:        userModel.ID,
			CreatedAt: userModel.CreatedAt.Unix(),
			UpdatedAt: userModel.UpdatedAt.Unix(),
		},
		Name:       userModel.Name,
		Email:      userModel.Email,
		Password:   userModel.Password,
		DeviceID:   userModel.DeviceID,
		IsActive:   userModel.IsActive,
		IsVerified: userModel.IsVerified,
		Language:   userModel.Language,
		Timezone:   userModel.Timezone,
	}
}
