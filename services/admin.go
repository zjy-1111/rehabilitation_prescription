package services

import (
	"rehabilitation_prescription/models"
)

type Admin struct {
	ID       int
	Username string
	Password string
}

func (a *Admin) Check() (bool, error) {
	return models.CheckAdmin(a.Username, a.Password)
}

func (a *Admin) ExistByName() (bool, error) {
	return models.ExistAdminByName(a.Username)
}

func (a *Admin) ExistByID() (bool, error) {
	return models.ExistAdminByID(a.ID)
}

func (a *Admin) GetAdminByID() (admin models.Admin, err error) {
	admin, err = models.GetAdminByID(a.ID)
	if err != nil {
		return admin, err
	}

	return admin, nil
}

func (a *Admin) GetAdmins() (admins []*models.Admin, err error) {
	admins, err = models.GetAdmins()
	if err != nil {
		return nil, err
	}

	return admins, err
}

func (a *Admin) Count() (int, error) {
	return models.GetAdminsTotal(map[string]interface{}{
		"deleted_on": 0,
	})
}

func (a *Admin) Add() error {
	return models.AddAdmin(a.Username, a.Password)
}

func (a *Admin) Edit() error {
	data := make(map[string]interface{})
	data["username"] = a.Username
	data["password"] = a.Password

	return models.EditAdmin(a.ID, data)
}

func (a *Admin) Delete() error {
	return models.DeleteAdmin(a.ID)
}
