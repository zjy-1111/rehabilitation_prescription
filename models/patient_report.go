package models

import (
	"github.com/jinzhu/gorm"
)

type PatientReport struct {
	Model

	ReportNum     string `json:"report_num"`
	PatientID     int    `json:"patient_id"`
	BodyType      string `json:"body_type"`
	Height        string `json:"height"`
	Weight        string `json:"weight"`
	Waist         string `json:"waist"`
	BloodPressure string `json:"blood_pressure"`
}

func ExistPatientReportByID(id int) (bool, error) {
	var p PatientReport
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if p.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPatientReportByPatienID(pageNum, pageSize, patientID int) ([]*PatientReport, error) {
	var ps []*PatientReport
	err := db.Where("patient_id = ? AND deleted_on = ?", patientID, 0).Offset(pageNum).Limit(pageSize).Find(&ps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ps, nil
}

func AddPatientReport(data map[string]interface{}) error {
	p := PatientReport{
		PatientID:     data["patient_id"].(int),
		BodyType:      data["body_type"].(string),
		Height:        data["height"].(string),
		Weight:        data["weight"].(string),
		Waist:         data["waist"].(string),
		BloodPressure: data["blood_pressure"].(string),
	}

	if err := db.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func EditPatientReport(id int, data interface{}) error {
	if err := db.Model(&Prescription{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelPatientReport(id int) error {
	if err := db.Where("id = ?", id).Delete(PatientReport{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanPatientReport() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&PatientReport{}).Error; err != nil {
		return err
	}

	return nil
}

func GetPatientReportsTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&PatientReport{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
