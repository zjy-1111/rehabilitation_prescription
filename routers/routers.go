package routers

import (
	"rehabilitation_prescription/pkg/setting"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	// r := gin.New()
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	r := gin.Default()

	gin.SetMode(setting.RunMode)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "hello world!",
		})
	})

	return r
}
