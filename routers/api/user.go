package api

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Username string `form:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
	UserType string `form:"user_type" valid:"MaxSize(10)"`
	Name     string `form:"name" valid:"MaxSize(100)"`
	Avatar   string `form:"avatar" valid:"MaxSize(255)"`
}

// @Summary 获取token
// @Produce json
// @param username query string true "Username"
// @param password query string true "Password"
// @Success 200 {object} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /admin [get]
func Login(c *gin.Context) {
	var form = UserForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	s := services.User{
		Username: form.Username,
		Password: form.Password,
		UserType: "3",
	}

	isExist, err := s.Check()
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

func GetUserByID(c *gin.Context) {
	s := services.User{ID: com.StrTo(c.Param("id")).MustInt()}

	user, err := s.GetUserByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ADMIN_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["user"] = user
	app.Response(c, http.StatusOK, e.SUCCESS, data)
}

func GetUsers(c *gin.Context) {
	s := services.User{
		UserType: c.Param("type"),
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	total, err := s.GetUsersTotalByType()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_COUNT_ADMINS_FAIL, nil)
		return
	}

	users, err := s.GetUsersByType()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_GET_ADMINS_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["items"] = users
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
func AddUser(c *gin.Context) {
	var form = UserForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	s := services.User{
		Username: form.Username,
		Password: form.Password,
		UserType: form.UserType,
		Name:     form.Name,
		Avatar:   form.Avatar,
	}

	exist, err := s.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}
	if exist {
		app.Response(c, http.StatusOK, e.ERROR_EXIST_AUTH, nil)
		return
	}

	err = s.Add()
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
func EditUser(c *gin.Context) {
	form := UserForm{}

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		app.Response(c, httpCode, errCode, nil)
		return
	}

	s := services.User{
		ID:       com.StrTo(c.Param("id")).MustInt(),
		Username: form.Username,
		Password: form.Password,
		Name:     form.Name,
		Avatar:   form.Avatar,
	}

	err := s.Edit()
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
func DeleteUser(c *gin.Context) {
	s := services.User{ID: com.StrTo(c.Param("id")).MustInt()}
	exist, err := s.ExistByID()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}

	if !exist {
		app.Response(c, http.StatusOK, e.ERROR_NOT_EXIST_AUTH, nil)
		return
	}

	if err := s.Delete(); err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_DELETE_AUTH_FAIL, nil)
		return
	}

	app.Response(c, http.StatusOK, e.SUCCESS, nil)
}
