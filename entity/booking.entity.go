package entity

import "time"

type Booking struct {
	Id     int       `gorm:"primaryKey"`
	Range  uint8     `gorm:"not null"`
	Hour   time.Time `gorm:"not null"`
	UserId int64
}
