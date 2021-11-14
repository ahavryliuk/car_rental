package storage

import (
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model  `json:"-"`
	UUID        uuid.UUID `gorm:"index;<-:create" sql:"type:uuid" json:"uuid"`
	AccessToken string    `gorm:"<-:create" json:"-"`
	DateFrom    time.Time `json:"date_from"`
	DateTo      time.Time `json:"date_to"`
	CarId       uint      `json:"-"`
	Car         Car       `json:"car"`
}

func (b *Booking) BeforeCreate(db *gorm.DB) (err error) {
	b.UUID = uuid.New()
	b.AccessToken = uuid.NewString()

	return
}

func (b *Booking) MarshalBinary() ([]byte, error) {
	return json.Marshal(b)
}

func (b *Booking) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, b); err != nil {
		return err
	}

	return nil
}
