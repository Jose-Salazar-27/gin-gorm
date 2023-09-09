package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Bookings  []Booking `gorm:"foreignKey:UserId"`
}

// Hooks
func (usr *User) BeforeSave(tx *gorm.DB) error {
	fmt.Println("Hook executed")
	return nil
}
