package db

import (
	"log"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{}) // Открытие бд по ссылке из аргумента

	if err != nil { // Обработка ошибок
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Song{}) // Создание таблицы в открытой бд по модели из пакета models

	return db // Возвращение открытой бд с таблицей
}
