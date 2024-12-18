package middleware

import (
	"github.com/KhaiHust/authen_service/core/constant"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPagingParams(c *gin.Context) (int, int) {
	pageSize, err := strconv.ParseInt(c.Query("pageSize"), 10, 32)
	if err != nil {
		pageSize = constant.DefaultPageSize
	}
	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	if err != nil {
		page = constant.DefaultOffset
	}
	return int(pageSize), int(page)
}
