package models

import "gorm.io/gorm"

type Genre struct { // Жанр
	gorm.Model        // Тут из особых полей только название, остальное гормом создаётся автоматически (айди, время создания, удаления и т.д.)
	Name       string `gorm:"unique"`
}
