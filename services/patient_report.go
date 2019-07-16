package services

import "rehabilitation_prescription/models"

type PatientReport struct {
	ID            int
	PatientID     int
	BodyType      string
	Height        string
	Weight        string
	Waist         string
	BloodPressure string

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (p *PatientReport) Add() error {
	report := map[string]interface{}{
		"patient_id":     p.PatientID,
		"body_type":      p.BodyType,
		"height":         p.Height,
		"weight":         p.Weight,
		"waist":          p.Waist,
		"blood_pressure": p.BloodPressure,
	}

	if err := models.AddPatientReport(report); err != nil {
		return err
	}

	return nil
}

func (p *PatientReport) Edit() error {
	return models.EditPatientReport(p.ID, map[string]interface{}{
		"patient_id":     p.PatientID,
		"body_type":      p.BodyType,
		"height":         p.Height,
		"weight":         p.Weight,
		"waist":          p.Waist,
		"blood_pressure": p.BloodPressure,
	})
}

func (p *PatientReport) Get() ([]*models.PatientReport, error) {
	prescriptions, err := models.GetPatientReportByPatienID(p.PageNum, p.PageSize, p.PatientID)
	if err != nil {
		return nil, err
	}

	return prescriptions, nil
}

func (p *PatientReport) Del() error {
	return models.DelPatientReport(p.ID)
}

func (p *PatientReport) Count() (int, error) {
	return models.GetPatientReportsTotal(map[string]interface{}{
		"deleted_on": 0,
		"doctor_id":  p.PatientID,
	})
}

func (p *PatientReport) ExistByID() (bool, error) {
	return models.ExistPatientReportByID(p.ID)
}
