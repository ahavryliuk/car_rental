package storage

import (
	"carRentalVivino/pkg/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func BuildDb(cfg config.MySQLConfig) *gorm.DB {
	dsn := buildDbDSN(cfg)

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Fatal("Error trying to open mysql sql driver")
	}

	return db
}

func buildDbDSN(cfg config.MySQLConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
}
