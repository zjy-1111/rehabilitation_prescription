package services

import (
	"rehabilitation_prescription/models"
)

type PatientsHandler struct {
	DoctorID int
	Offset int
	Limit int
}

type Patient struct {
	ID int
	Name string
}

func NewPatientsHandler(doctorID, offset, limit int) *PatientsHandler {
	return &PatientsHandler{
		DoctorID: doctorID,
		Offset: offset,
		Limit: limit,
	}
}

func (h *PatientsHandler) GetPatients() ([]*Patient, error) {
	IDs, err := h.GetPatientIDs()
	if err != nil {
		return nil, err
	}

	patients := make([]*Patient, len(IDs))
	for i := 0; i < len(IDs); i++ {
		p, err := models.GetUserByID(IDs[i])
		if err != nil { continue }
		patients[i] = &Patient{
			ID: p.ID,
			Name: p.Name,
		}
	}

	return patients, nil
}

func (h *PatientsHandler) GetPatientIDs() ([]int, error) {
	IDs, err := models.GetPatientsByDoctorID(h.Offset, h.Limit, h.DoctorID)
	if err != nil {
		return nil, err
	}

	return IDs, nil
}
