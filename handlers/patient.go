package handlers

import (
	"rehabilitation_prescription/models"
	"strings"
)

type PatientsHandler struct {
	DoctorID int
	Offset   int
	Limit    int
}

type Patient struct {
	ID          int    `json:"id"`
	PatientID   int    `json:"patient_id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Sex         string `json:"sex"`
	Age         int    `json:"age"`
	Description string `json:"description"`
}

func NewPatientsHandler(doctorID, offset, limit int) *PatientsHandler {
	return &PatientsHandler{
		DoctorID: doctorID,
		Offset:   offset,
		Limit:    limit,
	}
}

func (h *PatientsHandler) GetPatients(patientName string) ([]*Patient, error) {
	// [][0]是d-p关系行id，[][1]是patient_id
	var IDs [][2]int
	var err error
	if patientName != "" {
		IDs, err = models.GetAllPatients(h.DoctorID)
	} else {
		IDs, err = models.GetPatientsByDoctorID(h.Offset, h.Limit, h.DoctorID)
	}
	if err != nil {
		return nil, err
	}

	patients := make([]*Patient, 0)
	for i := 0; i < len(IDs); i++ {
		p, err := models.GetUserByID(IDs[i][1])
		// 当查询的patientName不为空时，跳过不包含patientName的patient
		if err != nil || (patientName != "" && !strings.Contains(p.Name, patientName)) {
			continue
		}
		patients = append(patients, &Patient{
			ID:          IDs[i][0],
			PatientID:   p.ID,
			Name:        p.Name,
			Sex:         p.Sex,
			Age:         p.Age,
			Description: p.Description,
			Avatar:      p.Avatar,
		})
	}

	return patients, nil
}
