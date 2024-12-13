package middleware

import (
	"errors"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golibs-starter/golib/log"
)

func GetUserID(c *gin.Context) (int64, error) {
	claims, ok := c.Get("claims")
	if !ok {
		log.Error(c, "error when getting claims from context")
		return 0, errors.New("error when getting claims from context")
	}
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		log.Error(c, "error when casting claims to map")

		return 0, errors.New("error when casting claims to map")
	}
	userID, ok := claimsMap[constant.CLAIM_USER_ID]
	if !ok {
		log.Error(c, "error when getting user id from claims")
		return 0, errors.New("error when getting user id from claims")
	}

	return int64(userID.(float64)), nil
}
