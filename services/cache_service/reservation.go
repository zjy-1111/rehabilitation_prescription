package cache_service

import (
	"strconv"
	"strings"

	"rehabilitation_prescription/pkg/e"
)

type Reservation struct {
	ID         int
	Name       string
	Time       int
	DoctorName string
	Address    string
	CreatedBy  string
	ModifiedBy string

	PageNum  int
	PageSize int
}

func (r *Reservation) GetReservationKey() string {
	keys := []string{
		e.CACHE_RESERVATION,
		"LIST",
	}

	if r.Name != "" {
		keys = append(keys, r.Name)
	}
	if r.PageNum > 0 {
		keys = append(keys, strconv.Itoa(r.PageNum))
	}
	if r.PageSize > 0 {
		keys = append(keys, strconv.Itoa(r.PageSize))
	}

	return strings.Join(keys, "_")
}
