package v1

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services/reservation_service"
	"rehabilitation_prescription/util"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	reservationService := reservation_service.Reservation{
		Name:      c.PostForm("name"),
		Date:      com.StrTo(c.PostForm("Date")).MustInt(),
		DoctorID:  com.StrTo(c.PostForm("doctor_id")).MustInt(),
		CreatedBy: c.PostForm("created_by"),
		PageNum:   util.GetPage(c),
		PageSize:  setting.AppSetting.PageSize,
	}

	appG := app.Gin{c}
	reservations, err := reservationService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESERVATIONS_FAIL, nil)
		return
	}

	count, err := reservationService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_RESERVATION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]interface{}{
		"list":  reservations,
		"total": count,
	})
}

type AddReservationForm struct {
	Name      string `form:"name" valid:"Required;MaxSize(100)"`
	Date      int    `form:"date" valid:"Required;Min(1)"`
	PeriodID  int    `form:"period_id" valid:"Required;Min(1)"`
	DoctorID  int    `form:"doctor_id" valid:"Required;Min(1)"`
	Address   string `form:"address" valid:"Required;MaxSize(100)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
}

func AddReservation(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form AddReservationForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	reservationService := reservation_service.Reservation{
		Name:      form.Name,
		Date:      form.Date,
		PeriodID:  form.PeriodID,
		Address:   form.Address,
		DoctorID:  form.DoctorID,
		CreatedBy: form.CreatedBy,
	}

	// 是否一个人可以预约多个
	/*exists, err := reservationService.ExistByName()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_reservation_FAIL, nil)
		return
	}
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_reservation, nil)
		return
	}*/

	err := reservationService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_ADD_RESERVATION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

type EditReservationForm struct {
	ID       int    `form:"id" valid:"Required;Min(1)"`
	Name     string `form:"name" valid:"Required;MaxSize(100)"`
	Date     int    `form:"date" valid:"Required;Min(1)"`
	PeriodID int    `form:"period_id" valid:"Required;Min(1)"`
	DoctorID int    `form:"doctor_id" valid:"Required;Min(1)"`
	Address  string `form:"address" valid:"Required;MaxSize(100)"`
}

func EditReservation(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		form = EditReservationForm{ID: com.StrTo(c.Param("id")).MustInt()}
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	reservationService := reservation_service.Reservation{
		ID:       form.ID,
		Name:     form.Name,
		Date:     form.Date,
		PeriodID: form.PeriodID,
		Address:  form.Address,
		DoctorID: form.DoctorID,
	}

	exists, err := reservationService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_RESERVATION_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESERVATION, nil)
		return
	}

	err = reservationService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_RESERVATION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func DeleteReservation(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}
	id := com.StrTo(c.Query("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	reservationService := reservation_service.Reservation{ID: id}
	exists, err := reservationService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_RESERVATION_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_RESERVATION, nil)
		return
	}

	if err := reservationService.Delete(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_RESERVATION_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
