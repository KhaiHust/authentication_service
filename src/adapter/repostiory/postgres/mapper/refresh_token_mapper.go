package mapper

import (
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/entity"
	"time"
)

func ToRefreshTokenModel(entity *entity.RefreshTokenEntity) *model.RefreshTokenModel {
	return &model.RefreshTokenModel{
		BaseModel:    model.BaseModel{ID: entity.ID},
		UserId:       entity.UserId,
		RefreshToken: entity.RefreshToken,
		ExpiredAt:    time.Unix(entity.ExpiredAt, 0),
		IpAddress:    entity.IpAddress,
		UserAgent:    entity.UserAgent,
	}
}
func ToRefreshTokenEntity(model *model.RefreshTokenModel) *entity.RefreshTokenEntity {
	return &entity.RefreshTokenEntity{
		BaseEntity: entity.BaseEntity{ID: model.ID,
			CreatedAt: model.CreatedAt.Unix(),
			UpdatedAt: model.UpdatedAt.Unix(),
		},
		UserId:       model.UserId,
		RefreshToken: model.RefreshToken,
		ExpiredAt:    model.ExpiredAt.Unix(),
		IpAddress:    model.IpAddress,
		UserAgent:    model.UserAgent,
	}
}
