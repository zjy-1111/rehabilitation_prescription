package handlers

import "rehabilitation_prescription/models"

// 根据处方id获取运动处方信息（多条）
func getTrainings(prescriptionID int) []*Training {
	prescriptionTrainings, _ := models.GetPrescriptionTrainings(prescriptionID)

	trainings := make([]*Training, 0)
	for i := 0; i < len(prescriptionTrainings); i++ {
		trainingInfo, err := models.GetTrainingVideo(prescriptionTrainings[i].TrainingID)
		if err != nil {
			continue
		}
		training := &Training{
			ID:        trainingInfo.ID,
			Title:     trainingInfo.Title,
			Desc:      trainingInfo.Description,
			VideoUrl:  trainingInfo.VideoUrl,
			CoverUrl:  trainingInfo.CoverUrl,
			Author:    trainingInfo.CreatedBy,
			CreatedOn: trainingInfo.CreatedOn,
		}
		trainings = append(trainings, training)
	}

	return trainings
}
