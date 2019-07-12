package models

import (
	"github.com/jinzhu/gorm"
)

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAuth checks if authentication information exists
func CheckAuth(username, password string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistAuthByName(name string) (bool, error) {
	var auth Auth
	err := db.Select("id").Where("username = ?", name).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func AddAuth(username, password string) error {
	auth := Auth{
		Username: username,
		Password: password,
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
	if err := db.Model(&Auth{}).Where("username = ?", username).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
