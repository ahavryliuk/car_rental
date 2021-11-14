package service

import (
	"carRentalVivino/pkg/storage"
	"errors"
	"github.com/google/uuid"
	"time"
)

func NewCarService(bookingManager *storage.BookingManager, carManager *storage.CarManager) *CarService {
	return &CarService{
		bookingManager: bookingManager,
		carManager:     carManager,
	}
}

type CarService struct {
	bookingManager *storage.BookingManager
	carManager     *storage.CarManager
}

func (s *CarService) AssertDateRangeAvailability(carUUID uuid.UUID, dateFrom, dateTo time.Time) error {
	cars, err := s.carManager.GetByUUIDAndDateRange(carUUID, dateFrom, dateTo)
	if err != nil {
		return err
	}

	if len(*cars) < 1 {
		return errors.New("no cars available for given date range")
	}

	return nil
}
