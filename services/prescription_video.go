package services

import "rehabilitation_prescription/models"

type PrescriptionVideo struct {
	ID             int
	PrescriptionID int
	VideoID        int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (p *PrescriptionVideo) Add() error {
	video := map[string]interface{}{
		"prescription_id": p.PrescriptionID,
		"video_id":        p.VideoID,
	}

	if err := models.AddPrescriptionVideo(video); err != nil {
		return err
	}

	return nil
}

func (p *PrescriptionVideo) Edit() error {
	return models.EditPrescriptionVideo(p.ID, map[string]interface{}{
		"prescription_id": p.PrescriptionID,
		"video_id":        p.VideoID,
	})
}

func (p *PrescriptionVideo) Get() ([]*models.PrescriptionVideo, error) {
	videos, err := models.GetPrescriptionVideos(p.PageNum, p.PageSize, p.PrescriptionID)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (p *PrescriptionVideo) Del() error {
	return models.DelPrescriptionVideo(p.ID)
}

func (p *PrescriptionVideo) Count() (int, error) {
	return models.GetPrescriptionVideosTotal(map[string]interface{}{
		"deleted_on": 0,
	})
}

func (p *PrescriptionVideo) ExistByID() (bool, error) {
	return models.ExistPrescriptionVideoByID(p.ID)
}
