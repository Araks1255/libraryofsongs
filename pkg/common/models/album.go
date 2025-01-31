package models

import "gorm.io/gorm"

type Album struct { // Структура альбома
	gorm.Model
	Name   string `gorm:"unique"`
	BandID uint   // Тут всё тоже самое, что в песне (столбец id создаётся автоматически)
	Band   Band   `gorm:"foreignKey:BandID;references:id"`
}
