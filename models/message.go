package models

import "github.com/jinzhu/gorm"

type Message struct {
	Model

	MessageTextID int
	Receiver      string
	Read          int
}

func ExistMessageByID(id int) (bool, error) {
	var m Message
	err := db.Select("id").Where("id = ? AND deleted_on = ?", id, 0).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if m.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetMessage(id int) (*Message, error) {
	var m Message
	err := db.Where("deleted_on = ? AND id = ?", 0, id).First(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &m, nil
}

func GetMessages(pageNum, pageSize int) ([]*Message, error) {
	var m []*Message
	err := db.Where("deleted_on = ?", 0).Offset(pageNum).Limit(pageSize).Find(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return m, nil
}

func GetUnReadMessages(pageNum, pageSize int) ([]*Message, error) {
	var m []*Message
	err := db.Where("deleted_on = ? AND read = ?", 0, 0).Offset(pageNum).Limit(pageSize).Find(&m).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return m, nil
}

func AddMessage(data map[string]interface{}) error {
	m := Message{
		MessageTextID: data["message_text_id"].(int),
		Receiver:      data["receiver"].(string),
		Read:          data["read"].(int),
	}

	if err := db.Create(&m).Error; err != nil {
		return err
	}

	return nil
}

func EditMessage(id int, data interface{}) error {
	if err := db.Model(&TrainingVideo{}).Where("id = ? And deleted_on = ?", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func DelMessage(id int) error {
	if err := db.Where("id = ?", id).Delete(TrainingVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanMessage() error {
	if err := db.Unscoped().Where("deleted_on != ?", 0).Delete(&TrainingVideo{}).Error; err != nil {
		return err
	}

	return nil
}

func GetMessagesTotal(maps interface{}) (int, error) {
	count := 0
	if err := db.Model(&Message{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
