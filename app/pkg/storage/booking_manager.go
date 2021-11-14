package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func NewBookingManager(db *gorm.DB) *BookingManager {
	return &BookingManager{
		db: db,
	}
}

type BookingManager struct {
	db *gorm.DB
}

func (m *BookingManager) GetByUUID(uuid uuid.UUID) (*Booking, error) {
	var booking Booking

	result := m.db.Preload(clause.Associations).Where(&Booking{UUID: uuid}).First(&booking)

	return &booking, result.Error
}

func (m *BookingManager) GetByUUIDAndAccessToken(uuid uuid.UUID, accessToken string) (*Booking, error) {
	var booking Booking

	result := m.db.Where(&Booking{UUID: uuid, AccessToken: accessToken}).First(&booking)

	return &booking, result.Error
}

func (m *BookingManager) RemoveByUUIDAndAccessToken(uuid uuid.UUID, accessToken string) error {
	result := m.db.Where(&Booking{UUID: uuid, AccessToken: accessToken}).Delete(&Booking{})

	return result.Error
}

func (m *BookingManager) Persist(booking *Booking) error {
	var result *gorm.DB

	if booking.ID == 0 {
		result = m.db.Create(booking)
	} else {
		result = m.db.Save(booking)
	}

	return result.Error
}
