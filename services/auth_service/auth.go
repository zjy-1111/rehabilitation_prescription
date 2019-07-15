package auth_service

import (
	"rehabilitation_prescription/models"
)

type Auth struct {
	ID       int
	Username string
	Password string
	UserType int
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password, a.UserType)
}

func (a *Auth) ExistByName() (bool, error) {
	return models.ExistAuthByName(a.Username)
}

func (a *Auth) ExistByID() (bool, error) {
	return models.ExistAuthByID(a.ID)
}

func (a *Auth) Add() error {
	return models.AddAuth(a.Username, a.Password, a.UserType)
}

func (a *Auth) Edit() error {
	data := make(map[string]interface{})
	data["username"] = a.Username
	data["password"] = a.Password

	return models.EditAuth(a.Username, data)
}

func (a *Auth) Delete() error {
	return models.DeleteAuth(a.Username)
}
