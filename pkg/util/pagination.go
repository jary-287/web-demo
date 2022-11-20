package util

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jary-287/web-demo/pkg/setting"
)

func GetPage(c *gin.Context) (result int) {
	page, _ := strconv.Atoi(c.Query("page"))
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}
	return
}
