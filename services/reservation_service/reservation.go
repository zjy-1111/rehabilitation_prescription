package reservation_service

import (
	"encoding/json"
	"rehabilitation_prescription/models"
	"rehabilitation_prescription/pkg/gredis"
	"rehabilitation_prescription/pkg/logging"

	"rehabilitation_prescription/services/cache_service"
)

type Reservation struct {
	ID         int
	Name       string
	Time       int
	DoctorName string
	Address    string
	CreatedBy  string
	ModifiedBy string
	PageNum    int
	PageSize   int
}

func (r *Reservation) ExistByID() (bool, error) {
	return models.ExistReservationByID(r.ID)
}

func (r *Reservation) Add() error {
	return models.AddReservation(r.Time, r.Name, r.DoctorName, r.Address, r.CreatedBy)
}

func (r *Reservation) Edit() error {
	data := make(map[string]interface{})
	data["modified_by"] = r.ModifiedBy
	data["name"] = r.Name
	data["time"] = r.Time
	data["doctor_name"] = r.DoctorName
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
		reservations, cacheReservations []models.Reservation
	)

	cache := cache_service.Reservation{
		PageNum:  r.PageNum,
		PageSize: r.PageSize,
	}
	key := cache.GetReservationKey()
	if gredis.Exists(key) {
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheReservations)
			return cacheReservations, nil
		}
	}

	reservations, err := models.GetReservations(r.PageNum, r.PageSize, r.getMaps())
	if err != nil {
		return nil, err
	}

	gredis.Set(key, reservations, 3600)
	return reservations, nil
}

func (r *Reservation) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["name"] = r.Name
	maps["time"] = r.Time
	maps["doctor_name"] = r.DoctorName
	maps["address"] = r.Address
	maps["created_by"] = r.CreatedBy

	return maps
}
