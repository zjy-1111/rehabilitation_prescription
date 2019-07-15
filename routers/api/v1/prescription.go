package v1

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services/auth_service"
	"rehabilitation_prescription/services/prescription_service"
	"rehabilitation_prescription/util"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetPrescriptions(c *gin.Context) {
	valid := validation.Validation{}

	doctorID := -1
	if arg := c.PostForm("doctor_id"); arg != "" {
		doctorID = com.StrTo(arg).MustInt()
		valid.Min(doctorID, 1, "doctor_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authService := auth_service.Auth{
		ID:       doctorID,
		UserType: 2,
	}
	checkValidAuth(c, authService)

	prescriptionService := prescription_service.Prescription{
		DoctorID: doctorID,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	total, err := prescriptionService.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_PATIENT_FAIL, nil)
		return
	}

	prescriptions, err := prescriptionService.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_PATIENTS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["lists"] = prescriptions
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddPrescriptionForm struct {
	DoctorID int    `form:"doctor_id" valid:"Required;Min(1)"`
	Desc     string `form:"desc" valid:"Required;MaxSize(255)"`
}

func AddPrescription(c *gin.Context) {
	form := AddPrescriptionForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{
		ID:       form.DoctorID,
		UserType: 2,
	}
	checkValidAuth(c, authService)

	prescriptionService := prescription_service.Prescription{
		DoctorID: form.DoctorID,
		Desc:     form.Desc,
	}

	if err := prescriptionService.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_PRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditPrescriptionForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	DoctorID int    `form:"doctor_id" valid:"Required;Min(1)"`
	Desc     string `form:"desc" valid:"Required;MaxSize(255)"`
}

func EditPrescription(c *gin.Context) {
	form := EditPrescriptionForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{
		ID:       form.DoctorID,
		UserType: 2,
	}
	checkValidAuth(c, authService)

	prescriptionService := prescription_service.Prescription{
		ID:       form.ID,
		DoctorID: form.DoctorID,
		Desc:     form.Desc,
	}
	exists, err := prescriptionService.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRESCRIPTION_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PRESCRIPTION, nil)
		return
	}

	err = prescriptionService.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_PRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelPrescription(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	prescriptionService := prescription_service.Prescription{ID: id}
	exist, err := prescriptionService.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_PRESCRIPTION_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_PRESCRIPTION, nil)
		return
	}

	err = prescriptionService.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_PRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
