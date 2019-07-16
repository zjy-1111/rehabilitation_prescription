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

func GetReportPrescriptions(c *gin.Context) {
	valid := validation.Validation{}

	reportID := -1
	if arg := c.PostForm("report_id"); arg != "" {
		reportID = com.StrTo(arg).MustInt()
		valid.Min(reportID, 1, "report_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	reportPrescriptionServ := services.ReportPrescription{
		PatientReportID: reportID,
		PageNum:         util.GetPage(c),
		PageSize:        setting.AppSetting.PageSize,
	}

	total, err := reportPrescriptionServ.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_REPORTPRESCRIPTION_FAIL, nil)
		return
	}

	prescriptions, err := reportPrescriptionServ.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_REPORTPRESCRIPTIONS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = prescriptions
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddReportPrescriptionForm struct {
	PatientReportID int    `form:"patient_report_id" valid:"Required;Min(1)"`
	PrescriptionID  int    `form:"prescription_id" valid:"Required;Min(1)"`
	Remark          string `form:"remark" valid:"Required;MaxSize(255)"`
}

func AddReportPrescription(c *gin.Context) {
	form := AddReportPrescriptionForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	// TODO check reportID是否存在

	reportPrescriptionServ := services.ReportPrescription{
		PatientReportID: form.PatientReportID,
		PrescriptionID:  form.PrescriptionID,
		Remark:          form.Remark,
	}

	if err := reportPrescriptionServ.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_REPORTPRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditReportPrescriptionForm struct {
	ID              int    `form:"id" valid:"Required;Min(1)"`
	PatientReportID int    `form:"patient_report_id" valid:"Required;Min(1)"`
	PrescriptionID  int    `form:"prescription_id" valid:"Required;Min(1)"`
	Remark          string `form:"remark" valid:"Required;MaxSize(255)"`
}

func EditReportPrescription(c *gin.Context) {
	form := EditReportPrescriptionForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	reportPrescriptionServ := services.ReportPrescription{
		ID:              form.ID,
		PatientReportID: form.PatientReportID,
		PrescriptionID:  form.PrescriptionID,
		Remark:          form.Remark,
	}
	exists, err := reportPrescriptionServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_REPORTPRESCRIPTION_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_REPORTPRESCRIPTION, nil)
		return
	}

	err = reportPrescriptionServ.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_REPORTPRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelReportPrescription(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	reportPrescriptionServ := services.ReportPrescription{ID: id}
	exist, err := reportPrescriptionServ.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_REPORTPRESCRIPTION_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_REPORTPRESCRIPTION, nil)
		return
	}

	err = reportPrescriptionServ.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_REPORTPRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
