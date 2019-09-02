package models

import "github.com/jinzhu/gorm"

type Examination struct {
	Model
	PatientID     int `json:"patient_id"`
	Height        int `json:"height"`
	Weight        int `json:"weight"`
	BloodPressure int `json:"blood_pressure"`
}

func GetExaminationList(pageNum, pageSize, patientID int) ([]*Examination, error) {
	var exList []*Examination
	err := db.Where("patient_id = ? AND deleted_on = ?", patientID, 0).Offset(pageNum).Limit(pageSize).Find(&exList).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return exList, nil
}

func AddExaminationInfo(data map[string]interface{}) error {
	ex := Examination{
		PatientID:     data["patient_id"].(int),
		Height:        data["height"].(int),
		Weight:        data["weight"].(int),
		BloodPressure: data["blood_pressure"].(int),
	}
	if err := db.Create(&ex).Error; err != nil {
		return err
	}
	return nil
}

func DelExamination(id int) error {
	if err := db.Where("id = ?", id).Delete(&Examination{}).Error; err != nil {
		return err
	}
	return nil
}
