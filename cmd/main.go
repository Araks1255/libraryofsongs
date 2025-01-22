package main

import (
	"github.com/Araks1255/libraryofsongs/pkg/common/db"
	"github.com/Araks1255/libraryofsongs/pkg/songs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./pkg/common/envs/.env")
	viper.ReadInConfig()

	port := viper.Get("PORT").(string) // Получение переменных среды
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default() // Инициализация роутера
	db := db.Init(dbUrl)    // И базы данных (самописным методом из пакета db)

	songs.RegisterRoutes(router, db) // Регистрируем роутеры функцией из пакета songs

	router.Run(port) // Запускаем
}
