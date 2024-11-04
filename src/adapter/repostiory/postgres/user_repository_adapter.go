package postgres

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/mapper"
	"github.com/KhaiHust/authen_service/adapter/repostiory/postgres/model"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
	"gorm.io/gorm"
)

type UserRepositoryAdapter struct {
	base
}

func (u *UserRepositoryAdapter) SaveUser(ctx *context.Context, userEntity *entity.UserEntity, tx *gorm.DB) (*entity.UserEntity, error) {
	userModel := mapper.EntityToUserModel(userEntity)
	if err := tx.WithContext(*ctx).Create(userModel).Error; err != nil {
		log.Error(ctx, "Error when save user ", err)
		return nil, err
	}
	return mapper.ModelToUserEntity(userModel), nil
}

func (u *UserRepositoryAdapter) GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error) {
	user := &model.UserModel{}

	if err := u.db.WithContext(*ctx).Where("email = ?", email).First(user).Error; err != nil {
		log.Error(ctx, "Error when get user by email ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(common.ErrUserNotFound)
		}
		return nil, err
	}
	return mapper.ModelToUserEntity(user), nil
}

func NewUserRepositoryAdapter(db *gorm.DB) port.IUserPort {
	return &UserRepositoryAdapter{
		base: base{db},
	}
}
