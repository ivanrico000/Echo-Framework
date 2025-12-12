package core

import "time"

type User struct {
	ID      int64  `gorm:"primaryKey;autoIncrement"`
	HotelID *int64 `gorm:"index"`
	RoleID  *int64 `gorm:"index"`

	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"size:255;not null"`
	Name     string `gorm:"not null"`
	Phone    string `gorm:"type:varchar(10)"`

	IsActive  bool      `gorm:"default:true"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}
