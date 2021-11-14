package storage

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestBookingManager_GetByUUID(t *testing.T) {
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()
	bookingManager := &BookingManager{db: dbMock}

	sqlMock.ExpectQuery("bookings").WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "access_token", "date_from", "date_to", "car_id"}).
			AddRow(1, uuid.New(), uuid.New(), time.Now(), time.Now(), 1),
	)
	sqlMock.ExpectQuery("cars").WillReturnRows(
		sqlmock.NewRows([]string{"id", "type", "uuid"}).
			AddRow(1, "Some type", uuid.New()),
	)

	booking, err := bookingManager.GetByUUID(uuid.New())
	if err != nil {
		t.Errorf("error was not expected while getting record: %s", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NotEmpty(t, booking.UUID)
	assert.NotEmpty(t, booking.UUID)
	assert.NotEmpty(t, booking.AccessToken)
}

func TestBookingManager_GetByUUIDAndAccessToken(t *testing.T) {
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()
	bookingManager := &BookingManager{db: dbMock}

	sqlMock.ExpectQuery("bookings").WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "access_token", "date_from", "date_to", "car_id"}).
			AddRow(1, uuid.New(), uuid.New(), time.Now(), time.Now(), 1),
	)

	booking, err := bookingManager.GetByUUIDAndAccessToken(uuid.New(), uuid.NewString())
	if err != nil {
		t.Errorf("error was not expected while getting record: %s", err)
	}

	assert.NotEmpty(t, booking)
}

func TestBookingManager_RemoveByUUIDAndAccessTokenOnFailure(t *testing.T) {
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()
	bookingManager := &BookingManager{db: dbMock}

	sqlMock.ExpectBegin()
	sqlMock.ExpectQuery("bookings").WillReturnRows()
	sqlMock.ExpectRollback()

	err := bookingManager.RemoveByUUIDAndAccessToken(uuid.New(), uuid.NewString())

	assert.Error(t, err)
}
