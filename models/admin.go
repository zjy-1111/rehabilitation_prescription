package models

import (
	"github.com/jinzhu/gorm"
)

type Admin struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// CheckAdmin checks if authentication information exists
func CheckAdmin(username, password string) (bool, error) {
	var auth Admin
	err := db.Select("id").Where(Admin{Username: username, Password: password}).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 && auth.DeletedOn == 0 {
		return true, nil
	}

	return false, nil
}

func ExistAdminByName(name string) (bool, error) {
	var auth Admin
	err := db.Select("id").Where("username = ? AND deleted_on = ?", name, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistAdminByID(id int) (bool, error) {
	var auth Admin
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&auth).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if auth.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetAdminByID(id int) (Admin, error) {
	var admin Admin
	err := db.Where("deleted_on = ? AND id = ?", 0, id).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin, err
	}

	return admin, nil
}

func GetAdmins() ([]*Admin, error) {
	var admins []*Admin
	err := db.Where("deleted_on = ?", 0).Find(&admins).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return admins, nil
}

func GetAdminsTotal(maps interface{}) (total int, err error) {
	if err := db.Model(&TrainingVideo{}).Where(maps).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func AddAdmin(username, password string) error {
	auth := Admin{
		Username: username,
		Password: password,
	}
	if err := db.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}

func DeleteAdmin(id int) error {
	if err := db.Where("id = ?", id).Delete(&Admin{}).Error; err != nil {
		return err
	}

	return nil
}

func EditAdmin(id int, data interface{}) error {
	if err := db.Model(&Admin{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func CleanDeletedAdmin() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Admin{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
