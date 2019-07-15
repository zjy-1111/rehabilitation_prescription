package prescription_service

import "rehabilitation_prescription/models"

type Prescription struct {
	ID       int
	Desc     string
	DoctorID int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (p *Prescription) Add() error {
	prescription := map[string]interface{}{
		"prescription_describe": p.Desc,
		"doctor_id":             p.DoctorID,
	}

	if err := models.AddPrescription(prescription); err != nil {
		return err
	}

	return nil
}

func (p *Prescription) Edit() error {
	return models.EditPrescription(p.ID, map[string]interface{}{
		"prescription_describe": p.Desc,
		"doctor_id":             p.DoctorID,
	})
}

func (p *Prescription) Get() ([]*models.Prescription, error) {
	prescriptions, err := models.GetPrescriptionByDoctorID(p.PageNum, p.PageSize, p.DoctorID)
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (p *Prescription) Del() error {
	return models.DelAppointment(p.ID)
}

func (p *Prescription) Count() (int, error) {
	return models.GetPatientsTotal(map[string]interface{}{
		"deleted_on": 0,
		"doctor_id":  p.DoctorID,
	})
}

func (p *Prescription) ExistByID() (bool, error) {
	return models.ExistPrescriptionByID(p.ID)
}
