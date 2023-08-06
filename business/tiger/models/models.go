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

type TigerSighting struct {
	gorm.Model
	UserID    uint
	User      models.User
	TigerID   uint
	Tiger     Tiger
	Timestamp int     `gorm:"not null"`
	Latitude  float64 `gorm:"not null"`
	Longitude float64 `gorm:"not null"`
	Image     string
}
