package api

import (
	"fmt"
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

type AdminForm struct {
	Username string `form:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
}

// @Summary 获取token
// @Produce json
// @param username query string true "Username"
// @param password query string true "Password"
// @Success 200 {object} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin [get]
func Login(c *gin.Context) {
	var form = AdminForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	adminService := services.Admin{
		Username: form.Username,
		Password: form.Password,
	}

	isExist, err := adminService.Check()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		app.Response(c, http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

func GetAdminByID(c *gin.Context) {
	service := services.Admin{ID: com.StrTo(c.Param("id")).MustInt()}

	admin, err := service.GetAdminByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ADMIN_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["admin"] = admin
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

func GetAdmins(c *gin.Context) {
	service := services.Admin{}
	total, err := service.Count()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_ADMINS_FAIL, nil)
		return
	}

	admins, err := service.GetAdmins()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ADMINS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = admins
	data["total"] = total
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

// @Summary Add admin
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin [post]
func AddAdmin(c *gin.Context) {
	var form = AdminForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	adminService := services.Admin{
		Username: form.Username,
		Password: form.Password,
	}
	exists, err := adminService.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}
	if exists {
		app.Response(c, http.StatusOK, e.ERROR_EXIST_AUTH, nil)
		return
	}

	err = adminService.Add()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_ADD_AUTH_FAIL, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary Update Admin
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/{username} [put]
func EditAdmin(c *gin.Context) {
	form := AdminForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	adminService := services.Admin{
		ID:       com.StrTo(c.Param("id")).MustInt(),
		Username: form.Username,
		Password: form.Password,
	}

	exists, err := adminService.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}

	if exists {
		app.Response(c, http.StatusOK, e.ERROR_EXIST_AUTH, nil)
		return
	}

	err = adminService.Edit()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EDIT_AUTH_FAIL, nil)
		return
	}

	token, err := util.GenerateToken(form.Username, form.Password)
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}

// @Summary Delete admin
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /admin/{username} [delete]
func DeleteAdmin(c *gin.Context) {
	adminService := services.Admin{ID: com.StrTo(c.Param("id")).MustInt()}
	fmt.Println(adminService.ID)
	exists, err := adminService.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}

	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_AUTH, nil)
		return
	}

	if err := adminService.Delete(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_AUTH_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
