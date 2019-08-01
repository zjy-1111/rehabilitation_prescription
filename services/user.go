package services

import (
	"rehabilitation_prescription/models"
)

type User struct {
	ID       int
	Username string
	Password string
	UserType string
	Name     string
	Avatar   string

	PageNum int
	PageSize int
}

func (u *User) Check() (bool, error) {
	return models.CheckUser(u.Username, u.Password, u.UserType)
}

func (u *User) ExistByName() (bool, error) {
	return models.ExistUserByName(u.Username)
}

func (u *User) ExistByID() (bool, error) {
	return models.ExistUserByID(u.ID)
}

func (u *User) GetUserByID() (admin models.User, err error) {
	admin, err = models.GetUserByID(u.ID)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (u *User) GetUsersByType() (users []*models.User, err error) {
	users, err = models.GetUsersByType(u.UserType, u.PageNum, u.PageSize)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (u *User) GetUsersTotalByType() (int, error) {
	return models.GetUsersTotalByType(u.UserType)
}

func (u *User) Add() error {
	return models.AddUser(u.Username, u.Password, u.UserType, u.Name, u.Avatar)
}

func (u *User) Edit() error {
	data := make(map[string]interface{})
	data["username"] = u.Username
	data["password"] = u.Password
	data["name"] = u.Name
	data["avatar"] = u.Avatar

	return models.EditUser(u.ID, data)
}

func (u *User) Delete() error {
	return models.DeleteUser(u.ID)
}
