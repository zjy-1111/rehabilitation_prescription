package v1

import (
	"net/http"
	"rehabilitation_prescription/handlers"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

type Query struct {
	DoctorID    int    `form:"doctor_id"`
	PatientName string `form:"patient_name"`
}

// @summary 病人列表
// @accept json
// @produce json
// @param doctor_id query int true "医生id"
// @param page query int true "分页"
// @param patient_name string false "按病人名称查询"
// @success 200 {object} models.Response
// @router api/v1/patients [get]
func GetPatients(c *gin.Context) {
	var q Query
	httpCode, errCode := app.BindAndValid(c, &q)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	//authService := services.User{
	//	ID:       doctorID,
	//	UserType: "2",
	//}
	//checkValidAuth(c, authService)

	h := handlers.NewPatientsHandler(q.DoctorID, util.GetPage(c), setting.AppSetting.PageSize)

	patients, err := h.GetPatients(q.PatientName)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_PATIENTS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = patients
	data["total"] = len(patients)
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddAppointmentForm struct {
	ID        int `form:"id" valid:"Required;Min(1)"`
	DoctorID  int `form:"doctor_id" valid:"Required;Min(1)"`
	PatientID int `form:"patient_id" valid:"Required;Min(1)"`
}

func AddAppointment(c *gin.Context) {
	form := AddAppointmentForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := services.User{
		ID:       form.DoctorID,
		UserType: "2",
	}
	checkValidAuth(c, authService)
	authService = services.User{
		ID:       form.PatientID,
		UserType: "1",
	}
	checkValidAuth(c, authService)

	appointService := services.Appointment{
		DoctorID:  form.DoctorID,
		PatientID: form.PatientID,
	}

	if err := appointService.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_APPOINTMENT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditAppointmentForm struct {
	ID        int `form:"id" valid:"Required;Min(1)"`
	DoctorID  int `form:"doctor_id" valid:"Required;Min(1)"`
	PatientID int `form:"patient_id" valid:"Required;Min(1)"`
}

func EditAppointment(c *gin.Context) {
	form := EditAppointmentForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := services.User{
		ID:       form.DoctorID,
		UserType: "2",
	}
	checkValidAuth(c, authService)
	authService = services.User{
		ID:       form.PatientID,
		UserType: "1",
	}
	checkValidAuth(c, authService)

	appointService := services.Appointment{
		ID:        form.ID,
		DoctorID:  form.DoctorID,
		PatientID: form.PatientID,
	}
	exists, err := appointService.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APPOINTMENT_FAIL, nil)
		return
	}
	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_APPOINTMENT, nil)
		return
	}

	err = appointService.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_APPOINTMENT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelAppointment(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	appointService := services.Appointment{ID: id}
	exist, err := appointService.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_CHECK_EXIST_APPOINTMENT_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_APPOINTMENT, nil)
		return
	}

	err = appointService.Del()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_APPOINTMENT_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func checkValidAuth(c *gin.Context, a services.User) {
	exist, err := a.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}
	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_AUTH, nil)
		return
	}
}
