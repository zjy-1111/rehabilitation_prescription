package api

import (
	"net/http"
	"rehabilitation_prescription/pkg/app"
	"rehabilitation_prescription/pkg/e"
	"rehabilitation_prescription/pkg/setting"
	"rehabilitation_prescription/services"
	"rehabilitation_prescription/util"
	"strconv"

	"github.com/Unknwon/com"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Username string `form:"username" valid:"Required;MaxSize(100)"`
	Password string `form:"password" valid:"Required;MaxSize(100)"`
	UserType string `form:"user_type" valid:"MaxSize(10)"`
	Avatar   string `form:"avatar" valid:"MaxSize(255)"`
}

func AdminLogin(c *gin.Context) {
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

	uid, isExist, err := s.Check()
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
		"uid":   strconv.Itoa(uid),
		"token": token,
	})
}

// @Summary 普通用户登录
// @Description get a token string
// @Accept json
// @Produce json
// @Param username query string true "用户名"
// @Param password query string true "密码"
// @Param user_type query string true "用户类型"
// @Success 200 {object} models.Response
// @Router /user/login [post]
func UserLogin(c *gin.Context) {
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
	}

	uid, isExist, err := s.Check()
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
		"uid":   strconv.Itoa(uid),
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

// @Summary 注册用户
// @accept json
// @produce json
// @param username query string true "用户名"
// @param password query string true "密码"
// @param user_type query string true "用户类型"
// @param name query string false "姓名"
// @param avatar query string false "头像"
// @success 200 {object} models.Response
// @router /user [post]
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
		Avatar:   form.Avatar,
	}

	exist, err := s.ExistByName()
	if err != nil {
		app.Response(c, http.StatusInternalServerError, e.ERROR_EXIST_AUTH_FAIL, nil)
		return
	}
	if exist {
		app.Response(c, http.StatusCreated, e.ERROR_EXIST_AUTH, nil)
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
