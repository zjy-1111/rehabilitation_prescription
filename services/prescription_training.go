package services

import "rehabilitation_prescription/models"

type RPrescriptionTraining struct {
	ID             int
	PrescriptionID int
	TrainingID     int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (p *RPrescriptionTraining) Add() error {
	training := map[string]interface{}{
		"prescription_id": p.PrescriptionID,
		"training_id":     p.TrainingID,
	}

	if err := models.AddRPrescriptionTraining(training); err != nil {
		return err
	}

	return nil
}

func (p *RPrescriptionTraining) Edit() error {
	return models.EditRPrescriptionTraining(p.ID, map[string]interface{}{
		"prescription_id": p.PrescriptionID,
		"training_id":     p.TrainingID,
	})
}

//func (p *RPrescriptionTraining) Get() ([]*models.RPrescriptionTraining, error) {
//	training_ids, err := models.GetRPrescriptionTrainings(p.PageNum, p.PageSize, p.PrescriptionID)
//	if err != nil {
//		return nil, err
//	}
//
//	return training_ids, nil
//}

func (p *RPrescriptionTraining) Del() error {
	return models.DelRPrescriptionTraining(p.ID)
}

func (p *RPrescriptionTraining) Count() (int, error) {
	return models.GetRPrescriptionTrainingsTotal(map[string]interface{}{
		"deleted_on": 0,
	})
}

func (p *RPrescriptionTraining) ExistByID() (bool, error) {
	return models.ExistRPrescriptionTrainingByID(p.ID)
}
