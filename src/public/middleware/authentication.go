package middleware

import (
	"errors"
	"fmt"
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/golibs-starter/golib-security/web/config"
)

func GetInfoFromToken(props *config.JwtSecurityProperties) gin.HandlerFunc {
	return func(c *gin.Context) {
		jwtKeyFunc, err := getJwtPublicKeyFunc(props)
		if err != nil {
			apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
			return
		}
		jwtExtractor := request.AuthorizationHeaderExtractor
		jwtParser := request.WithParser(jwt.NewParser(jwt.WithValidMethods([]string{props.Algorithm})))
		token, err := request.ParseFromRequest(c.Request, jwtExtractor, jwtKeyFunc, jwtParser)
		if err != nil {
			apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			apihelper.AbortErrorHandle(c, common.GeneralUnauthorized)
			return
		}
		c.Set(constant.CLAIMS, claims)
		c.Next()

	}

}
func getJwtPublicKeyFunc(props *config.JwtSecurityProperties) (func(token *jwt.Token) (interface{}, error), error) {
	if len(props.PublicKey) == 0 {
		return nil, errors.New("jwt public key must be defined when using jwt authentication")
	}
	var err error
	var publicKey interface{}
	if len(props.PublicKey) > 0 {
		if props.IsAlgEs() {
			publicKey, err = jwt.ParseECPublicKeyFromPEM([]byte(props.PublicKey))
		} else if props.IsAlgRs() {
			publicKey, err = jwt.ParseRSAPublicKeyFromPEM([]byte(props.PublicKey))
		} else {
			err = fmt.Errorf("unsupported jwt algorithm: [%v], required startswith RS or ES",
				props.Algorithm)
		}
		if err != nil {
			return nil, err
		}
	}
	return func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	}, nil
}
