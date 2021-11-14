package request

import (
	"github.com/google/uuid"
	"time"
)

type BookingCreate struct {
	DateFrom time.Time `json:"date_from" binding:"required"`
	DateTo   time.Time `json:"date_to" binding:"required"`
	CarUUID  uuid.UUID `json:"car_id" binding:"required"`
}

type BookingUpdate struct {
	DateFrom time.Time `json:"date_from" binding:"required"`
	DateTo   time.Time `json:"date_to" binding:"required"`
	CarUUID  uuid.UUID `json:"car_id" binding:"required"`
}
