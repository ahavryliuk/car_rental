package storage

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCarManager_GetByUUID(t *testing.T) {
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()

	carManager := &CarManager{db: dbMock}

	sqlMock.ExpectQuery("cars").WillReturnRows(
		sqlmock.NewRows([]string{"id", "type", "uuid"}).
			AddRow(1, "Some type", uuid.New()),
	)

	car, err := carManager.GetByUUID(uuid.New())
	if err != nil {
		t.Errorf("error was not expected while getting record: %s", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.NotEmpty(t, car)
}

func TestCarManager_GetAvailableByDateRange(t *testing.T) {
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()
	carManager := &CarManager{db: dbMock}

	sqlMock.ExpectQuery("cars").WillReturnRows(
		sqlmock.NewRows([]string{"id", "type", "uuid"}).
			AddRow(1, "Some type", uuid.New()).
			AddRow(2, "Some other type", uuid.New()),
	)

	cars, err := carManager.GetAvailableByDateRange(time.Now(), time.Now().AddDate(0, 0, 1))
	if err != nil {
		t.Errorf("error was not expected while getting record: %s", err)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	assert.True(t, len(*cars) == 2)
}
