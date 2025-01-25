package models

import "gorm.io/gorm"

type Album struct {
	gorm.Model
	Name   string `gorm:"unique"`
	BandID uint
	Band   Band `gorm:"foreignKey:BandID;references:id"`
	//Songs []Song `gorm:"foreignKey:Name;references:Name"`
}
