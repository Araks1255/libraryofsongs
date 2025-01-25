package models

import (
	"gorm.io/gorm"
)

type Song struct {
	gorm.Model
	Name    string `gorm:"unique"`
	AlbumID uint
	Album   Album `gorm:"foreignKey:AlbumID;references:id"`
	// Band          string
	// Album         string
	// Genre         string
	// YearOfRelease string
}
