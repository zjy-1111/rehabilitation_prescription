package models

import (
	"github.com/jinzhu/gorm"
)

type RPrescriptionTraining struct {
	Model
	PrescriptionID int `json:"prescription_id"`
	TrainingID     int `json:"training_id"`
}

func ExistRPrescriptionTrainingByID(id int) (bool, error) {
	var p RPrescriptionTraining
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if p.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPrescriptionTrainings(prescriptionID int) ([]*RPrescriptionTraining, error) {
	var p []*RPrescriptionTraining
	err := db.Where("prescription_id = ? AND deleted_on = ?", prescriptionID, 0).Find(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return p, nil
}

func AddRPrescriptionTraining(data map[string]interface{}) error {
	p := RPrescriptionTraining{
		PrescriptionID: data["prescription_id"].(int),
		TrainingID:     data["training_id"].(int),
	}

	if err := db.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func EditRPrescriptionTraining(id int, data interface{}) error {
	if err := db.Model(&RPrescriptionTraining{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelRPrescriptionTraining(id int) error {
	if err := db.Where("id = ?", id).Delete(RPrescriptionTraining{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanRPrescriptionTraining() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&RPrescriptionTraining{}).Error; err != nil {
		return err
	}

	return nil
}

func GetRPrescriptionTrainingsTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&RPrescriptionTraining{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
