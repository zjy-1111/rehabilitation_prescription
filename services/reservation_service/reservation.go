package reservation_service

import (
	"rehabilitation_prescription/models"
)

type Reservation struct {
	ID        int
	Name      string
	Date      int
	PeriodID  int
	DoctorID  int
	Address   string
	CreatedBy string
	PageNum   int
	PageSize  int
}

func (r *Reservation) ExistByID() (bool, error) {
	return models.ExistReservationByID(r.ID)
}

func (r *Reservation) Add() error {
	return models.AddReservation(r.Date, r.PeriodID, r.DoctorID, r.Name, r.Address, r.CreatedBy)
}

func (r *Reservation) Edit() error {
	data := make(map[string]interface{})
	data["name"] = r.Name
	data["Date"] = r.Date
	data["period_id"] = r.PeriodID
	data["doctor_id"] = r.DoctorID
	data["address"] = r.Address

	return models.EditReservation(r.ID, data)
}

func (r *Reservation) Delete() error {
	return models.DeleteReservation(r.ID)
}

func (r *Reservation) Count() (int, error) {
	return models.GetReservationTotal(r.getMaps())
}

func (r *Reservation) GetAll() ([]models.Reservation, error) {
	var (
		reservations []models.Reservation
		// cacheReservations []models.Reservation
	)

	//cache := cache_service.Reservation{
	//	Name:      r.Name,
	//	Date:      r.Date,
	//	DoctorID:  r.DoctorID,
	//	CreatedBy: r.CreatedBy,
	//	PageNum:   r.PageNum,
	//	PageSize:  r.PageSize,
	//}
	//key := cache.GetReservationKey()
	//if gredis.Exists(key) {
	//	data, err := gredis.Get(key)
	//	if err != nil {
	//		logging.Info(err)
	//	} else {
	//		json.Unmarshal(data, &cacheReservations)
	//		return cacheReservations, nil
	//	}
	//}

	reservations, err := models.GetReservations(r.PageNum, r.PageSize, r.getMaps())
	if err != nil {
		return nil, err
	}

	//gredis.Set(key, reservations, 3600)
	return reservations, nil
}

func (r *Reservation) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if r.Name != "" {
		maps["name"] = r.Name
	}
	if r.DoctorID != 0 {
		maps["doctor_id"] = r.DoctorID
	}
	if r.Date != 0 {
		maps["Date"] = r.Date

	}
	if r.PeriodID != 0 {
		maps["period_id"] = r.PeriodID
	}
	if r.CreatedBy != "" {
		maps["created_by"] = r.CreatedBy
	}

	return maps
}
