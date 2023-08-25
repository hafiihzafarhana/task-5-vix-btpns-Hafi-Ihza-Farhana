package models

import "time"

type RoleModel struct {
	ID        uint      `gorm:"primaryKey"`
	Role      string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
