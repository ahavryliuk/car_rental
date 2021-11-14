package storage

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
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
