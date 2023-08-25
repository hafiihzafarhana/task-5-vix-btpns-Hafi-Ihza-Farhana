package models

import "time"

type PhotoModel struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	Caption   string    `gorm:"not null"`
	PhotoUrl  string    `gorm:"not null"`
	UserID    uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
