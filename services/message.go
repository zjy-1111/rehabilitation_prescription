package services

import (
	"rehabilitation_prescription/models"
)

type Message struct {
	ID         int
	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	MessageTextID int
	Receiver      string
	Read          int

	PageNum  int
	PageSize int
}

func (m *Message) Add() error {
	msg := map[string]interface{}{
		"message_text_id": m.MessageTextID,
		"receiver":        m.Receiver,
		"read":            m.Read,
	}

	if err := models.AddMessage(msg); err != nil {
		return err
	}

	return nil
}

func (m *Message) Edit() error {
	return models.EditMessage(m.ID, map[string]interface{}{
		"message_text_id": m.MessageTextID,
		"receiver":        m.Receiver,
		"read":            m.Read,
	})
}

func (m *Message) Get() ([]*models.Message, error) {
	msgs, err := models.GetMessages(m.PageNum, m.PageSize)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (m *Message) Del() error {
	return models.DelMessage(m.ID)
}

func (m *Message) Count() (int, error) {
	return models.GetMessagesTotal(map[string]interface{}{
		"receiver":   m.Receiver,
		"deleted_on": 0,
	})
}

func (m *Message) UnReadCount() (int, error) {
	return models.GetMessagesTotal(map[string]interface{}{
		"receiver":   m.Receiver,
		"deleted_on": 0,
		"read":       0,
	})
}

func (m *Message) ExistByID() (bool, error) {
	return models.ExistMessageByID(m.ID)
}
