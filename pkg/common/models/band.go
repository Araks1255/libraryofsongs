package models

import "gorm.io/gorm"

type Band struct { // Структура группы (музыкальной)
	gorm.Model
	Name    string `gorm:"unique"`
	GenreID uint   // Всё аналогично песне и альбому
	Genre   Genre  `gorm:"foreignKey:GenreID;references:id"`
}
