package handler

import (
	"carRentalVivino/pkg/config"
	"carRentalVivino/pkg/service"
	"carRentalVivino/pkg/storage"
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func getMockedConnection(t *testing.T) (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	conMock, sqlMock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormMock, _ := gorm.Open(mysql.New(mysql.Config{
		Conn: conMock,
	}), &gorm.Config{})

	return gormMock, sqlMock, conMock
}

func TestBookingHandler_Get(t *testing.T) {
	gin.SetMode(gin.TestMode)
	dbMock, sqlMock, conMock := getMockedConnection(t)
	defer conMock.Close()
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)

	bookingManager := storage.NewBookingManager(dbMock)
	carManager := storage.NewCarManager(dbMock)

	bookingHandler := BookingHandler{
		bookingManager: bookingManager,
		carManager:     carManager,
		carService:     service.NewCarService(bookingManager, carManager),
		redisService: service.NewRedisService(context.Background(), &config.Config{
			Redis: config.RedisConfig{
				Host:     "",
				Port:     0,
				Password: "",
			},
		}),
	}

	sqlMock.ExpectQuery("bookings").WillReturnRows(
		sqlmock.NewRows([]string{"id", "uuid", "access_token", "date_from", "date_to", "car_id"}).
			AddRow(1, uuid.New(), uuid.New(), time.Now(), time.Now(), 1),
	)
	sqlMock.ExpectQuery("cars").WillReturnRows(
		sqlmock.NewRows([]string{"id", "type", "uuid"}).
			AddRow(1, "Some type", uuid.New()),
	)

	_, err := http.NewRequest(http.MethodGet, "https://app/api/v1/bookings/201a9cdd-2d09-485f-baac-9c2f2ebfb420", nil)
	if err != nil {
		t.Fatalf("an error '%s' was not expected while creating request", err)
	}

	ctx.Params = []gin.Param{
		{
			Key:   "uuid",
			Value: "201a9cdd-2d09-485f-baac-9c2f2ebfb420",
		},
	}
	bookingHandler.Get(ctx)

	if recorder.Code != 200 {
		t.Fatalf("expected status code to be 200, but got: %d", recorder.Code)
	}

	if err := sqlMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
