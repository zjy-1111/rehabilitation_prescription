package models

import (
	"github.com/jinzhu/gorm"
)

type RDoctorPatient struct {
	Model

	PatientID int `json:"patient_id"`
	DoctorID  int `json:"doctor_id"`
}

func ExistAppointmentByID(id int) (bool, error) {
	var a RDoctorPatient
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&a).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if a.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPatientsByDoctorID(pageNum, pageSize, doctorID int) ([][2]int, error) {
	var as []*RDoctorPatient
	err := db.Select("id,patient_id").Where(
		"doctor_id = ? AND deleted_on = ?", doctorID, 0).Offset(pageNum).Limit(pageSize).Find(&as).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	res := make([][2]int, len(as))
	for i, a := range as {
		res[i][0] = a.ID
		res[i][1] = a.PatientID
	}

	return res, nil
}

func GetAllPatients(doctorID int) ([][2]int, error) {
	var ps []*RDoctorPatient
	err := db.Select("id,patient_id").Where(
		"doctor_id = ? AND deleted_on = ?", doctorID, 0).Find(&ps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	res := make([][2]int, len(ps))
	for i, p := range ps {
		res[i][0] = p.ID
		res[i][1] = p.PatientID
	}

	return res, nil
}


func AddAppointment(data map[string]interface{}) error {
	a := RDoctorPatient{
		PatientID: data["patient_id"].(int),
		DoctorID:  data["doctor_id"].(int),
	}

	if err := db.Create(&a).Error; err != nil {
		return err
	}

	return nil
}

func EditAppointment(id int, data interface{}) error {
	if err := db.Model(&RDoctorPatient{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelAppointment(id int) error {
	if err := db.Where("id = ?", id).Delete(RDoctorPatient{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAppointment() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&RDoctorPatient{}).Error; err != nil {
		return err
	}

	return nil
}

func GetPatientsTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&RDoctorPatient{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
