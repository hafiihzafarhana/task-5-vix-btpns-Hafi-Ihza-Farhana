package models

import "time"

type UserSessionModel struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       string    `gorm:"not null"`
	Status       string    `gorm:"not null"`
	AccessToken  string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}
