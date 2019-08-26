package services

import "rehabilitation_prescription/models"

type TrainingVideo struct {
	ID          int
	Title       string
	Description string
	VideoUrl    string
	CoverUrl    string
	Duration    int

	CreatedOn  int
	ModifiedOn int
	DeletedOn  int

	PageNum  int
	PageSize int
}

func (t *TrainingVideo) Add() error {
	video := map[string]interface{}{
		"title":       t.Title,
		"description": t.Description,
		"video_url":   t.VideoUrl,
		"cover_url":   t.CoverUrl,
		"duration":    t.Duration,
	}

	if err := models.AddTrainingVideo(video); err != nil {
		return err
	}

	return nil
}

func (t *TrainingVideo) Edit() error {
	return models.EditTrainingVideo(t.ID, map[string]interface{}{
		"title": t.Title,
		"description": t.Description,
		"video_url": t.VideoUrl,
		"cover_url": t.CoverUrl,
		"duration":  t.Duration,
	})
}

func (t *TrainingVideo) GetVideoByID() (*models.TrainingVideo, error) {
	video, err := models.GetTrainingVideo(t.ID)
	if err != nil {
		return nil, err
	}

	return video, nil
}

func (t *TrainingVideo) Get() ([]*models.TrainingVideo, error) {
	videos, err := models.GetTrainingVideos(1)
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
