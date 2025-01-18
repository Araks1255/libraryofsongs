package db

import (
	"log"

	"github.com/Araks1255/libraryofsongs"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config)

	if err != nil {
		log.Fatal(err)
	}

	return db
}