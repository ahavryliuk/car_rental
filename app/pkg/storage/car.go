package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Car struct {
	gorm.Model `json:"-"`
	UUID       uuid.UUID `json:"uuid"`
	Type       string    `json:"type"`
}
