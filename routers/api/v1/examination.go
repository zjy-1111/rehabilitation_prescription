package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"
)

type ExaminationListParams struct {
	PatientID int `form:"patient_id" valid:"Required;Min(1)"`
}

func GetExaminationList(c *gin.Context) {
	var params ExaminationListParams
	httpCode, errCode := app.BindAndValid(c, &params)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	s := &services.Examination{PageNum: util.GetPage(c), PageSize: setting.AppSetting.PageSize, PatientID: params.PatientID}
	exList, err := s.Get()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_EXAMINATIONS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = exList
	data["total"] = len(exList)
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

type ExaminationForm struct {
	PatientID     int `form:"patient_id" valid:"Required;Min(1)"`
	Height        int `form:"height" valid:"Required;Min(50);Max(250)"`
	Weight        int `form:"weight" valid:"Required;Min(1);"`
	BloodPressure int `form:"blood_pressure" valid:"Required;Min(1)"`
}

func AddExamination(c *gin.Context) {
	var form ExaminationForm
	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	s := &services.Examination{
		PatientID:     form.PatientID,
		Height:        form.Height,
		Weight:        form.Weight,
		BloodPressure: form.BloodPressure,
	}
	if err := s.Add(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_EXAMINATION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}

func DelExamination(c *gin.Context) {
	valid := validation.Validation{}
	id := com.StrTo(c.Param("id")).MustInt()
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		app.Response(c, http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}

	s := &services.Examination{ID: id}
	if err := s.Del(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_EXAMINATION_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
