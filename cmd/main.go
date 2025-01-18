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

	port := viper.Get("PORT").(string)
	dbUrl := viper.Get("DB_URL").(string)

	router := gin.Default()
	db := db.Init(dbUrl)

	songs.RegisterRoutes(router, db)

	router.Run(port)
}
