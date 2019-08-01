package v1

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetTrainingVideo(c *gin.Context) {
	s := services.TrainingVideo{
		ID: com.StrTo(c.Param("id")).MustInt(),
	}

	video, err := s.GetVideoByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.	ERROR_GET_TRAININGVIDEO_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["video"] = video
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

func GetTrainingVideos(c *gin.Context) {
	traningVideoServ := services.TrainingVideo{
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := traningVideoServ.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_TRAININGVIDEO_FAIL, nil)
		return
	}

	videos, err := traningVideoServ.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_TRAININGVIDEOS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = videos
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddTrainingVideoForm struct {
	VideoUrl string `json:"video_url" valid:"Required;MaxSize(255)"`
	CoverUrl string `json:"cover_url"`
	Duration int    `json:"duration"`
}

func AddTrainingVideo(c *gin.Context) {
	form := AddTrainingVideoForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	trainingVideoServ := services.TrainingVideo{
		VideoUrl: form.VideoUrl,
		CoverUrl: form.CoverUrl,
		Duration: form.Duration,
	}

	if err := trainingVideoServ.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_TRAININGVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditTrainingVideoForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	VideoUrl string `json:"video_url" valid:"Required;MaxSize(255)"`
	CoverUrl string `json:"cover_url" valid:"MaxSize(255)"`
	Duration int    `json:"duration" valid:"Min(1)"`
}

func EditTrainingVideo(c *gin.Context) {
	form := EditTrainingVideoForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	trainingVideoServ := services.TrainingVideo{
		ID:       form.ID,
		VideoUrl: form.VideoUrl,
		CoverUrl: form.CoverUrl,
		Duration: form.Duration,
	}
	exists, err := trainingVideoServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_TRAININGVIDEO_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_TRAININGVIDEO, nil)
		return
	}

	err = trainingVideoServ.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_TRAININGVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelTrainingVideo(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	trainingVideoServ := services.TrainingVideo{ID: id}
	exist, err := trainingVideoServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_TRAININGVIDEO_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_TRAININGVIDEO, nil)
		return
	}

	err = trainingVideoServ.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_TRAININGVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
