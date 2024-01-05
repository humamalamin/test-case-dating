package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID         string         `json:"id"`
	FirstName  string         `json:"first_name"`
	LastName   string         `json:"last_name"`
	Email      string         `json:"email"`
	Gender     int            `json:"gender"`
	Password   string         `json:"password"`
	BirthDate  time.Time      `json:"birth_date"`
	Location   string         `json:"location"`
	Lat        float64        `json:"lat"`
	Lng        float64        `json:"lng"`
	IsPremium  bool           `json:"is_premium"`
	VerifiedAt time.Time      `json:"verified_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

func (rm *User) BeforeCreate(tx *gorm.DB) (err error) {
	id := uuid.New()
	now := time.Now()

	rm.ID = id.String()
	rm.CreatedAt = now
	rm.UpdatedAt = now

	return
}
