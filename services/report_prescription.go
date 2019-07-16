package services

import "rehabilitation_Prescription/models"

type ReportPrescription struct {
	ID              int
	PatientReportID int
	PrescriptionID  int
	Remark          string

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (r *ReportPrescription) Add() error {
	ReportPrescription := map[string]interface{}{
		"patient_report_id": r.PatientReportID,
		"prescription_id":   r.PrescriptionID,
		"remark":            r.Remark,
	}

	if err := models.AddReportPrescription(ReportPrescription); err != nil {
		return err
	}

	return nil
}

func (r *ReportPrescription) Edit() error {
	return models.EditReportPrescription(r.ID, map[string]interface{}{
		"patient_report_id": r.PatientReportID,
		"prescription_id":   r.PrescriptionID,
		"remark":            r.Remark,
	})
}

func (r *ReportPrescription) Get() ([]*models.ReportPrescription, error) {
	ReportPrescriptions, err := models.GetPrescriptionsOfReport(r.PageNum, r.PageSize, r.PatientReportID)
	if err != nil {
		return nil, err
	}

	return ReportPrescriptions, nil
}

func (r *ReportPrescription) Del() error {
	return models.DelAppointment(r.ID)
}

func (r *ReportPrescription) Count() (int, error) {
	return models.GetReportPrescriptionsTotal(map[string]interface{}{
		"deleted_on":        0,
		"patient_report_id": r.PatientReportID,
	})
}

func (r *ReportPrescription) ExistByID() (bool, error) {
	return models.ExistReportPrescriptionByID(r.ID)
}
