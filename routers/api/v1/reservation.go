package v1

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services/reservation_service"
	"rehabilitation_prescription/util"

	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/astaxie/beego/validation"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

func GetReservations(c *gin.Context) {
	name := c.Query("name")

	reservation_service := reservation_service.Reservation{
		Name:     name,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}

	appG := app.Gin{c}
	reservations, err := reservation_service.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_RESERVATIONS_FAIL, nil)
		return
	}

	count, err := reservation_service.Count()
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
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	Time       int    `form:"time" valid:"Required;MaxSize(10)"`
	DoctorName string `form:"doctor_name" valid:"Required;MaxSize(100)"`
	Address    string `form:"address" valid:"Required;MaxSize(100)"`
	CreatedBy  string `form:"created_by" valid:"Required;MaxSize(100)"`
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
		Name:       form.Name,
		Time:       form.Time,
		Address:    form.Address,
		DoctorName: form.DoctorName,
		CreatedBy:  form.CreatedBy,
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
	ID         int    `form:"id" valid:"Required;Min(1)"`
	Name       string `form:"name" valid:"Required;MaxSize(100)"`
	Times      timestamp.Timestamp
	Time       int    `form:"time" valid:"Required;MaxSize(10)"`
	DoctorName string `form:"doctor_name" valid:"Required;MaxSize(100)"`
	Address    string `form:"address" valid:"Required;MaxSize(100)"`
	ModifiedBy string `form:"modified_by" valid:"Required;MaxSize(100)"`
}

func EditTag(c *gin.Context) {
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
		ID:         form.ID,
		Name:       form.Name,
		Time:       form.Time,
		Address:    form.Address,
		DoctorName: form.DoctorName,
		ModifiedBy: form.ModifiedBy,
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
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
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
