package models

import (
	"github.com/jinzhu/gorm"
)

type ReportPrescription struct {
	Model

	PatientReportID int    `json:"patient_report_id"`
	PrescriptionID  int    `json:"prescription_id"`
	Remark          string `json:"remark"`
}

func ExistReportPrescriptionByID(id int) (bool, error) {
	var r ReportPrescription
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&r).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if r.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPrescriptionsOfReport(pageNum, pageSize, reportID int) ([]*ReportPrescription, error) {
	var rs []*ReportPrescription
	err := db.Where("patient_report_id = ? AND deleted_on = ?", reportID, 0).Offset(pageNum).Limit(pageSize).Find(&rs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return rs, nil
}

func AddReportPrescription(data map[string]interface{}) error {
	r := ReportPrescription{
		PatientReportID: data["patient_report_id"].(int),
		PrescriptionID:  data["prescription_id"].(int),
		Remark:          data["remark"].(string),
	}

	if err := db.Create(&r).Error; err != nil {
		return err
	}

	return nil
}

func EditReportPrescription(id int, data interface{}) error {
	if err := db.Model(&ReportPrescription{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelReportPrescription(id int) error {
	if err := db.Where("id = ?", id).Delete(ReportPrescription{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanReportPrescription() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&ReportPrescription{}).Error; err != nil {
		return err
	}

	return nil
}

func GetReportPrescriptionsTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&ReportPrescription{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
