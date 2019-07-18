package services

import "rehabilitation_prescription/models"

type TrainingVideo struct {
	ID       int
	VideoUrl string
	CoverUrl string
	Duration int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (t *TrainingVideo) Add() error {
	video := map[string]interface{}{
		"video_url": t.VideoUrl,
		"cover_url": t.CoverUrl,
		"duration":  t.Duration,
	}

	if err := models.AddTrainingVideo(video); err != nil {
		return err
	}

	return nil
}

func (t *TrainingVideo) Edit() error {
	return models.EditTrainingVideo(t.ID, map[string]interface{}{
		"video_url": t.VideoUrl,
		"cover_url": t.CoverUrl,
		"duration":  t.Duration,
	})
}

func (t *TrainingVideo) Get() ([]*models.TrainingVideo, error) {
	videos, err := models.GetTrainingVideos(t.PageNum, t.PageSize)
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func (t *TrainingVideo) Del() error {
	return models.DelTrainingVideo(t.ID)
}

func (t *TrainingVideo) Count() (int, error) {
	return models.GetTrainingVideosTotal(map[string]interface{}{
		"deleted_on": 0,
	})
}

func (t *TrainingVideo) ExistByID() (bool, error) {
	return models.ExistTrainingVideoByID(t.ID)
}