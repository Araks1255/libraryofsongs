package models

import "gorm.io/gorm"

type Band struct {
	gorm.Model
	Name    string `gorm:"unique"`
	GenreID uint
	Genre   Genre `gorm:"foreignKey:GenreID;references:id"`
	//Albums []Album `gorm:"foreignKey:Name;references:Name"`
}
