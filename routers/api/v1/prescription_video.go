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

func GetPrescriptionVideos(c *gin.Context) {
	valid := validation.Validation{}

	prescriptionID := -1
	if arg := c.PostForm("prescription_id"); arg != "" {
		prescriptionID = com.StrTo(arg).MustInt()
		valid.Min(prescriptionID, 1, "prescription_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	prescriptionVideoServ := services.PrescriptionVideo{
		PageNum:        util.GetPage(c),
		PageSize:       setting.AppSetting.PageSize,
		PrescriptionID: prescriptionID,
	}

	total, err := prescriptionVideoServ.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}

	videos, err := prescriptionVideoServ.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_PRESCRIPTIONVIDEOS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = videos
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddPrescriptionVideoForm struct {
	PrescriptionID int `form:"prescription_id" valid:"Required;Min(1)"`
	VideoID        int `form:"video_id" valid:"Required;Min(1)"`
}

func AddPrescriptionVideo(c *gin.Context) {
	form := AddPrescriptionVideoForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	prescriptionVideoServ := services.PrescriptionVideo{
		PrescriptionID: form.PrescriptionID,
		VideoID:        form.VideoID,
	}

	if err := prescriptionVideoServ.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditPrescriptionVideoForm struct {
	ID             int `form:"id" valid:"Required;Min(1)"`
	PrescriptionID int `form:"prescription_id" valid:"Required;Min(1)"`
	VideoID        int `form:"video_id" valid:"Required;Min(1)"`
}

func EditPrescriptionVideo(c *gin.Context) {
	form := EditPrescriptionVideoForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	prescriptionVideoServ := services.PrescriptionVideo{
		ID:             form.ID,
		PrescriptionID: form.PrescriptionID,
		VideoID:        form.VideoID,
	}
	exists, err := prescriptionVideoServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PRESCRIPTIONVIDEO, nil)
		return
	}

	err = prescriptionVideoServ.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelPrescriptionVideo(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	prescriptionVideoServ := services.PrescriptionVideo{ID: id}
	exist, err := prescriptionVideoServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PRESCRIPTIONVIDEO, nil)
		return
	}

	err = prescriptionVideoServ.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_PRESCRIPTIONVIDEO_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
