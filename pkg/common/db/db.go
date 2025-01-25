package db

import (
	"log"

	"github.com/Araks1255/libraryofsongs/pkg/common/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}) // Открытие бд по ссылке из аргумента

	if err != nil { // Обработка ошибок
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Song{}, &models.Album{}, &models.Band{}, &models.Genre{}) // Создание таблицы в открытой бд по модели из пакета models

	return db // Возвращение открытой бд с таблицей
}
