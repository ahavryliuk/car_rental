package request

import "time"

type CarList struct {
	DateFrom time.Time `json:"date_from" binding:"required"`
	DateTo   time.Time `json:"date_to" binding:"required"`
}
