package services

import "rehabilitation_prescription/models"

type Examination struct {
	ID            int
	PatientID     int
	Height        int
	Weight        int
	BloodPressure int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (e *Examination) Get() ([]*models.Examination, error) {
	exList, err := models.GetExaminationList(e.PageNum, e.PageSize, e.PatientID)
	if err != nil {
		return nil, err
	}
	return exList, nil
}

func (e *Examination) Add() error {
	data := map[string]interface{}{
		"patient_id":     e.PatientID,
		"height":         e.Height,
		"weight":         e.Weight,
		"blood_pressure": e.BloodPressure,
	}
	err := models.AddExaminationInfo(data)
	if err != nil {
		return err
	}
	return nil
}

func (e *Examination) Del() error {
	return models.DelExamination(e.ID)
}
