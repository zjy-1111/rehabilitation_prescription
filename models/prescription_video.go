package models

import (
	"github.com/jinzhu/gorm"
)

type PrescriptionVideo struct {
	Model
	PrescriptionID int `json:"prescription_id"`
	VideoID        int `json:"video_id"`
}

func ExistPrescriptionVideoByID(id int) (bool, error) {
	var p PrescriptionVideo
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if p.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetPrescriptionVideos(pageNum, pageSize, prescriptionID int) ([]*PrescriptionVideo, error) {
	var p []*PrescriptionVideo
	err := db.Where("prescription_id = ? AND deleted_on = ?", prescriptionID).Offset(pageNum).Limit(pageSize).Find(&p).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return p, nil
}

func AddPrescriptionVideo(data map[string]interface{}) error {
	p := PrescriptionVideo{
		PrescriptionID: data["prescription_id"].(int),
		VideoID:        data["video_id"].(int),
	}

	if err := db.Create(&p).Error; err != nil {
		return err
	}

	return nil
}

func EditPrescriptionVideo(id int, data interface{}) error {
	if err := db.Model(&PrescriptionVideo{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelPrescriptionVideo(id int) error {
	if err := db.Where("id = ?", id).Delete(PrescriptionVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanPrescriptionVideo() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&PrescriptionVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func GetPrescriptionVideosTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&PrescriptionVideo{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
