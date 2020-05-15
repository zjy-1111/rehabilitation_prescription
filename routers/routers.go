package routers

import (
	"net/http"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/routers/api"
	v1 "rehabilitation_prescription/routers/api/v1"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "rehabilitation_prescription/docs"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// r := gin.Default()
	r.Use(Cors())

	gin.SetMode(setting.ServerSetting.RunMode)

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "200",
			"message": "hello world!",
		})
	})

	r.GET("/user/:id", api.GetUserByID)
	r.GET("/users/:type", api.GetUsers)
	r.POST("/admin/login", api.AdminLogin)
	r.POST("/user/login", api.UserLogin)
	r.POST("/user", api.AddUser)
	r.PUT("/user/:id", api.EditUser)
	r.DELETE("/user/:id", api.DeleteUser)

	r.POST("/upload_image", api.UploadImage)
	r.POST("/upload_video", api.UploadVideo)

	apiv1 := r.Group("/api/v1")
	//apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/patients", v1.GetPatients)
		apiv1.GET("/doctors", v1.GetDoctors)
		apiv1.POST("/appoint", v1.AddAppointment)
		apiv1.PUT("/appoint/:id", v1.EditAppointment)
		apiv1.DELETE("/patient/:id", v1.DelAppointment)

		// 体测相关
		apiv1.GET("/examination", v1.GetExaminationList)
		apiv1.POST("/examination", v1.AddExamination)
		apiv1.DELETE("/examination/:id", v1.DelExamination)

		// 处方相关
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

		apiv1.GET("/training_video/:id", v1.GetTrainingVideo)
		apiv1.GET("/training_videos", v1.GetTrainingVideos)
		apiv1.POST("/training_video", v1.AddTrainingVideo)
		apiv1.PUT("/training_video/:id", v1.EditTrainingVideo)
		apiv1.DELETE("/training_video/:id", v1.DelTrainingVideo)

		//apiv1.GET("/prescription_video", v1.GetRPrescriptionTrainings)
		apiv1.POST("/prescription_video", v1.AddRPrescriptionTraining)
		apiv1.PUT("/prescription_video/:id", v1.EditRPrescriptionTraining)
		apiv1.DELETE("/prescription_video/:id", v1.DelRPrescriptionTraining)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
