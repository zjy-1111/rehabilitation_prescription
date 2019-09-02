package services

import (
	"rehabilitation_prescription/models"
)

type Prescription struct {
	ID             int
	Title          string
	Desc           string
	PatientID      int
	TrainingIDList []int
	//Name      string
	//Sex       string
	//Age       int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (p *Prescription) Add() error {
	prescription := map[string]interface{}{
		"title":      p.Title,
		"patient_id": p.PatientID,
		"desc":       p.Desc,
		//"name":       p.Name,
		//"sex":        p.Sex,
		//"age":        p.Age,
	}

	if err := models.AddPrescription(prescription, p.TrainingIDList); err != nil {
		return err
	}

	return nil
}

func (p *Prescription) Edit() error {
	return models.EditPrescription(p.ID, map[string]interface{}{
		"title":      p.Title,
		"patient_id": p.PatientID,
		//"name":       p.Name,
		//"sex":        p.Sex,
		//"age":        p.Age,
	})
}

func (p *Prescription) Get() ([]*models.Prescription, error) {
	prescriptions, err := models.GetPatientPrescriptions(p.PageNum, p.PageSize, p.PatientID)
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (p *Prescription) Del() error {
	return models.DelPrescription(p.ID)
}

func (p *Prescription) Count() (int, error) {
	return models.GetPrescriptionsTotal(map[string]interface{}{
		"deleted_on": 0,
		"patient_id": p.PatientID,
	})
}

func (p *Prescription) ExistByID() (bool, error) {
	return models.ExistPrescriptionByID(p.ID)
}
