package models

import (
	"github.com/jinzhu/gorm"
)

type TrainingVideo struct {
	Model

	Title       string `json:"title"`
	Description string `json:"desc"`
	VideoUrl    string `json:"video_url"`
	CoverUrl    string `json:"cover_url"`
	Duration    int    `json:"duration"`
	CreatedBy   string `json:"created_by"`
}

func ExistTrainingVideoByID(id int) (bool, error) {
	var t TrainingVideo
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if t.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetTrainingVideo(id int) (*TrainingVideo, error) {
	var t TrainingVideo
	err := db.Where("deleted_on = ? AND id = ?", 0, id).First(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &t, nil
}

func GetTrainingVideos() ([]*TrainingVideo, error) {
	var t []*TrainingVideo
	err := db.Where("deleted_on = ?", 0).Find(&t).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return t, nil
}

func AddTrainingVideo(data map[string]interface{}) error {
	t := TrainingVideo{
		Title:       data["title"].(string),
		Description: data["description"].(string),
		VideoUrl:    data["video_url"].(string),
		CoverUrl:    data["cover_url"].(string),
		Duration:    data["duration"].(int),
	}

	if err := db.Create(&t).Error; err != nil {
		return err
	}

	return nil
}

func EditTrainingVideo(id int, data interface{}) error {
	if err := db.Model(&TrainingVideo{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelTrainingVideo(id int) error {
	if err := db.Where("id = ?", id).Delete(TrainingVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanTrainingVideo() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&TrainingVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func GetTrainingVideosTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&TrainingVideo{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
