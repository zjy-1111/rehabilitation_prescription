package services

import (
	"rehabilitation_prescription/models"
)

type Appointment struct {
	ID        int
	PatientID int
	DoctorID  int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (a *Appointment) Add() error {
	appointment := map[string]interface{}{
		"patient_id": a.PatientID,
		"doctor_id":  a.DoctorID,
	}

	if err := models.AddAppointment(appointment); err != nil {
		return err
	}

	return nil
}

func (a *Appointment) Edit() error {
	return models.EditAppointment(a.ID, map[string]interface{}{
		"patient_id": a.PatientID,
		"doctor_id":  a.DoctorID,
	})
}

//func (a *Appointment) Get() ([]int, error) {
//	patientIDs, err := models.GetPatientsByDoctorID(a.PageNum, a.PageSize, a.DoctorID, )
//	if err != nil {
//		return nil, err
//	}
//
//	return patientIDs, nil
//}

func (a *Appointment) Del() error {
	return models.DelAppointment(a.ID)
}

func (a *Appointment) Count() (int, error) {
	return models.GetPatientsTotal(map[string]interface{}{
		"deleted_on": 0,
		"doctor_id":  a.DoctorID,
	})
}

func (a *Appointment) ExistByID() (bool, error) {
	return models.ExistAppointmentByID(a.ID)
}
