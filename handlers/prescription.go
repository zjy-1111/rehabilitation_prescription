package handlers

import (
	"rehabilitation_prescription/models"
)

type PrescriptionHandler struct {
	PatientID int
	Offset    int
	Limit     int
}

type Prescription struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Desc        string      `json:"desc"`
	CreatedOn   int         `json:"created_on"`
	PatientName string      `json:"patient_name"`
	Sex         string      `json:"sex"`
	Age         int         `json:"age"`
	Trainings   []*Training `json:"trainings"`
}

type Training struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Desc      string `json:"desc"`
	VideoUrl  string `json:"video_url"`
	CoverUrl  string `json:"cover_url"`
	Author    string `json:"author"`
	CreatedOn int    `json:"created_on"`
}

func NewPrescriptionHandler(patientID, offset, limit int) *PrescriptionHandler {
	return &PrescriptionHandler{
		PatientID: patientID,
		Offset:    offset,
		Limit:     limit,
	}
}

func (h *PrescriptionHandler) GetPrescriptions() ([]*Prescription, error) {
	// 用户信息
	userInfo, err := models.GetUserByID(h.PatientID)
	if err != nil {
		return nil, err
	}

	// 处方信息
	prescriptions, err := models.GetPatientPrescriptions(h.Offset, h.Limit, h.PatientID)
	if err != nil {
		return nil, err
	}

	ps := make([]*Prescription, len(prescriptions))
	for i := 0; i < len(prescriptions); i++ {
		ps[i] = &Prescription{
			ID:          prescriptions[i].ID,
			Title:       prescriptions[i].Title,
			Desc:        prescriptions[i].Desc,
			CreatedOn:   prescriptions[i].CreatedOn,
			PatientName: userInfo.Name,
			Sex:         userInfo.Sex,
			Age:         userInfo.Age,
			Trainings:   h.getTrainings(prescriptions[i].ID),
		}
	}

	return ps, nil
}

// 根据处方id获取运动处方信息（多条）
func (h *PrescriptionHandler) getTrainings(prescriptionID int) []*Training {
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
