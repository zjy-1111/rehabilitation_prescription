package handlers

import (
	"rehabilitation_prescription/models"
	"rehabilitation_prescription/pkg/logging"
)

type PrescriptionHandler struct {
	DoctorID  int
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

func NewPrescriptionHandler(doctorID, patientID, offset, limit int) *PrescriptionHandler {
	return &PrescriptionHandler{
		DoctorID:  doctorID,
		PatientID: patientID,
		Offset:    offset,
		Limit:     limit,
	}
}

func (h *PrescriptionHandler) GetPrescriptions() ([]*Prescription, error) {
	if h.DoctorID != 0 {
		return h.GetHistoryPrescriptions()
	}
	return h.GetPatientPrescriptions()
}

func (h *PrescriptionHandler) GetPatientPrescriptions() ([]*Prescription, error) {
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
			Trainings:   getTrainings(prescriptions[i].ID),
		}
	}

	return ps, nil
}

func (h *PrescriptionHandler) GetHistoryPrescriptions() ([]*Prescription, error) {
	prescriptions, err := models.GetDoctorPrescriptions(h.Offset, h.Limit, h.DoctorID)
	if err != nil {
		return nil, err
	}

	ps := make([]*Prescription, len(prescriptions))
	for i := 0; i < len(prescriptions); i++ {
		userInfo, err := models.GetUserByID(prescriptions[i].PatientID)
		if err != nil {
			logging.Error("get userInfo failed, error: %v", err)
			continue
		}
		ps[i] = &Prescription{
			ID:          prescriptions[i].ID,
			Title:       prescriptions[i].Title,
			Desc:        prescriptions[i].Desc,
			CreatedOn:   prescriptions[i].CreatedOn,
			PatientName: userInfo.Name,
			Sex:         userInfo.Sex,
			Age:         userInfo.Age,
			Trainings:   getTrainings(prescriptions[i].ID),
		}
	}

	return ps, err
}
