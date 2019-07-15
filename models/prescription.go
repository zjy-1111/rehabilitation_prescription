package models

import (
	"github.com/jinzhu/gorm"
)

type Prescription struct {
	Model

	Desc     string `json:"desc"`
	DoctorID int    `json:"doctor_id"`
}

func ExistPrescriptionByID(id int) (bool, error) {
	var p Prescription
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if p.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPrescriptionByDoctorID(pageNum, pageSize, doctorID int) ([]*Prescription, error) {
	var ps []*Prescription
	err := db.Where("doctor_id = ? AND deleted_on = ?", doctorID, 0).Offset(pageNum).Limit(pageSize).Find(&ps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ps, nil
}

func AddPrescription(data map[string]interface{}) error {
	p := Prescription{
		Desc:     data["prescription_describe"].(string),
		DoctorID: data["doctor_id"].(int),
	}

	if err := db.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func EditPrescription(id int, data interface{}) error {
	if err := db.Model(&Prescription{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelPrescription(id int) error {
	if err := db.Where("id = ?", id).Delete(Prescription{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanPrescription() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&Prescription{}).Error; err != nil {
		return err
	}

	return nil
}

func GetPrescriptionsTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&Prescription{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
