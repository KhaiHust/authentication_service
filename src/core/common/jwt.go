package common

import (
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/properties"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userEntity *entity.UserEntity, props *properties.TokenProperties) (string, error) {
	claims := dto.ClaimTokenDto{
		Email:  userEntity.Email,
		UserId: userEntity.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(props.TokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(props.PrivateKey))
}
func GenerateRefreshToken(userEntity *entity.UserEntity, props *properties.TokenProperties) (string, error) {
	claims := dto.ClaimTokenDto{
		UserId: userEntity.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(props.RefreshTokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(props.PrivateKey))
}
