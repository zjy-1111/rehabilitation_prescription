package models

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"` // 用户类型（1普通用户，2医生, 3管理员）
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string, userType string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password, UserType: userType}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 && auth.DeletedOn == 0 {
		return true, nil
	}

	return false, nil
}

func ExistAuthByName(name string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("username = ? AND deleted_on = ?", name, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistAuthByID(id int) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddAuth(username, password string, userType string) error {
	auth := Auth{
		Username: username,
		Password: password,
		UserType: userType,
	}
	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAuth(username string) error {
	if err := db.Where("username = ?", username).Delete(&Auth{}).Error; err != nil {
		return err
	}

	return nil
}

func EditAuth(username string, data interface{}) error {
	if err := db.Model(&Auth{}).Where("username = ? AND deleted_on = ?", username, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func CleanDeletedAuth() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Auth{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
