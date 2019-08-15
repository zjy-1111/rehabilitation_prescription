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

func GetPatientReport(c *gin.Context) {
	valid := validation.Validation{}

	patientID := -1
	if arg := c.PostForm("patient_id"); arg != "" {
		patientID = com.StrTo(arg).MustInt()
		valid.Min(patientID, 1, "patient_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := services.User{
		ID:       patientID,
		UserType: "1",
	}
	checkValidAuth(c, authService)

	reportServ := services.PatientReport{
		PatientID: patientID,
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}

	total, err := reportServ.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_PATIENTREPORT_FAIL, nil)
		return
	}

	reports, err := reportServ.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_PATIENTREPORTS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = reports
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddPatientReportForm struct {
	PatientID     int    `form:"patient_id" valid:"Required;Min(1)"`
	BodyType      string `form:"body_type" valid:"Required;MaxSize(32)"`
	Height        string `form:"height" valid:"Required;MaxSize(16)"`
	Weight        string `form:"weight" valid:"Required;MaxSize(16)"`
	Waist         string `form:"waist" valid:"Required;MaxSize(16)"`
	BloodPressure string `form:"blood_pressure" valid:"Required;MaxSize(16)"`
}

func AddPatientReport(c *gin.Context) {
	form := AddPatientReportForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := services.User{
		ID:       form.PatientID,
		UserType: "1",
	}
	checkValidAuth(c, authService)

	reportServ := services.PatientReport{
		PatientID:     form.PatientID,
		BodyType:      form.BodyType,
		Height:        form.Height,
		Weight:        form.Weight,
		Waist:         form.Waist,
		BloodPressure: form.BloodPressure,
	}

	if err := reportServ.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_PATIENTREPORT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditPatientReportForm struct {
	ID            int    `form:"id" valid:"Required;Min(1)"`
	PatientID     int    `form:"patient_id" valid:"Required;Min(1)"`
	BodyType      string `form:"body_type" valid:"Required;MaxSize(32)"`
	Height        string `form:"height" valid:"Required;MaxSize(16)"`
	Weight        string `form:"weight" valid:"Required;MaxSize(16)"`
	Waist         string `form:"waist" valid:"Required;MaxSize(16)"`
	BloodPressure string `form:"blood_pressure" valid:"Required;MaxSize(16)"`
}

func EditPatientReport(c *gin.Context) {
	form := EditPatientReportForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := services.User{
		ID:       form.PatientID,
		UserType: "1",
	}
	checkValidAuth(c, authService)

	reportServ := services.PatientReport{
		ID:            form.ID,
		PatientID:     form.PatientID,
		BodyType:      form.BodyType,
		Height:        form.Height,
		Weight:        form.Weight,
		Waist:         form.Waist,
		BloodPressure: form.BloodPressure,
	}
	exists, err := reportServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PATIENTREPORT_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PATIENTREPORT, nil)
		return
	}

	err = reportServ.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_PATIENTREPORT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelPatientReport(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	reportServ := services.PatientReport{ID: id}
	exist, err := reportServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PATIENTREPORT_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PATIENTREPORT, nil)
		return
	}

	err = reportServ.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_PATIENTREPORT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
