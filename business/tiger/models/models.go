package models

import (
	"github.com/sourava/tiger/business/user/models"
	"gorm.io/gorm"
)

type Tiger struct {
	gorm.Model
	UserID            uint
	User              models.User
	Name              string  `gorm:"not null"`
	DateOfBirth       string  `gorm:"not null"`
	LastSeenTimestamp int     `gorm:"not null"`
	LastSeenLatitude  float64 `gorm:"not null"`
	LastSeenLongitude float64 `gorm:"not null"`
}
