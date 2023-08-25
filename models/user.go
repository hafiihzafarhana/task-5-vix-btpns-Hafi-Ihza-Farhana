package models

import "time"

type UserModel struct {
	ID       uint         `gorm:"primaryKey"`
	Username string       `gorm:"not null"`
	Email    string       `gorm:"not null"`
	Password string       `gorm:"not null"`
	Photos   []PhotoModel `gorm:"foreignKey:UserID"`
	// UserSessions []UserSessionModel `gorm:"foreignKey:UserID"`
	RoleID    uint      `gorm:"not null"` // Relasi one-to-one ke UserRoleModel
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Role      RoleModel `gorm:"foreignKey:RoleID"`
}
