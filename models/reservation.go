package models

import (
	"github.com/jinzhu/gorm"
)

type Reservation struct {
	Model
	Name       string `json:"name"`
	Time       int    `json:"time"`
	DoctorName string `json:"doctor_name"`
	Address    string `json:"address"`
	CreatedBy  string `json:"created_by"`
}

// ExistReservationByID checks if an article exists based on ID
func ExistReservationByID(id int) (bool, error) {
	var reserv Reservation
	err := db.Select("id").Where("id = ?", id).First(&reserv).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if reserv.ID > 0 {
		return true, nil
	}

	return false, nil
}

// AddReservation Add a reservation
func AddReservation(t int, name, doctorName, addr, createdBy string) error {
	reserv := Reservation{
		Name:       name,
		Time:       t,
		DoctorName: doctorName,
		Address:    addr,
		CreatedBy:  createdBy,
	}

	if err := db.Create(&reserv).Error; err != nil {
		return err
	}

	return nil
}

// GetReservations gets a list of reservations based on paging and constraints
func GetReservations(pageNum, pageSize int, maps interface{}) ([]Reservation, error) {
	var (
		reservations []Reservation
		err          error
	)

	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&reservations).Error
	} else {
		err = db.Where(maps).Find(&reservations).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return reservations, nil
}

// GetReservationTotal counts the total number of reservations based on the constraint
func GetReservationTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Reservation{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteReservation delete a reservation
func DeleteReservation(id int) error {
	if err := db.Where("id = ?", id).Delete(&Reservation{}).Error; err != nil {
		return err
	}

	return nil
}

// EditReservation modify a single reservation
func EditReservation(id int, data interface{}) error {
	if err := db.Model(&Reservation{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}
