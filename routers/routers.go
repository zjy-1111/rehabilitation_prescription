package routers

import (
	"rehabilitation_prescription/middleware/jwt"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/routers/api"
	v1 "rehabilitation_prescription/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "rehabilitation_prescription/docs"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r := gin.Default()

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "hello world!",
		})
	})
	r.GET("/auth", api.GetAuth)
	r.POST("/auth", api.AddAuth)
	r.PUT("/auth/", api.EditAuth)
	r.DELETE("/auth/:username", api.DeleteAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/reservation", v1.GetReservations)
		apiv1.POST("/reservation", v1.AddReservation)
		apiv1.PUT("/reservation", v1.EditReservation)
		apiv1.DELETE("/reservation", v1.DeleteReservation)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
