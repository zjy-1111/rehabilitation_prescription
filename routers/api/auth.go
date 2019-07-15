package api

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/services/auth_service"
	"rehabilitation_prescription/util"

	"github.com/gin-gonic/gin"
)

type AuthForm struct {
	Username string `form:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
	UserType int    `form:"user_type" valid:"Required;Range(1,2)"`
}

// @Summary 获取token
// @Produce json
// @param username query string true "Username"
// @param password query string true "Password"
// @Success 200 {object} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	var form = AuthForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{
		Username: form.Username,
		Password: form.Password,
		UserType: form.UserType,
	}

	isExist, err := authService.Check()
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

// @Summary Add auth
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func AddAuth(c *gin.Context) {
	var form = AuthForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{
		Username: form.Username,
		Password: form.Password,
		UserType: form.UserType,
	}
	exists, err := authService.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}
	if exists {
		app.Response(c, http.StatusOK, e.ERROR_EXIST_AUTH, nil)
		return
	}

	err = authService.Add()
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

// @Summary Update Auth
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/{username} [put]
func EditAuth(c *gin.Context) {
	form := AuthForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{
		Username: form.Username,
		Password: form.Password,
		UserType: form.UserType,
	}

	exists, err := authService.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}

	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_AUTH, nil)
		return
	}

	err = authService.Edit()
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

// @Summary Delete auth
// @Produce  json
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth/{username} [delete]
func DeleteAuth(c *gin.Context) {
	authService := auth_service.Auth{Username: c.Param("username")}
	exists, err := authService.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}

	if !exists {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_AUTH, nil)
		return
	}

	if err := authService.Delete(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_AUTH_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
