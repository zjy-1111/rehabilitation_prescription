package v1

import (
	"net/http"
	"rehabilitation_prescription/handlers"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"
	"strconv"
	"strings"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetPrescriptions(c *gin.Context) {
	valid := validation.Validation{}

	patientID := -1
	if arg := c.Query("patient_id"); arg != "" {
		patientID = com.StrTo(arg).MustInt()
		valid.Min(patientID, 1, "patient_id")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	//authService := services.User{
	//	ID:       patientID,
	//	UserType: "2",
	//}
	//checkValidAuth(c, authService)

	h := handlers.NewPrescriptionHandler(patientID, util.GetPage(c), setting.AppSetting.PageSize)

	prescriptions, err := h.GetPrescriptions()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_PRESCRIPTIONS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = prescriptions
	data["total"] = len(prescriptions)
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type AddPrescriptionForm struct {
	Title          string `form:"title" valid:"Required;MaxSize(255)"`
	PatientID      int    `form:"patient_id" valid:"Required;Min(1)"`
	Desc           string `form:"desc" valid:"Required;MaxSize(255)"`
	TrainingIDList string `form:"training_id_list"`
}

func AddPrescription(c *gin.Context) {
	form := AddPrescriptionForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	//authService := services.User{
	//	ID:       form.DoctorID,
	//	UserType: "2",
	//}
	//checkValidAuth(c, authService)

	lists := strings.Split(form.TrainingIDList, ",")
	trainingIDList := make([]int, len(lists))
	for i := 0; i < len(lists); i++ {
		trainingIDList[i], _ = strconv.Atoi(lists[i])
	}

	s := services.Prescription{
		Title:          form.Title,
		PatientID:      form.PatientID,
		Desc: form.Desc,
		TrainingIDList: trainingIDList,
	}

	if err := s.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_PRESCRIPTION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

type EditPrescriptionForm struct {
	ID        int    `form:"id" valid:"Required;Min(1)"`
	PatientID int    `form:"patient_id" valid:"Required;Min(1)"`
	Title     string `form:"title" valid:"Required;MaxSize(255)"`
}

func EditPrescription(c *gin.Context) {
	form := EditPrescriptionForm{ID: com.StrTo(c.Param("id")).MustInt()}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	//authService := services.User{
	//	ID:       form.PatientID,
	//	UserType: "2",
	//}
	//checkValidAuth(c, authService)

	prescriptionService := services.Prescription{
		ID:        form.ID,
		PatientID: form.PatientID,
		Title:     form.Title,
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

	prescriptionService := services.Prescription{ID: id}
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
