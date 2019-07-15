package app

import (
	"rehabilitation_prescription/pkg/e"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, gin.H{
		"code": httpCode,
		"msg":  e.GetMsg(errCode),
		"data": data,
	})

	return
}
