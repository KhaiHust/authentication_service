package postgres

import (
	"context"
	"errors"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/port"
	"github.com/golibs-starter/golib/log"
	"gorm.io/gorm"
)

type UserRepositoryAdapter struct {
	base
}

func (u *UserRepositoryAdapter) GetUserByEmail(ctx *context.Context, email string) (*entity.UserEntity, error) {
	user := entity.UserEntity{}
	err := u.db.WithContext(*ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Error(ctx, "Error when get user by email ", err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New(common.ErrUserNotFound)
		}
		return nil, err
	}
	return &user, nil
}

func NewUserRepositoryAdapter(db *gorm.DB) port.IUserPort {
	return &UserRepositoryAdapter{
		base: base{db},
	}
}
