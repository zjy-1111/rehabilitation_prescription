package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string `json:"username"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
	Name     string `json:"name"`
}

// CheckUser checks if uentication information exists
func CheckUser(username, password, userType string) (bool, error) {
	var u User
	err := db.Select("id").Where(User{Username: username, Password: password, UserType: userType}).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 && u.DeletedOn == 0 {
		return true, nil
	}

	return false, nil
}

func ExistUserByName(name string) (bool, error) {
	var u User
	err := db.Select("id").Where("username = ? AND deleted_on = ?", name, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 {
		return true, nil
	}

	return false, nil
}

func ExistUserByID(id int) (bool, error) {
	var u User
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&u).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if u.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetUserByID(id int) (User, error) {
	var admin User
	err := db.Where("deleted_on = ? AND id = ?", 0, id).First(&admin).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return admin, err
	}

	return admin, nil
}

func GetUsersByType(userType string) ([]*User, error) {
	var users []*User
	err := db.Where("deleted_on = ? AND user_type = ?", 0, userType).Find(&users).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return users, nil
}

func GetUsersTotalByType(userType string) (total int, err error) {
	if err := db.Model(&User{}).Where("deleted_on = ? AND user_type = ?", 0, userType).Count(&total).Error; err != nil {
		return 0, err
	}

	return total, nil
}

func AddUser(username, password, userType, name string) error {
	u := User{
		Username: username,
		Password: password,
		UserType: userType,
		Name:     name,
	}
	if err := db.Create(&u).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	if err := db.Where("id = ?", id).Delete(&User{}).Error; err != nil {
		return err
	}

	return nil
}

func EditUser(id int, data interface{}) error {
	if err := db.Model(&User{}).Where("id = ? AND deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func CleanDeletedUser() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&User{}).Error; err != nil {
		return false, err
	}

	return true, nil
}
