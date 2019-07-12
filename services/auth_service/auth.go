package auth_service

import (
	"rehabilitation_prescription/models"
)

type Auth struct {
	ID       int
	Username string
	Password string
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}

func (a *Auth) ExistByName() (bool, error) {
	return models.ExistAuthByName(a.Username)
}

func (a *Auth) Add() error {
	return models.AddAuth(a.Username, a.Password)
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
