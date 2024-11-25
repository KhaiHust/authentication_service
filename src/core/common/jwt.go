package common

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/KhaiHust/authen_service/core/entity"
	"github.com/KhaiHust/authen_service/core/entity/dto"
	"github.com/KhaiHust/authen_service/core/properties"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

func GenerateToken(userEntity *entity.UserEntity, props *properties.TokenProperties) (string, error) {
	// Parse the private key
	block, _ := pem.Decode([]byte(props.PrivateKey))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("invalid RSA private key format")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Create claims
	claims := dto.ClaimTokenDto{
		Email:  userEntity.Email,
		UserId: userEntity.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userEntity.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(props.TokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Sign the token with the private key
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userEntity *entity.UserEntity, props *properties.TokenProperties) (string, error) {
	// Parse the RSA private key
	block, _ := pem.Decode([]byte(props.PrivateKey))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("invalid private key format")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	// Set up claims for the refresh token
	claims := dto.ClaimTokenDto{
		UserId: userEntity.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatInt(userEntity.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(props.RefreshTokenExpired) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Generate the token with RSA signing method
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
