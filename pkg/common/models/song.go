package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Name          string
	Band          string
	Album         string
	Genre         string
	YearOfRelease string
}
