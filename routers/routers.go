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
		apiv1.GET("/appoint", v1.GetPatients)
		apiv1.POST("/appoint", v1.AddAppointment)
		apiv1.PUT("/appoint/:id", v1.EditAppointment)
		apiv1.DELETE("/appoint/:id", v1.DelAppointment)

		apiv1.GET("/prescription", v1.GetPrescriptions)
		apiv1.POST("/prescription", v1.AddPrescription)
		apiv1.PUT("/prescription/:id", v1.EditPrescription)
		apiv1.DELETE("/prescription/:id", v1.DelPrescription)

		apiv1.GET("/patient_report", v1.GetPatientReport)
		apiv1.POST("/patient_report", v1.AddPatientReport)
		apiv1.PUT("/patient_report/:id", v1.EditPatientReport)
		apiv1.DELETE("/patient_report/:id", v1.DelPatientReport)

		apiv1.GET("/report_prescription", v1.GetReportPrescriptions)
		apiv1.POST("/report_prescription", v1.AddReportPrescription)
		apiv1.PUT("/report_prescription/:id", v1.EditReportPrescription)
		apiv1.DELETE("/report_prescription/:id", v1.DelReportPrescription)

		apiv1.GET("/training_video", v1.GetTrainingVideos)
		apiv1.POST("/training_video", v1.AddTrainingVideo)
		apiv1.PUT("/training_video/:id", v1.EditTrainingVideo)
		apiv1.DELETE("/training_video/:id", v1.DelTrainingVideo)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
