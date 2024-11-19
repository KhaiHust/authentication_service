package middleware

import (
	"github.com/KhaiHust/authen_service/core/common"
	"github.com/KhaiHust/authen_service/public/apihelper"
	"github.com/gin-gonic/gin"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				apihelper.AbortErrorHandle(c, common.GeneralServiceUnavailable)
			}
		}()
		c.Next()
	}
}
