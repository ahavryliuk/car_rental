package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func NewCarManager(db *gorm.DB) *CarManager {
	return &CarManager{
		db: db,
	}
}

type CarManager struct {
	db *gorm.DB
}

func (m *CarManager) GetByUUID(uuid uuid.UUID) (*Car, error) {
	var car Car

	result := m.db.Where(&Car{UUID: uuid}).First(&car)

	return &car, result.Error
}

func (m *CarManager) GetAvailableByDateRange(dateFrom, dateTo time.Time) (*[]Car, error) {
	var cars []Car

	query := `
	SELECT * FROM cars c
	WHERE NOT EXISTS
		 (SELECT * FROM bookings
		  WHERE car_id = c.id
			 AND date_from < ?
			 AND date_to > ?)`

	result := m.db.Raw(query, dateTo, dateFrom).Scan(&cars)

	return &cars, result.Error
}

func (m *CarManager) GetByUUIDAndDateRange(uuid uuid.UUID, dateFrom, dateTo time.Time) (*[]Car, error) {
	var cars []Car

	query := `
	SELECT * FROM cars c
	WHERE uuid = ?
		 AND NOT EXISTS
		 (SELECT * FROM bookings
		  WHERE car_id = c.id
			 AND date_from < ?
			 AND date_to > ?)`

	result := m.db.Raw(query, uuid, dateTo, dateFrom).Scan(&cars)

	return &cars, result.Error
}
